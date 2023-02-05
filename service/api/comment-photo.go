package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Write a comment on a photo of a user
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get photoID from path
	// The Photo ID in the path is a string and coincides with the photo we are commenting
	photo_id := rt.getPathParameter("photo_id", ps)
	if photo_id == "" {
		ctx.Logger.Error("wrong photo_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 2. Get comment from request body
	// If comment is a reply comment, parentId will contain the comment id of parent comment;
	// "" otherwise
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		ctx.Logger.Error("The body is not a parseable JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !comment.IsValid() {
		// Here we validated the comment structure content
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "Invalid comment. Invalid characters inserted or comment too long."}`))
		return
	}
	// 3. Check if the user is authenticated
	// We want to allow only to logged users to write comments.
	// Therefore the authentication token in the header should coincides with the username of the writer
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)
	id, _ := rt.db.GetIdByName(comment.Author)
	if authtoken == "" || authtoken != id {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		return
	}

	// 4. Write comment in the database
	comment_db, err := rt.db.CommentPhoto(photo_id, comment.ToDatabase(rt.db))
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("error while adding the comment to the database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment.FromDatabase(comment_db, rt.db)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)

}

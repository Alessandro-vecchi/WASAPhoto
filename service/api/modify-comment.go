package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) modifyComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1. Retrieve Id of the comment the user want to modify.
	comment_id := rt.getPathParameter("comment_id", ps)
	if comment_id == "" {
		ctx.Logger.Error("wrong comment_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 2. Get comment from request body (comment body, author name)
	var comment models.Comment
	comment.CommentId = comment_id
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		ctx.Logger.Error("The body is not a parseable JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !comment.IsValid() {
		// Here we validated the comment structure is valid
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "Invalid comment. Invalid characters inserted or comment too long."}`))
		return
	}
	// 3. Check if user who want to modify comment is the same who wrote it.
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)
	id, _ := rt.db.GetIdByName(comment.Author)
	err = checkUserIdentity(authtoken, id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "User does not exist"}`))
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		return
	}
	// 4. Modify comment from the database
	_, err = rt.db.ModifyComment(comment.ToDatabase(rt.db))

	if errors.Is(err, database.ErrCommentNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "The comment does not exist"}`))
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("comment_id", comment_id).Error("can't modify comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send partial comment modifications to the client
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)
}

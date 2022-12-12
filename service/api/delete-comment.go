package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1. Retrieve Id of the comment the user want to delete.
	comment_id := rt.getPathParameter("comment_id", ps)
	if comment_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Get comment from request body
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	fmt.Println(comment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.IsValid() {
		// Here we validated the comment structure is valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 3. Check if user who want to delete comment is the same who wrote it.
	authtoken := r.Header.Get("authToken")
	log.Printf("The authentication token in the header is: %v", authtoken)
	id, _ := rt.db.GetIdByName(comment.Author)
	err = checkUserIdentity(authtoken, id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 4. Delete comment from the database
	err = rt.db.UncommentPhoto(comment_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

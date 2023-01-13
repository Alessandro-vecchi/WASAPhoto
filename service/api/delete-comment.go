package api

import (
	"errors"
	"log"
	"net/http"

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
	// 2. Check if the comment exists.
	comment_db, err := rt.db.GetSingleComment(comment_id)
	if errors.Is(err, database.ErrCommentNotExists) {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 3. Check if user who want to delete comment is the same who wrote it.
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)
	err = checkUserIdentity(authtoken, comment_db.UserId, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
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

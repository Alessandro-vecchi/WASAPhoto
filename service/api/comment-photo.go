package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

// Write a comment on a photo of a user
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get photoID from path
	// The Photo ID in the path is a string and coincides with the photo we are commenting
	rawPhoto_id := ps.ByName("photo_id")
	photo_id := strings.TrimPrefix(rawPhoto_id, ":photo_id=")
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Check if the user is authenticated
	// We want to allow only to the owner of the profile to upload photo,
	// Therefore the user_id must coincides with the authentication token in the header
	authtoken := r.Header.Get("authToken")
	log.Printf("The authentication token in the header is: %v", authtoken)

	err := rt.db.CheckUserIdentity(authtoken, user_id)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 3. get comment from request body
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	fmt.Println(comment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.IsValid() {
		// Here we validated the fountain structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the fountain data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - chiamare metodo DB
	comment_db, err := rt.db.CommentPhoto(photo_id, comment.ToDatabase(rt.db))
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment.FromDatabase(comment_db, rt.db)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)

}

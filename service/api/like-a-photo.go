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

// user b follows user A
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Retrieve photo ID from path.
	photo_id := rt.getPathParameter("photo_id", ps)
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Retrieve ID of the user who want to put like from path.
	user_id := rt.getPathParameter("like_id", ps)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 3. Check if the user is authenticated
	// We want to allow only to logged users to put likes.
	// Therefore the authentication token in the header should coincide with the id of the user who is liking the photo
	authtoken := r.Header.Get("authToken")
	log.Printf("The authentication token in the header is: %v", authtoken)
	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 4. Check that the user is not putting like to an own photo
	if models.IsLikingHimself(photo_id, user_id, rt.db) {
		ctx.Logger.WithError(database.ErrUserCantLikeHimself).Error("like not possible ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 5. Add a like to the database if it's not present,
	// 	  otherwise don't do anything
	err = rt.db.LikePhoto(photo_id, user_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		// 404 Not Found
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the photo that triggered the error.
		ctx.Logger.WithError(err).WithField("id", photo_id).Error("can't put a like to the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5. The like is created
	// Succesfull response: 201 Created
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(user_id)
}

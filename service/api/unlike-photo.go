package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1. Retrieve photo ID from path.
	photo_id := rt.getPathParameter("photo_id", ps)
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Retrieve ID of the user who want to put like from path.
	user_id := rt.getPathParameter("likes", ps)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3. Check if the user is authenticated
	// We want to allow only to logged users to put likes.
	// Therefore the authentication token in the header should coincide with the id of the user who is liking the photo
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)
	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 4. Check that the user is not putting like to an own photo, if the photo exists.
	flag, err := models.IsLikingHimself(photo_id, user_id, rt.db)
	if errors.Is(err, database.ErrPhotoNotExists) {
		ctx.Logger.WithError(database.ErrPhotoNotExists).Error("the photo doesn't exist")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if flag {
		ctx.Logger.WithError(database.ErrUserCantLikeHimself).Error("like not possible ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 5. Remove a like from the database if it's present,
	// 	  otherwise don't do anything
	err = rt.db.UnlikePhoto(photo_id, user_id)
	if errors.Is(err, database.ErrLikeNotPresent) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		ctx.Logger.WithError(database.ErrLikeNotPresent).Error("like already not present")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the photo that triggered the error.
		ctx.Logger.WithError(err).WithField("id", photo_id).Error("database error. Can't unlike photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Succesful response: 204 No Content
	w.WriteHeader(http.StatusNoContent)
}

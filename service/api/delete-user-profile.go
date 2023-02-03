package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1. Get userID from path
	// The User ID in the path is a string and coincides with the profile we are in
	user_id := rt.getPathParameter("user_id", ps)
	if user_id == "" {
		ctx.Logger.Error("wrong user_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 2. Check if the user is authenticated
	// We want to allow only to the owner of the profile to upload photo,
	// Therefore the user_id must coincides with the authentication token in the header
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)

	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		_, _ = w.Write([]byte(`{"error": "User does not exist"}`))
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Delete all images from the folder
	photos, err := rt.db.GetListUserPhotos(user_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, photo := range photos {
		err := rt.deleteImageFromFolder(photo.PhotoId, w, ctx)
		if err != nil {
			ctx.Logger.WithError(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	// Delete user profile. The cascade make sure all the info related are being removed
	err = rt.db.DeleteUserProfile(user_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the fountain that triggered the error.
		ctx.Logger.WithError(err).WithField("id", user_id).Error("can't delete the profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("profile deleted succesfully: %v", err)
	w.WriteHeader(http.StatusNoContent)
}

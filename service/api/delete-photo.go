package api

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The photo ID in the path is a string. Let's parse it.
	photo_id := rt.getPathParameter("photo_id", ps)
	if photo_id == "" {
		ctx.Logger.Error("wrong photo_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Get user id from photo]
	photo, err := rt.db.GetUserPhoto(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		_, _ = w.Write([]byte(`{"error": "The photo does not exist"}`))
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if user who want to delete comment is the same who wrote it.
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)
	err = checkUserIdentity(authtoken, photo.UserId, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		_, _ = w.Write([]byte(`{"error": "User does not exist"}`))
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = rt.deleteImageFromFolder(photo_id, w, ctx)
	if err != nil {
		ctx.Logger.WithError(err).Error("error in deleting images from folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.DeletePhoto(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		_, _ = w.Write([]byte(`{"error": "Photo does not exist"}`))
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the fountain that triggered the error.
		ctx.Logger.WithError(err).WithField("id", photo_id).Error("can't delete the photo from the database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) deleteImageFromFolder(photo_id string, w http.ResponseWriter, ctx reqcontext.RequestContext) error {
	photo, err := rt.db.GetUserPhoto(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		_, _ = w.Write([]byte(`{"error": "The photo does not exist"}`))
		w.WriteHeader(http.StatusNotFound)
		return err
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide photo")
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	current_directory, _ := os.Getwd()
	folder_name := "images"
	path := filepath.Join(current_directory, folder_name, photo.Image)
	err = os.Remove(path)
	if err != nil {
		// handle the error
		ctx.Logger.WithError(err).WithField("id", photo_id).Error("image could not be removed from the folder")
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

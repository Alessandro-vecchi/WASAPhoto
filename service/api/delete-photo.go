package api

import (
	"errors"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := rt.deleteImageFromFolder(photo_id, w, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.DeletePhoto(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
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
	/* current_directory, _ := os.Getwd()
		folder_name := "images"
		path := filepath.Join(current_directory, folder_name, photo_id)
	    pattern := path + "*.{jpeg,jpg,png}"
	    files, _ := filepath.Glob(pattern)
		err = os.Remove(files[0]) */

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) deleteImageFromFolder(photo_id string, w http.ResponseWriter, ctx reqcontext.RequestContext) error {
	photo, err := rt.db.GetUserPhoto(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
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

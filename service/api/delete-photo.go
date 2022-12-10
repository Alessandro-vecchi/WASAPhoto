package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The user ID in the path is a string. Let's parse it.
	photo_id := ps.ByName("photo_id")
	photo_id = strings.TrimPrefix(photo_id, ":photo_id=")
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.DeletePhoto(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the fountain that triggered the error.
		ctx.Logger.WithError(err).WithField("id", photo_id).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

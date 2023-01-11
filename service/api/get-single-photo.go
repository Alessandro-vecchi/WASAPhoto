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

// Get a single photo
func (rt *_router) getUserPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The User ID in the path is a string
	photo_id := rt.getPathParameter("photos", ps)
	if photo_id == "" {
		log.Println("invalid photo_id in the URL")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photo, err := rt.db.GetUserPhoto(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Translating from database to api
	var p models.Photo
	p.FromDatabase(photo, rt.db)
	// checking that the profile is valid
	if p.IsValid() {
		// Send the user profile to the user
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(p)
	}

}

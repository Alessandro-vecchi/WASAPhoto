package api

import (
	"encoding/json"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

// enrollNewUser enrolls a new student in the system, and provides the public key.
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var profile database.Profile_db
	var p models.Profile
	if r.URL.Query().Has("username") {
		name := r.URL.Query().Get("username")
		profile, err = rt.db.GetUserProfileByUsername(name)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the user profile to the user
	w.Header().Set("Content-Type", "application/json")
	// translating from database to api
	p.FromDatabase(profile)
	_ = json.NewEncoder(w).Encode(p)
}

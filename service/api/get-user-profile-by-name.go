package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

// Get the profile of a user
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var profile database.Profile_db
	// 1. Retrieve username searched from the URL query, if properly formed.
	if r.URL.Query().Has("username") {
		name := r.URL.Query().Get("username")
		profile, err = rt.db.GetUserProfileByUsername(name)

	} else {
		// No username field founded
		ctx.Logger.WithError(err).Error("query has no field username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if errors.Is(err, database.ErrUserExists) {
		// User not found in the database
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("database error: can't provide user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// 2. Translating from database to api
	var p models.Profile
	p.FromDatabase(profile, rt.db)
	// Checking that the profile is valid
	if p.IsValid() {
		w.WriteHeader(http.StatusOK)
		// Send the user profile to the user
		_ = json.NewEncoder(w).Encode(p)
	} else {
		// Invalid user profile
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

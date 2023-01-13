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

// Get the profile of a user
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var name string
	// 1. Retrieve username searched from the URL query, if properly formed.
	if r.URL.Query().Has("username") {
		name = r.URL.Query().Get("username")
		// If the username is "" it means that we are in our own profile page
		if name == "" {
			user_id := r.Header.Get("Authorization")
			name, _ = rt.db.GetNameById(user_id)
		}

	} else {
		// No username field founded
		log.Printf("query has no field username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	profile, err := rt.db.GetUserProfileByUsername(name)
	if errors.Is(err, database.ErrUserNotExists) {
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
	// 2. Translating from database to api
	var p models.Profile
	p.FromDatabase(profile, rt.db)
	// Checking that the profile is valid
	w.WriteHeader(http.StatusOK)
	// Send the user profile to the user
	_ = json.NewEncoder(w).Encode(p)

}

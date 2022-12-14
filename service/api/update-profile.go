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

func (rt *_router) updateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1. Get Id of the user whose profile is being updated
	user_id := rt.getPathParameter("user_id", ps)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Check if the user is authenticated
	// We want to allow only to logged users in their own profile to modify the profile.
	// Therefore the authentication token in the header should coincides with the id of the profile
	authtoken := r.Header.Get("authToken")
	log.Printf("The authentication token in the header is: %v", authtoken)
	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 3. Read new profile information from request body
	var p models.Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		// The body was not a parseable JSON object, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !p.IsValid() {
		// Profile data is invalid
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
	p.ID = user_id
	_, err = rt.db.UpdateUserProfile(false, p.ToDatabase())

	if err != nil {
		ctx.Logger.WithError(err).WithField("user_id", user_id).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// w.WriteHeader(http.StatusNoContent)
	// Send the user profile to the user
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p)
}

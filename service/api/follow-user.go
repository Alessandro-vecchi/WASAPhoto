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

// user b follows user A
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get user ID of user A from path
	// The User ID in the path is a string and coincides with the profile we watching
	user_id_A := rt.getPathParameter("user_id", ps)
	if user_id_A == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Get user ID of user B from path
	// It coincides with the user that want to follow
	user_id_B := rt.getPathParameter("followers_id", ps)
	if user_id_B == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 3. Check if the users are the same
	if models.AreTheSame(user_id_A, user_id_B) {
		// A user can't follow himself
		ctx.Logger.WithError(database.ErrUserCantFollowHimself).Error("a user can't follow himself ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 4. Check if the user B is authenticated
	// We want to allow only to a logged in user to follow another user,
	// Therefore the user_id of the user that want to follow must coincide with the authentication token in the header
	authtoken := r.Header.Get("authToken")
	log.Printf("The authentication token in the header is: %v", authtoken)

	err := checkUserIdentity(authtoken, user_id_B, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// 4. Follow user A
	err = rt.db.FollowUser(user_id_A, user_id_B)
	if !errors.Is(err, database.ErrFollowerAlreadyPresent) {
		// user B already follows user A
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5. The follow is created
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	// 6. Encode the name
	_, name_B, err := models.Conversion(user_id_A, user_id_B, rt.db)
	if err != nil {
		ctx.Logger.WithError(err).Error("error while converting user_id and username")
		return
	}
	_ = json.NewEncoder(w).Encode(name_B)
}

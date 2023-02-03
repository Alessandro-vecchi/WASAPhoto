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

// user B bans user A
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get ID of the profile of user A from path
	// The User ID in the path is a string and correspondes with the profile we're watching
	user_id_A := rt.getPathParameter("user_id", ps)
	if user_id_A == "" {
		ctx.Logger.Error("wrong user_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 2. Get ID of user B from path
	// It coincides with the user that want to ban
	user_id_B := rt.getPathParameter("ban_id", ps)
	if user_id_B == "" {
		ctx.Logger.Error("wrong ban_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 3. Check that user is not banning himself
	if models.AreTheSame(user_id_A, user_id_B) {
		// A user can't ban himself
		_, _ = w.Write([]byte(`{"error": "An user can't ban himself"}`))
		w.WriteHeader(http.StatusConflict)
		return
	}
	// 4. Check if the user B is authenticated
	// We want to allow only to a logged in user to ban another user.
	// Therefore the user_id of the user that want to follow must coincide with the authentication token in the header
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)

	err := checkUserIdentity(authtoken, user_id_B, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		_, _ = w.Write([]byte(`{"error": "User does not exist"}`))
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Conversion
	name_A, _, err := models.Conversion(user_id_A, user_id_B, rt.db)
	if err != nil {
		ctx.Logger.WithError(err).Error("error while converting user_id in username")
		return
	}
	// 4. Ban user A
	err = rt.db.BanUser(user_id_A, user_id_B)
	if !errors.Is(err, database.ErrBanAlreadyPresent) {
		// user B already banned user A
		_, _ = w.Write([]byte(`{"error": "You already banned " ` + name_A + `}`))
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("error in the database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5. The ban is created
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	// 6. Encode the name
	_ = json.NewEncoder(w).Encode(name_A)
}

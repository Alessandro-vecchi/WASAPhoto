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

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	// 1. Read the username from the request body.
	var username models.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !username.IsValid() {
		// Checking if the username match its regex
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("invalid characters in username")
		return
	}

	// 2. If it doesn't exist, create the user in the database.
	// Otherwise, return the identifier already present.
	// Note that this function will return a user identifier.
	user_id, err := rt.db.DoLogin(username.Name)
	if errors.Is(err, database.ErrUserExists) {
		// This means that the user already exists in the database,
		// so we are not going to create a new user but just returning the one we have.
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user_id)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 3. The user didn't exist. Return 201 created.
	w.WriteHeader(http.StatusCreated)
	// 4. Send the user_id to the user.
	_ = json.NewEncoder(w).Encode(user_id)
}

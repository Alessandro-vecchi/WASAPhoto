package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Read the new user from the request body.
	var username models.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !username.IsValid() {
		// Here we validated the fountain structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the fountain data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the user in the database. Note that this function will return a user identifier.
	user_id, err := rt.db.DoLogin(username.Name)
	fmt.Println("ciao " + user_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user_id)
}

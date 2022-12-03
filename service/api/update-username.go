package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The User ID in the path is a string
	user_id := ps.ByName("user_id")
	if user_id == "" {
		fmt.Println("siiiii")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// read new username from request body
	// var username string
	var p Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		// The body was not a parseable JSON object, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !p.isValid() {
		// Profile data is invalid
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
	p.ID = user_id
	up, err := rt.db.UpdateUsername(p.ToDatabase())
	var pp Profile
	pp.FromDatabase(up)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("user_id", user_id).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusNoContent)
	// Send the user profile to the user

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(pp)
}

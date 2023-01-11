package api

import (
	"encoding/json"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//  1. Retrieve ID of the user profile we want to follow from the path
	user_id := rt.getPathParameter("users", ps)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followers, err := rt.db.GetFollowers(user_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(followers)
}

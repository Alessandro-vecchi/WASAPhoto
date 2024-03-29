package api

import (
	"encoding/json"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowed(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The User ID in the path is a string
	user_id := rt.getPathParameter("user_id", ps)
	if user_id == "" {
		ctx.Logger.Error("wrong user_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	following, err := rt.db.GetFollowing(user_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	name, _ := rt.db.GetNameById(r.Header.Get("Authorization"))
	var short_prof models.Short_profile
	short_prof.FromDatabase(following, name)
	// Send the list to the user.
	_ = json.NewEncoder(w).Encode(short_prof)
}

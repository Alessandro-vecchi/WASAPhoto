package api

import (
	"encoding/json"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Retrieve photo ID from path.
	photo_id := rt.getPathParameter("photo_id", ps)
	if photo_id == "" {
		ctx.Logger.Error("wrong photo_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	likes, err := rt.db.GetLikes(photo_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list likes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Check if the logged user added the like
	name, _ := rt.db.GetNameById(r.Header.Get("Authorization"))
	var short_prof models.Short_profile
	short_prof.FromDatabase(likes, name)
	// Send list of likes to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(short_prof)
}

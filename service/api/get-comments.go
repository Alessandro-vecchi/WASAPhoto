package api

import (
	"encoding/json"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var listCommentsDb []database.Comment_db

	// 1. Retrieve ID of the photo from the path
	photo_id := rt.getPathParameter("photo_id", ps)
	if photo_id == "" {
		// If empty it's because the function returned a bad values
		ctx.Logger.Error("photo undefined")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	listCommentsDb, err = rt.db.GetComments(photo_id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list comments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var listCommentsAPI []models.Comment
	for _, element := range listCommentsDb {
		var c models.Comment
		c.FromDatabase(element, rt.db)

		listCommentsAPI = append(listCommentsAPI, c)
	}
	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(listCommentsAPI)
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error
	var listPhotos []database.Photo_db

	// The User ID in the path is a string
	user_id := ps.ByName("user_id")
	user_id = strings.TrimPrefix(user_id, ":user_id=")
	fmt.Println(user_id)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.URL.Query().Has("latitude") && r.URL.Query().Has("longitude") {
		fmt.Println("hi")
	} else {
		// Request an unfiltered list of fountains from the DB
		listPhotos, err = rt.db.GetListUserPhotos(user_id)
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var listPhotosAPI []models.Photo
	for _, element := range listPhotos {
		var m models.Photo
		m.FromDatabase(element, rt.db)

		fmt.Println(m.Username)
		listPhotosAPI = append(listPhotosAPI, m)
	}
	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(listPhotosAPI)
}

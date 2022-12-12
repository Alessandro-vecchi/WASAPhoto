package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The User ID in the path is a string
	user_id := rt.getPathParameter("user_id", ps)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// read new username from request body
	// var username string
	var p models.Profile
	p.ID = user_id

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		// The body was not a parseable JSON object, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !p.IsValid() {
		// Profile data is invalid
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name, err := rt.db.GetNameById(user_id)
	if err != nil {
		log.Println("couldn't get name from id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userProfile, err := rt.db.GetUserProfileByUsername(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var pp models.Profile
	pp.FromDatabase(userProfile, rt.db)

	oldProfileBytes, err := json.Marshal(&pp)
	if err != nil {
		log.Println("Error creating json patch", err.Error())
		return
	}
	// fmt.Println(`{"op":"replace", "path": "/username", "value": "` + p.Username + `"}`)
	patchJSON := []byte(`[{"op":"replace", "path": "/username", "value": "` + p.Username + `"}]`)
	// fmt.Println(patchJSON)
	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		log.Println("Error Decoding patch json ", err.Error())
		return
	}
	modified, err := patch.Apply(oldProfileBytes)
	if err != nil {
		log.Println("Error Applying patch json ", err.Error())
		return
	}

	var jsonProfile models.Profile
	err = json.Unmarshal(modified, &jsonProfile)
	if err != nil {
		log.Println("Error unmarshaling patch json ", err.Error())
		return
	}
	// updating profile in the database
	_, err = rt.db.UpdateUserProfile(true, jsonProfile.ToDatabase())
	if err != nil {
		log.Println("Error UPDATING the database ", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(jsonProfile)
}

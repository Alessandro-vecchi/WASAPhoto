package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get ID of the user that want to change name
	user_id := rt.getPathParameter("user_id", ps)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Check if the user is authenticated
	// We want to allow only to the owner of the profile to upload photo,
	// Therefore the user_id must coincides with the authentication token in the header
	authtoken := r.Header.Get("authToken")
	log.Printf("The authentication token in the header is: %v", authtoken)

	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 3. Read new username from request body
	var p models.Profile
	p.ID = user_id

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		// The body was not a parseable JSON object, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if rt.db.CountStuffs("username", "profile", p.Username) > 0 {
		// User Already Exists
		ctx.Logger.WithError(err).WithField("username", p.Username).Error("Cannot use username that already exists")
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
	// Updating profile in the database
	_, err = rt.db.UpdateUserProfile(true, jsonProfile.ToDatabase())
	if err != nil {
		log.Println("Error UPDATING the database ", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Send updated username in response
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(jsonProfile)
}

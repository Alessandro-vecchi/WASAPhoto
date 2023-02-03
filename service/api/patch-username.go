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
		ctx.Logger.Error("wrong user_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 2. Check if the user is authenticated
	// We want to allow only to the owner of the profile to upload photo,
	// Therefore the user_id must coincides with the authentication token in the header
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)

	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "User does not exist"}`))
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		return
	}
	// 3. Read new username from request body
	var u models.Username
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		// The body was not a parseable JSON object, reject it
		ctx.Logger.Error("The body is not a parseable JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if rt.db.CountStuffs("username", "profile", u.Username) > 0 {
		// User Already Exists
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "This username already exists. Please choose another username."}`))
		return
	} else if !u.IsValid() {
		// Username is invalid
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "Invalid characters in username or invalid length.\\n It should be betweeen 3 and 16 characters."}`))
		return
	}
	var p models.Profile
	p.ID = user_id
	p.Username = u.Username

	name, err := rt.db.GetNameById(user_id)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not get name from id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userProfile, err := rt.db.GetUserProfileByUsername(name)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not get user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var pp models.Profile
	pp.FromDatabase(userProfile, rt.db)

	oldProfileBytes, err := json.Marshal(&pp)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating JSON patch")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println(`{"op":"replace", "path": "/username", "value": "` + p.Username + `"}`)
	patchJSON := []byte(`[{"op":"replace", "path": "/username", "value": "` + p.Username + `"}]`)
	// fmt.Println(patchJSON)
	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding JSON patch")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	modified, err := patch.Apply(oldProfileBytes)
	if err != nil {
		log.Println("Error Applying patch json ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jsonProfile models.Profile
	err = json.Unmarshal(modified, &jsonProfile)
	if err != nil {
		log.Println("Error unmarshaling patch json ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Updating profile in the database
	_, err = rt.db.UpdateUserProfile(true, jsonProfile.ToDatabase())
	if err != nil {
		log.Println("Error UPDATING the database ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Send updated username in response
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(jsonProfile)
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The User ID in the path is a string
	user_id := ps.ByName("user_id")
	user_id = strings.TrimPrefix(user_id, ":user_id=")
	fmt.Println(user_id)
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
	// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
	name, err := rt.db.GetNameById(user_id)
	if err != nil {
		fmt.Println("couldn't get name from id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userProfile, err := rt.db.GetUserProfileByUsername(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var pp models.Profile
	pp.FromDatabase(userProfile)
	fmt.Println("bella", pp)
	oldProfileBytes, err := json.Marshal(&pp)
	if err != nil {
		fmt.Println("Error creating json patch", err.Error())
		return
	}
	fmt.Println(`{"op":"replace", "path": "/username", "value": "` + p.Username + `"}`)
	patchJSON := []byte(`[{"op":"replace", "path": "/username", "value": "` + p.Username + `"}]`)
	fmt.Println(patchJSON)
	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		fmt.Println("Error Decoding patch json ", err.Error())
		return
	}
	modified, err := patch.Apply(oldProfileBytes)
	if err != nil {
		fmt.Println("Error Applying patch json ", err.Error())
		return
	}

	var jsonProfile models.Profile
	fmt.Println("hey", jsonProfile)
	err = json.Unmarshal(modified, &jsonProfile)
	if err != nil {
		fmt.Println("Error unmarshaling patch json ", err.Error())
		return
	}
	// updating profile in the database
	_, err = rt.db.UpdateUserProfile(jsonProfile.ToDatabase())
	if err != nil {
		fmt.Println("Error UPDATING the database ", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(jsonProfile)
}

/*
curl -X PATCH http://localhost:3001/users/:user_id=2Idw46q6s1oYd9xk8Uy0iB9bc8g -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{"username":"paola"}'
curl -X POST "http://localhost:3001/session/" -H 'Content-Type: application/json' -d '{"username":"gianni"}'
curl "http://localhost:3001/users/?username=michi"
curl -X DELETE "http://localhost:3001/users/:user_id=5e9b1511-f78c-4163-a436-9e19fed724ad"
curl -X POST "http://localhost:3001/users/:user_id=5e9b1511-f78c-4163-a436-9e19fed724ad/photos/" -H 'Content-Type: application/json' -d '{"caption":"bella raga", "image":"https://www.sono_scema.it/michi.png"}'
curl "http://localhost:3001/users/:user_id=5e9b1511-f78c-4163-a436-9e19fed724ad/photos/"
curl -X DELETE "http://localhost:3001/photos/:photo_id=ef08d077-070f-435f-b4ae-b3929e759543"
curl -X GET "http://localhost:3001/photos/:photo_id=ef08d077-070f-435f-b4ae-b3929e759543"
*/

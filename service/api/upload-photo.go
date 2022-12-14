package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get userID from path
	// The User ID in the path is a string and coincides with the profile we are in
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

	// 3. Get photo from the request body
	// Decode information inserted by the user in the request body and check validity
	var media models.Photo
	err = json.NewDecoder(r.Body).Decode(&media)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !media.IsValid() {
		// Here we validated the fountain structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the fountain data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4. Upload the photo given by the user in the database
	media_db, err := rt.db.UploadPhoto(user_id, media.ToDatabase(rt.db))
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	media.FromDatabase(media_db, rt.db)
	// fmt.Println(media.LikesCount, media.Timestamp)
	// 5. Set the content type as application/json and encode the media written
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(media)
}

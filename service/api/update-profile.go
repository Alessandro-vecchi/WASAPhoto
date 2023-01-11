package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1. Get Id of the user whose profile is being updated
	user_id := rt.getPathParameter("user_id", ps)
	fmt.Println(user_id)
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2. Check if the user is authenticated
	// We want to allow only to logged users in their own profile to modify the profile.
	// Therefore the authentication token in the header should coincides with the id of the profile
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)
	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 3. Read new profile information from request body
	var p models.Profile
	// Get bio
	p.Bio = r.FormValue("bio")
	fmt.Println(r.Form)
	// Get Username
	p.Username = r.FormValue("username")
	// If username is not modified, get the old one
	if p.Username == "" {
		p.Username, _ = rt.db.GetNameById(user_id)
	} else if rt.db.CountStuffs("username", "profile", p.Username) > 0 {
		// User Already Exists
		ctx.Logger.WithError(err).WithField("username", p.Username).Error("Username already exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	} /*

		// 4. Get photo from the request body
		// Decode information inserted by the user in the request body
		photo, fileHeader, err := r.FormFile("image")
		if err != nil {
			ctx.Logger.WithError(err).Error("error: could not parse photo")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer photo.Close()
		buff := make([]byte, 512)
		_, err = photo.Read(buff)
		if err != nil {
			ctx.Logger.WithError(err).Error("error: could not read photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 5 - Check if the photo is valid
		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
			ctx.Logger.WithError(err).Error("error: The provided file format is not allowed. Please upload a JPEG,JPG or PNG image")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = photo.Seek(0, io.SeekStart)
		if err != nil {
			ctx.Logger.WithError(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 6 - Generate an ID that univoquely identifies the image
		rawPhotoId, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("failed to get UUID: %v", err)
		}
		log.Printf("generated Version 4 UUID: %v", rawPhotoId)
		photoId := rawPhotoId.String()

		// 7 - Save the photo in the images folder exploiting the image id
		f, err := os.Create(fmt.Sprintf("@/webui/src/assets/images/%s%s", photoId, filepath.Ext(fileHeader.Filename)))
		if err != nil {
			ctx.Logger.WithError(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, photo)
		if err != nil {
			ctx.Logger.WithError(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// 8 - Create picture url
		picURL := fmt.Sprintf("@/webui/src/assets/images/%s%s", photoId, filepath.Ext(fileHeader.Filename)) */

	p.ProfilePictureUrl = r.FormValue("profile_picture_url")

	// 9 - Update user_id
	// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
	p.ID = user_id

	// 10 - Save the profile information in the database
	_, err = rt.db.UpdateUserProfile(false, p.ToDatabase())

	if err != nil {
		ctx.Logger.WithError(err).WithField("user_id", user_id).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// w.WriteHeader(http.StatusNoContent)
	// Send the user profile to the user
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p)
}

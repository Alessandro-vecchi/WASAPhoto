package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get Id of the user whose profile is being updated
	user_id := rt.getPathParameter("user_id", ps)
	if user_id == "" {
		ctx.Logger.Error("wrong user_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
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
		_, _ = w.Write([]byte(`{"error": "User does not exist"}`))
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		return
	}
	// 3. Decode information inserted by the user in the request body
	err = r.ParseMultipartForm(32 << 20) // 32MB is the maximum file size
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error: could not parse form")
		_, _ = w.Write([]byte(`{"error": "32MB is the maximum file size"}`))
		return
	}

	// 4. Read new profile information from request body
	var p models.Profile
	// Get Username
	var u models.Username
	u.Username = r.FormValue("username")
	// If modified username coincides with an existing one different from my old one, send error to the user
	oldName, _ := rt.db.GetNameById(user_id)
	if u.Username != oldName && rt.db.CountStuffs("username", "profile", u.Username) > 0 {
		// User Already Exists
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "This username already exists. Please choose another username."}`))
		return
	} else if !u.IsValid() {
		// Username is invalid
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "Invalid characters in username or username too short. Its length should be betweeen 3 and 16 characters."}`))
		return
	}
	p.Username = u.Username
	// Get bio
	p.Bio = r.FormValue("bio")
	if !(len(p.Bio) <= 200 && models.BioRx.MatchString(p.Bio)) {
		// Username is invalid
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "Invalid characters inserted in the bio or bio too long. Maximum 200 characters long."}`))
		return
	}
	p.ProfilePictureUrl, _ = rt.db.GetProfilePic(user_id)
	// Get photo from the request body
	var file_name string
	_, _, err = r.FormFile("image")
	// log.Printf("err: %v, err2: %v", err, errors.New("http: no such file"))
	if fmt.Sprintf("%v", err) == "http: no such file" {
		log.Printf("Using old profile picture")
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error: could not parse photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		file_name, err = createImage(r, ctx)
		if errors.Is(err, errors.New("bad image format")) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "The provided file format is not allowed. Please upload a JPEG,JPG or PNG image."}`))
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 8 - Create picture url
		log.Printf("image path name: %s", file_name)
		p.ProfilePictureUrl = file_name
	}
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

func createImage(r *http.Request, ctx reqcontext.RequestContext) (string, error) {
	photo, fileHeader, err := r.FormFile("image")
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not parse photo")
		return "", err
	}
	defer photo.Close()
	buff := make([]byte, 512)
	_, err = photo.Read(buff)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not read photo")
		return "", err
	}

	// 5 - Check if the photo is valid
	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
		return "", errors.New("bad image format")
	}
	_, err = photo.Seek(0, io.SeekStart)
	if err != nil {
		ctx.Logger.WithError(err)
		return "", err
	}
	// 6 - Generate an ID that univoquely identifies the image
	rawPhotoId, err := uuid.NewV4()
	if err != nil {
		ctx.Logger.WithError(err).Error("failed to get UUID")
		return "", err
	}
	log.Printf("generated Version 4 photoID: %v", rawPhotoId)
	photoId := rawPhotoId.String()

	// 7 - Save the photo in the images folder exploiting the image id

	image_directory := "tmp"
	folder_name := "images"
	file_name := fmt.Sprintf("%s%s", photoId, filepath.Ext(fileHeader.Filename))
	path := filepath.Join(image_directory, folder_name, file_name)
	log.Printf("Current directory: %v, Folder: %v, Filename: %v, path: %v,", image_directory, folder_name, file_name, path)

	f, err := os.Create(path)
	if err != nil {
		ctx.Logger.WithError(err)
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, photo)
	if err != nil {
		ctx.Logger.WithError(err)
		return "", err
	}
	return file_name, nil
}

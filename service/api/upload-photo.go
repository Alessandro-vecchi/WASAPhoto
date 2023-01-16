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
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)

	err := checkUserIdentity(authtoken, user_id, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {

		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// 3. Decode information inserted by the user in the request body
	err = r.ParseMultipartForm(32 << 20) // 32MB is the maximum file size
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not parse form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4. Read new photo info from request body
	// Get caption
	caption := r.FormValue("caption")
	// Get photo from the request body
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
	log.Printf("generated Version 4 photoID: %v", rawPhotoId)
	photoId := rawPhotoId.String()

	// 7 - Save the photo in the images folder exploiting the image id
	f, err := os.Create(fmt.Sprintf("./webui/src/assets/images/%s%s", photoId, filepath.Ext(fileHeader.Filename)))

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
	picURL := fmt.Sprintf("./src/assets/images/%s%s", photoId, filepath.Ext(fileHeader.Filename))
	log.Printf("image path name: %s", picURL)

	// 9 - create media object
	var media models.Photo
	media.PhotoId = photoId
	media.Caption = caption
	media.Image = picURL

	// 4. Upload the photo given by the user in the database
	media_db, err := rt.db.UploadPhoto(user_id, media.ToDatabase(rt.db))
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("photo can't be uploaded in the database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	media.FromDatabase(media_db, rt.db)
	// fmt.Println(media.LikesCount, media.Timestamp)
	// 5. Set the content type as application/json and encode the media written
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(media)
}

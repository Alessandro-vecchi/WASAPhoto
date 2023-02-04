package api

import (
	_ "image/jpeg"
	"io"
	"log"
	"path/filepath"

	_ "image/png"
	"net/http"
	"os"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// get photo from photos folder
func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var image_file_name string
	if r.URL.Query().Has("image_name") {
		image_file_name = r.URL.Query().Get("image_name")

	} else {
		// No image field founded
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "query has no field image_name"}`))
		return
	}

	// Obtain the path for the image in the server
	image_directory := "/tmp"
	folder_name := "images"
	path := filepath.Join(image_directory, folder_name, image_file_name)
	log.Printf("Current directory: %v, Folder: %v, Filename: %v, path: %v,", image_directory, folder_name, image_file_name, path)

	// Open the image
	img, err := os.Open(path)
	if err != nil {
		// error opening file
		ctx.Logger.WithError(err).Error("error: could not open photo file")
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer img.Close()

	// 2 - Send the image back to the front
	w.Header().Set("Content-Type", "blob")
	_, err = io.Copy(w, img)
	if err != nil {
		// error copying image
		ctx.Logger.WithError(err).Error("error: could not copy photo")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

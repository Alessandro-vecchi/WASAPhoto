package api

import (
	_ "image/jpeg"
	"io"

	_ "image/png"
	"net/http"
	"os"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Get image from images folder
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - Open file
	img, err := os.Open("." + r.URL.Path)
	if err != nil {
		// error opening file
		ctx.Logger.Error("error: could not open photo file")
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer img.Close()

	// 2 - prepare response
	w.Header().Set("Content-Type", "image/jpeg")
	_, err = io.Copy(w, img)
	if err != nil {
		// error copying image
		ctx.Logger.Error("error: could not copy photo")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

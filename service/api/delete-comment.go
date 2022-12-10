package api

import (
	"net/http"
	"strings"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The user ID in the path is a string. Let's parse it.
	photo_id := ps.ByName("photo_id")
	photo_id = strings.TrimPrefix(photo_id, ":photo_id=")
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

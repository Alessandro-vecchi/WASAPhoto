package api

/*
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) patchUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	profileBytes, err := json.Marshal(&Profile{})
	if err != nil {
		fmt.Println("Error creating patch json ", err.Error())
		return
	}
	fmt.Println(string(profileBytes))
	patchJSON, _ := ioutil.ReadAll(r.Body)
	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		fmt.Println("Error Decoding patch json ", err.Error())
		return
	}
	modified, err := patch.Apply(profileBytes)
	if err != nil {
		fmt.Println("Error applying patch json ", err.Error())
		return
	}
	w.Write(modified)
	w.Header().Set("Content-Type", "application/json")
}
*/

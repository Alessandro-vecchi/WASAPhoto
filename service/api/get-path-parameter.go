package api

import (
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPathParameter(path string, ps httprouter.Params) string {

	raw_path_param := ps.ByName(path)
	path_param := strings.TrimPrefix(raw_path_param, ":"+path+"=")
	if path_param == "" {
		return ""
	}
	return path_param
}

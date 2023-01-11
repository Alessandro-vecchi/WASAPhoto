package api

import (
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPathParameter(path_parameter_name string, ps httprouter.Params) string {

	path_param := ps.ByName(path_parameter_name)
	if models.IsValidUUID(path_param) {
		return path_param
	}
	return ""
}

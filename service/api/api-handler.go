package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// Do login
	rt.router.POST("/session/", rt.wrap(rt.doLogin))

	// Get user profile by username
	rt.router.GET("/users/", rt.wrap(rt.getUserProfile))

	// Update username
	//rt.router.PATCH("/users/:user_id/", rt.wrap(rt.patchUsername))

	// Update username
	rt.router.PUT("/users/:user_id/", rt.wrap(rt.setMyUserName))

	// Delete user profile
	rt.router.DELETE("/users/:user_id/", rt.wrap(rt.deleteUserProfile))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

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
	rt.router.PATCH("/users/:user_id", rt.wrap(rt.setMyUserName))

	// Update username
	rt.router.PUT("/users/:user_id", rt.wrap(rt.updateProfile))

	// Delete user profile
	rt.router.DELETE("/users/:user_id", rt.wrap(rt.deleteUserProfile))

	// Upload a photo
	rt.router.POST("/users/:user_id/photos/", rt.wrap(rt.uploadPhoto))

	// Retrieve collection of photos of a user
	rt.router.GET("/users/:user_id/photos/", rt.wrap(rt.getUserPhotos))

	// Get a single user photo
	rt.router.GET("/photos/:photo_id", rt.wrap(rt.getUserPhoto))

	// Delete a single user photo
	rt.router.DELETE("/photos/:photo_id", rt.wrap(rt.deletePhoto))

	// Upload a comment
	rt.router.POST("/photos/:photo_id/comments/", rt.wrap(rt.commentPhoto))

	// Delete a comment
	rt.router.DELETE("/comments/:comments_id", rt.wrap(rt.uncommentPhoto))

	// Follow a user
	rt.router.PUT("/users/:user_id/followers/:followers_id", rt.wrap(rt.followUser))

	// Unfollow a user
	rt.router.DELETE("/users/:user_id/followers/:followers_id", rt.wrap(rt.unfollowUser))

	// Get list followers
	rt.router.GET("/users/:user_id/followers/", rt.wrap(rt.getFollowers))

	// Get list followed
	rt.router.GET("/users/:user_id/following/", rt.wrap(rt.getFollowed))

	// Put a like to a photo
	rt.router.PUT("/photos/:photo_id/likes/:likes_id", rt.wrap(rt.likePhoto))

	// Unlike a photo
	rt.router.DELETE("/photos/:photo_id/likes/:likes_id", rt.wrap(rt.unlikePhoto))

	// Get list of the users that added a like
	rt.router.GET("/photos/:photo_id/likes/", rt.wrap(rt.getLikes))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

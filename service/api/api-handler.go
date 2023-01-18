package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	/* LOGIN */
	rt.router.POST("/session/", rt.wrap(rt.doLogin))

	/* USER PROFILE */
	// Get user profile by username
	rt.router.GET("/users/", rt.wrap(rt.getUserProfile))
	// Update username
	rt.router.PATCH("/users/:user_id", rt.wrap(rt.setMyUserName))
	// Update user profile
	rt.router.PUT("/users/:user_id", rt.wrap(rt.updateProfile))
	// Delete user profile
	rt.router.DELETE("/users/:user_id", rt.wrap(rt.deleteUserProfile))

	/* PHOTO */
	// Upload a photo
	rt.router.POST("/users/:user_id/photos/", rt.wrap(rt.uploadPhoto))
	// Get a single user photo
	rt.router.GET("/photos/:photo_id", rt.wrap(rt.getUserPhoto))
	// Retrieve collection of photos of a user
	rt.router.GET("/users/:user_id/photos/", rt.wrap(rt.getUserPhotos))
	// Delete a single user photo
	rt.router.DELETE("/photos/:photo_id", rt.wrap(rt.deletePhoto))
	// Get Image
	rt.router.GET("/images/", rt.wrap(rt.getImage))

	/* COMMENTS */
	// Upload a comment
	rt.router.POST("/photos/:photo_id/comments/", rt.wrap(rt.commentPhoto))
	// Retrieve list of comments under a photo
	rt.router.GET("/photos/:photo_id/comments/", rt.wrap(rt.getPhotoComments))
	// Modify an own comment
	rt.router.PUT("/comments/:comment_id", rt.wrap(rt.modifyComment))
	// Delete a comment
	rt.router.DELETE("/comments/:comment_id", rt.wrap(rt.uncommentPhoto))

	/* LIKES */
	// Put a like to a photo
	rt.router.PUT("/photos/:photo_id/likes/:like_id", rt.wrap(rt.likePhoto))
	// Get list of the users that added a like
	rt.router.GET("/photos/:photo_id/likes/", rt.wrap(rt.getLikes))
	// Unlike a photo
	rt.router.DELETE("/photos/:photo_id/likes/:like_id", rt.wrap(rt.unlikePhoto))

	/* FOLLOWERS */
	// Follow a user
	rt.router.PUT("/users/:user_id/followers/:follower_id", rt.wrap(rt.followUser))
	// Get list followers
	rt.router.GET("/users/:user_id/followers/", rt.wrap(rt.getFollowers))
	// Get list following
	rt.router.GET("/users/:user_id/following/", rt.wrap(rt.getFollowed))
	// Unfollow a user
	rt.router.DELETE("/users/:user_id/followers/:follower_id", rt.wrap(rt.unfollowUser))

	/* BANS */
	// Ban a user
	rt.router.PUT("/users/:user_id/bans/:ban_id", rt.wrap(rt.banUser))
	// Get list of banned users
	rt.router.GET("/users/:user_id/bans/", rt.wrap(rt.getBannedUsers))
	// Unban a user
	rt.router.DELETE("/users/:user_id/bans/:ban_id", rt.wrap(rt.unbanUser))

	/* STREAM */
	// Get stream of the user
	rt.router.GET("/users/:user_id/stream/", rt.wrap(rt.getMyStream))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

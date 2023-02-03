package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/Alessandro-vecchi/WASAPhoto/service/api/models"
	"github.com/Alessandro-vecchi/WASAPhoto/service/api/reqcontext"
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1. Get user ID of user A from path
	// The User ID in the path is a string and coincides with the profile we watching
	user_id_A := rt.getPathParameter("user_id", ps)
	if user_id_A == "" {
		ctx.Logger.Error("wrong user_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 2. Get user ID of user B from path
	// It coincides with the user that want to follow
	user_id_B := rt.getPathParameter("follower_id", ps)
	if user_id_B == "" {
		ctx.Logger.Error("wrong follower_id path parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. Check that the user is not unfollowing himself
	if models.AreTheSame(user_id_A, user_id_B) {
		// A user can't unfollow himself
		ctx.Logger.WithError(database.ErrUserCantFollowHimself).Error("a user can't unfollow himself ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4. Check if the user B is authenticated
	// We want to allow only to logged users to put likes.
	// Therefore the authentication token in the header should coincide with the id of the user who is liking the photo
	authtoken := r.Header.Get("Authorization")
	log.Printf("The authentication token in the header is: %v", authtoken)
	// If I ban a user and he's following me, his follow should be removed from the database
	user := user_id_A
	if user != authtoken {
		user = user_id_B
	}
	err := checkUserIdentity(authtoken, user, rt.db)
	if errors.Is(err, database.ErrUserNotExists) {
		_, _ = w.Write([]byte(`{"error": "User does not exist"}`))
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAuthenticationFailed) {
		_, _ = w.Write([]byte(`{"error": "You are not authenticated"}`))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = rt.db.UnfollowUser(user_id_A, user_id_B)
	if errors.Is(err, database.ErrFollowerNotPresent) {
		// User B wasn't following user A, reject the action indicating an error on the client side.
		_, _ = w.Write([]byte(`{"error": "You are not following the user"}`))
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the fountain that triggered the error.
		ctx.Logger.WithError(err).WithField("id", user_id_B).Error("can't unfollow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

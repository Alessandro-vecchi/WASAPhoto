package api

import (
	"errors"
	"log"

	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
)

func checkUserIdentity(authtoken string, user_id string, db database.AppDatabase) error {

	_, err := db.GetNameById(user_id)
	if errors.Is(err, database.ErrUserNotExists) {
		return database.ErrUserNotExists
	}
	// user exists in the database
	if authtoken == "" || authtoken != user_id {
		log.Printf("Authorization token %v doesn't match user_id %v: Unauthorized", authtoken, user_id)
		return database.ErrAuthenticationFailed
	}
	return nil

}

func contains(list []database.Short_profile_db, filter string) bool {
	for _, item := range list {
		if item.Username == filter {
			return true
		}
	}
	return false
}

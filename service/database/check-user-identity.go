package database

import (
	"errors"
	"fmt"
)

func (db *appdbimpl) CheckUserIdentity(authtoken string, user_id string) error {

	_, err := db.GetNameById(user_id)
	if errors.Is(err, ErrUserNotExists) {
		return ErrUserNotExists
	}
	// user exists in the database
	if authtoken != user_id {
		fmt.Printf("authtoken %v doesn't match user_id %v: Unauthorized", authtoken, user_id)
		return ErrAuthenticationFailed
	}
	return nil

}

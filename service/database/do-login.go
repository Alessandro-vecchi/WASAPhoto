package database

import (
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func (db *appdbimpl) DoLogin(user_name string) (string, error) {

	p, err := db.GetUserProfileByUsername(user_name)

	if err == nil {
		// profile already exists, returning user id of the existing profile
		return p.ID, nil
	} else if err == ErrUserNotExists {
		// user does not exist, creating a new username
		u4, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("failed to get UUID: %v", err)
		}
		log.Printf("generated Version 4 UUID: %v", u4)
		uid := u4.String()
		_, err = db.c.Exec(`INSERT INTO profile (user_id, username) VALUES (?,?)`, uid, user_name)
		fmt.Println(err)
		if err != nil {
			return "", fmt.Errorf("userID could not be created: %w", err)
		}
		return uid, nil
	}
	// scan error
	return "", fmt.Errorf("error while scanning the profile: %w", err)
}

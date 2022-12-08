package database

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

func (db *appdbimpl) DoLogin(user_name string) (string, error) {

	p, err := db.GetUserProfileByUsername(user_name)
	fmt.Println(err)
	if err == nil {
		// profile already exists, returning user id of the existing profile
		return p.ID, nil
	} else if err == ErrUserNotExists {
		// user does not exist, creating a new username
		row_uid := ksuid.New()
		uuid := row_uid.String()
		fmt.Println("bella" + uuid)
		_, err = db.c.Exec(`INSERT INTO profile (user_id, username) VALUES (?,?)`, uuid, user_name)
		fmt.Println(err)
		if err != nil {
			return "", fmt.Errorf("userID could not be created: %w", err)
		}
		return uuid, nil
	}
	// scan error
	return "", fmt.Errorf("error while scanning the profile: %w", err)
}

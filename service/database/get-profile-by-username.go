package database

import (
	"log"
)

func (db *appdbimpl) GetUserProfileByUsername(username string) (Profile_db, error) {

	const query = `
SELECT *
FROM profile
WHERE username = ?`

	var ret Profile_db
	// Issue the query, using the username as filter
	err := db.c.QueryRow(query, username).Scan(&ret.ID, &ret.Username, &ret.ProfilePictureUrl, &ret.Bio)
	if err != nil {
		log.Print("error while querying the profile: %w", err)
		return Profile_db{}, ErrUserNotExists
	}
	return ret, nil

}

package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) FollowUser(userId_A string, followerId_B string) error {

	// check if user B is already following user A
	const query = `
SELECT *
FROM follow
WHERE user_id = ? AND follower_id = ?`

	err := db.c.QueryRow(query, userId_A, followerId_B).Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		// user B already follows user A
		return ErrFollowerAlreadyPresent
	}
	// user B doesn't follow user A
	_, err = db.c.Exec(`INSERT INTO follow (follower_id, followed_id) VALUES (?,?)`, followerId_B, userId_A)
	if err != nil {
		return fmt.Errorf("error when following a new user: %w", err)
	}
	return nil
}

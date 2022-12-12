package database

import "fmt"

func (db *appdbimpl) UnfollowUser(followedId_A string, followerId_B string) error {

	const query = `
		DELETE
		FROM follow
		WHERE follower_id = ? AND followed_id = ?`
	res, err := db.c.Exec(query, followerId_B, followedId_A)
	if err != nil {
		return fmt.Errorf("error when unfollowing an user: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then user B wasn't following user A
		return ErrFollowerNotPresent
	}
	// succesfully unfollowed the user
	return nil
}

package database

import "fmt"

func (db *appdbimpl) UnfollowUser(userId_A string, followerId_B string) error {

	const query = `
		DELETE
		FROM follow
		WHERE user_id = ? AND follower_id = ?`
	res, err := db.c.Exec(query, userId_A, followerId_B)
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
	// successfully unfollowed the user
	return nil
}

package database

import "fmt"

func (db *appdbimpl) UnbanUser(banned_user string, banner_user string) error {

	const query = `
		DELETE
		FROM ban
		WHERE banner_id = ? AND banned_id = ?`

	res, err := db.c.Exec(query, banner_user, banned_user)
	if err != nil {
		return fmt.Errorf("error when unbanning an user: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then user B did not ban user A
		return ErrBanNotPresent
	}
	// succesfully unfbanned the user
	return nil
}

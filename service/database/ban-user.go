package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) BanUser(banned_user string, banner_user string) error {

	// check if user B is already following user A
	const query = `
		SELECT *
		FROM ban
		WHERE banner_id = ? AND banned_id = ?`

	err := db.c.QueryRow(query, banner_user, banned_user).Scan()
	fmt.Println(err)
	if !errors.Is(err, sql.ErrNoRows) {
		// user B already follows user A
		return ErrBanAlreadyPresent
	}
	// user B doesn't follow user A
	_, err = db.c.Exec(`INSERT INTO ban (banner_id, banned_id) VALUES (?,?)`, banner_user, banned_user)
	if err != nil {
		return fmt.Errorf("error when banning a new user: %w", err)
	}
	return nil
}

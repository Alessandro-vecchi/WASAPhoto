package database

import (
	"fmt"
)

func (db *appdbimpl) DeleteUserProfile(userID string) error {

	// delete user profile
	table_name := "profile"
	query := `DELETE FROM ` + table_name + ` WHERE user_id =?`
	err := db.deleteTable(query, table_name, userID)
	if err != nil {
		//
		return err
	}
	return nil
}

func (db *appdbimpl) deleteTable(query string, table_name string, user_id string) error {

	res, err := db.c.Exec(query, user_id)
	if err != nil {
		return fmt.Errorf("could not delete the "+table_name+" table. error: %w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the id didn't exist
		return ErrUserNotExists
	}
	// table succesfully deleted
	return nil
}

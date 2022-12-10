package database

import "fmt"

func (db *appdbimpl) DeleteUserProfile(userID string) error {
	res, err := db.c.Exec("DELETE FROM profile WHERE user_id =?", userID)
	if err != nil {
		//
		return fmt.Errorf("could not delete the user profile. error: %v", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the user didn't exist
		return ErrUserNotExists
	}

	_, err = db.c.Exec("DELETE FROM photos WHERE user_id =?", userID)
	if err != nil {
		//
		return fmt.Errorf("could not delete the user media. error: %v", err)
	}

	return nil
}

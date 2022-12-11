package database

import "fmt"

func (db *appdbimpl) DeletePhoto(photoId string) error {
	res, err := db.c.Exec("DELETE FROM photos WHERE photo_id =?", photoId)
	if err != nil {
		// error deleting the photo
		return fmt.Errorf("error deleting the photo: %v", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the photo didn't exist
		return ErrPhotoNotExists
	}
	return nil
}

package database

import "fmt"

func (db *appdbimpl) UnlikePhoto(photoId string, userId string) error {

	const query = `
		DELETE
		FROM likes
		WHERE photo_id = ? AND liker_id = ?`
	res, err := db.c.Exec(query, photoId, userId)
	if err != nil {
		return fmt.Errorf("error when unliking a photo: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the like was not present
		return ErrLikeNotPresent
	}
	// successfully removed the like
	return nil
}

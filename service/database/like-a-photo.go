package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) LikePhoto(photoId string, userId string) error {

	// check if user already liked photo with given id
	const query = `
SELECT *
FROM likes
WHERE photo_id = ? AND user_id = ?`

	err := db.c.QueryRow(query, photoId, userId).Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		// user already liked photo

		return ErrLikeAlreadyPut
	}
	// inserting like in the database
	_, err = db.c.Exec(`INSERT INTO likes (photo_id, user_id) VALUES (?,?)`, photoId, userId)
	if err != nil {
		return fmt.Errorf("error when liking a photo: %w", err)
	}
	return nil
}

package database

import (
	"fmt"
	"time"
)

func (db *appdbimpl) UploadPhoto(user_id string, p Photo_db) (Photo_db, error) {

	// Generate the timestamp
	p.Timestamp = time.Now().Format(time.RFC3339)
	p.UserId = user_id
	_, err := db.c.Exec(`INSERT INTO photos (user_id, photo_id, timestamp, caption, image) VALUES (?,?,?,?,?)`, p.UserId, p.PhotoId, p.Timestamp, p.Caption, p.Image)
	if err != nil {
		return Photo_db{}, fmt.Errorf("photo could not be created: %w", err)
	}
	return p, nil

}

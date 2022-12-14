package database

import (
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

func (db *appdbimpl) UploadPhoto(userId string, p Photo_db) (Photo_db, error) {

	// photo does not exist, creating a new username
	rawPhotoId, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to get UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID: %v", rawPhotoId)
	photoId := rawPhotoId.String()
	p.Timestamp = time.Now().Format(time.RFC3339)
	_, err = db.c.Exec(`INSERT INTO photos (user_id, photo_id, timestamp, caption, image) VALUES (?,?,?,?,?)`, userId, photoId, p.Timestamp, p.Caption, p.Image)
	if err != nil {
		return Photo_db{}, fmt.Errorf("photo could not be created: %w", err)
	}
	return p, nil

}

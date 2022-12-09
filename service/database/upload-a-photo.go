package database

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

func (db *appdbimpl) UploadPhoto(userId string, p Photo_db) (Photo_db, error) {

	// user does not exist, creating a new username
	rawPhotoId := ksuid.New()
	photoId := rawPhotoId.String()
	p.Timestamp = time.Now().Format(time.RFC3339)
	fmt.Println(p.Timestamp)
	_, err := db.c.Exec(`INSERT INTO photos (user_id, photoId, timestamp, likesCount, commentsCount, caption, image) VALUES (?,?,?,?,?,?,?)`, userId, photoId, p.Timestamp, 0, 0, p.Caption, p.Image)
	if err != nil {
		return Photo_db{}, fmt.Errorf("photo could not be created: %w", err)
	}
	return p, nil

}

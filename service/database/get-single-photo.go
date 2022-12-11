package database

import "fmt"

func (db *appdbimpl) GetUserPhoto(photoId string) (Photo_db, error) {

	var p Photo_db
	const query = `
SELECT *
FROM photos
WHERE photo_id = ?`

	err := db.c.QueryRow(query, photoId).Scan(&p.UserId, &p.PhotoId, &p.Timestamp, &p.Caption, &p.Image)
	fmt.Println(p, err)
	if err != nil {

		return Photo_db{}, ErrPhotoNotExists
	}
	return p, nil
}

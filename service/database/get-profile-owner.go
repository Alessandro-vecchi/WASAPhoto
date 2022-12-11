package database

func (db *appdbimpl) GetProfileOwner(photo_id string) (string, error) {

	var uid string
	const query = `
		SELECT user_id
		FROM photos
        WHERE photo_id = ?`

	err := db.c.QueryRow(query, photo_id).Scan(&uid)
	if err != nil {

		return "", ErrPhotoNotExists
	}
	return uid, nil
}

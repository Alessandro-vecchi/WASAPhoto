package database

func (db *appdbimpl) GetUserPhoto(user_id string, photoId string) (Photo_db, error) {

	var name string
	const query = `
SELECT *
FROM photo
WHERE user_id = ?`

	err := db.c.QueryRow(query, user_id).Scan(&name)
	if err != nil {

		return Photo_db{}, ErrUserNotExists
	}
	return Photo_db{}, nil
}

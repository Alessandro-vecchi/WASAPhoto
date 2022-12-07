package database

func (db *appdbimpl) SetMyUserName(p Profile_db) (Profile_db, error) {
	var p_empty Profile_db
	const query = `
SELECT username, picturesCount, followersCount, followsCount, profilePictureUrl, bio
FROM Profile
WHERE username = ?`

	_, err := db.c.Exec(`UPDATE Profile SET username =? WHERE user_id=?`, p.Username, p.ID)

	if err != nil {
		return old, err
	}

	return p, nil
}

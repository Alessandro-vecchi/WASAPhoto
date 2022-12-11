package database

func (db *appdbimpl) UpdateUserProfile(p Profile_db) (Profile_db, error) {
	id := p.ID
	const query = `
SELECT username, picturesCount, followersCount, followsCount, profilePictureUrl, bio
FROM profile
WHERE user_id = ?`
	var old Profile_db
	user := db.c.QueryRow(query, id)

	err := user.Scan(&old.Username, &old.ProfilePictureUrl, &old.Bio)
	if err != nil {
		return Profile_db{}, err
	}
	_, err = db.c.Exec(`UPDATE profile SET username =? WHERE user_id=?`, p.Username, p.ID)

	if err != nil {
		return old, err
	}

	return p, nil
}

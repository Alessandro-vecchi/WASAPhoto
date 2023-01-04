package database

func (db *appdbimpl) UpdateUserProfile(isPatch bool, p Profile_db) (Profile_db, error) {

	const query = `
		SELECT username, profilePictureUrl, bio
		FROM profile
		WHERE user_id = ?`

	var old Profile_db
	user := db.c.QueryRow(query, p.ID)

	err := user.Scan(&old.Username, &old.ProfilePictureUrl, &old.Bio)
	if err != nil {
		return Profile_db{}, err
	}
	if isPatch {
		_, err = db.c.Exec(`UPDATE profile SET username = ? WHERE user_id = ?`, p.Username, p.ID)
	} else {
		if p.Username == "" {
			p.Username = old.Username
		}
		_, err = db.c.Exec(`UPDATE profile SET username = ?, profilePictureUrl = ?, bio = ?  WHERE user_id = ?;`, p.Username, p.ProfilePictureUrl, p.Bio, p.ID)

	}
	if err != nil {
		// if there is an error return old profile
		return old, err
	}

	return p, nil
}

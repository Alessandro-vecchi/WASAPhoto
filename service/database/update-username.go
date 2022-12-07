package database

func (db *appdbimpl) SetMyUserName(p Profile_db) (Profile_db, error) {
	id := p.ID
	const query = `
SELECT username, picturesCount, followersCount, followsCount, profilePictureUrl, bio
FROM profile
WHERE uuid = ?`
	var old Profile_db
	user, _ := db.c.Query(query, id)
	defer func() { _ = user.Close() }()

	err := user.Scan(&old.Username, &old.PicturesCount, &old.FollowersCount, &old.FollowsCount, &old.ProfilePictureUrl, &old.Bio)
	if err != nil {
		return Profile_db{}, err
	}
	_, err = db.c.Exec(`UPDATE Profile SET username =? WHERE user_id=?`, p.Username, p.ID)

	if err != nil {
		return old, err
	}

	return p, nil
}

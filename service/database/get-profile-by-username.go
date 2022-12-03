package database

func (db *appdbimpl) GetUserProfileByUsername(username string) (Profile, error) {

	const query = `
SELECT username, picturesCount, followersCount, followsCount, profilePictureUrl, bio
FROM Profile
WHERE username = ?`

	var ret Profile
	// Issue the query, using the username as filter
	userPage, err := db.c.Query(query)
	if err != nil {

		return ret, ErrUserNotExists
	}

	defer func() { _ = userPage.Close() }()

	err = userPage.Scan(&ret.Username, &ret.PicturesCount, &ret.FollowersCount, &ret.FollowsCount, &ret.ProfilePictureUrl, &ret.Bio)
	if err != nil {
		var p Profile
		return p, err
	}
	return ret, nil

}

package database

import "fmt"

func (db *appdbimpl) GetUserProfileByUsername(username string) (Profile_db, error) {

	const query = `
SELECT username, picturesCount, followersCount, followsCount, profilePictureUrl, bio
FROM profiles
WHERE username = ?`

	var ret Profile_db
	// Issue the query, using the username as filter
	userPage, err := db.c.Query(query, username)
	if err != nil {

		return Profile_db{}, err
	}
	// why not defer userPage.Close()
	defer func() { _ = userPage.Close() }()

	// updating the profile in place
	//fmt.Println(userPage.Err())
	for userPage.Next() {
		err = userPage.Scan(&ret.Username, &ret.PicturesCount, &ret.FollowersCount, &ret.FollowsCount, &ret.ProfilePictureUrl, &ret.Bio)
		fmt.Println("blo")
		if err != nil {
			return Profile_db{}, fmt.Errorf("error while scanning the profile: %w", err)
		}

	}

	if !userPage.Next() {
		return Profile_db{}, ErrUserNotExists
	}

	return ret, nil

}

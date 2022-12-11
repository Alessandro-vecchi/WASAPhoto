package database

import "fmt"

func (db *appdbimpl) GetUserProfileByUsername(username string) (Profile_db, error) {

	const query = `
SELECT *
FROM profile
WHERE username = ?`

	var ret Profile_db
	// Issue the query, using the username as filter
	userPage, err := db.c.Query(query, username)
	if err != nil {

		return Profile_db{}, err
	}
	// why not defer userPage.Close()
	defer func() { _ = userPage.Close() }()

	//fmt.Println(userPage.Err())
	res := userPage.Next()
	if !res {
		return Profile_db{}, ErrUserNotExists
	}
	err = userPage.Scan(&ret.ID, &ret.Username, &ret.ProfilePictureUrl, &ret.Bio)
	if err != nil {
		return Profile_db{}, fmt.Errorf("error while scanning the profile: %w", err)
	}

	return ret, nil

}

package database

import "fmt"

// get the list of the users that is followed by a user
func (db *appdbimpl) GetFollowing(follower_id string) ([]Short_profile_db, error) {

	const query = `
	SELECT profilePictureUrl, user_id 
	FROM profile, ( SELECT followed_id FROM follow WHERE follower_id =? ) as following
	WHERE profile.user_id = following.followed_id;`

	var short_profile []Short_profile_db

	rows, err := db.c.Query(query, follower_id)
	if err != nil {
		return []Short_profile_db{}, fmt.Errorf("error fetching following: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all followers in the result set
	for rows.Next() {
		var s_p Short_profile_db
		var followed_id string
		err = rows.Scan(&s_p.ProfilePictureUrl, &followed_id)
		if err != nil {
			return []Short_profile_db{}, fmt.Errorf("error scanning following: %w", err)
		}
		s_p.Username, err = db.GetNameById(followed_id)
		if err != nil {
			return []Short_profile_db{}, err
		}
		short_profile = append(short_profile, s_p)
	}
	if err = rows.Err(); err != nil {
		return []Short_profile_db{}, fmt.Errorf("error encountered during iteration: %w", err)
	}
	// successfully retrieved following
	return short_profile, nil

}

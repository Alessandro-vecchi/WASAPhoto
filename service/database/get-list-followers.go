package database

import "fmt"

// get the list of the users that follows a user
func (db *appdbimpl) GetFollowers(followed_id string) ([]Short_profile_db, error) {

	const query = `
	SELECT profilePictureUrl, user_id 
	FROM profile, ( SELECT follower_id FROM follow WHERE followed_id =? ) as followers 
	WHERE profile.user_id = followers.follower_id;`

	var short_profile []Short_profile_db

	rows, err := db.c.Query(query, followed_id)
	if err != nil {
		return []Short_profile_db{}, fmt.Errorf("error fetching followers: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all followers in the result set
	for rows.Next() {
		var s_p Short_profile_db
		var follower_id string
		err = rows.Scan(&s_p.ProfilePictureUrl, &follower_id)
		if err != nil {
			return []Short_profile_db{}, fmt.Errorf("error scanning followers: %w", err)
		}
		s_p.Username, err = db.GetNameById(follower_id)
		if err != nil {
			return []Short_profile_db{}, err
		}
		short_profile = append(short_profile, s_p)
	}
	if err = rows.Err(); err != nil {
		return []Short_profile_db{}, fmt.Errorf("error encountered during iteration: %w", err)
	}
	// successfully retrieved followers
	return short_profile, nil

}

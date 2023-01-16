package database

import "fmt"

// get list of the users that are banned from another user
func (db *appdbimpl) GetBannedUsers(banner_id string) ([]Short_profile_db, error) {

	const query = `
	SELECT profilePictureUrl, user_id 
	FROM profile, ( SELECT banned_id FROM ban WHERE banner_id = ? ) as banned 
	WHERE profile.user_id = banned.banned_id;`

	var banned_users []Short_profile_db

	rows, err := db.c.Query(query, banner_id)
	if err != nil {
		return []Short_profile_db{}, fmt.Errorf("error fetching banned users: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all banned users in the result set
	for rows.Next() {
		var short_prof Short_profile_db
		var banned_id string
		err = rows.Scan(&short_prof.ProfilePictureUrl, &banned_id)
		if err != nil {
			return []Short_profile_db{}, fmt.Errorf("error scanning banned users: %w", err)
		}
		short_prof.Username, err = db.GetNameById(banned_id)
		if err != nil {
			return []Short_profile_db{}, err
		}
		banned_users = append(banned_users, short_prof)
	}
	if err = rows.Err(); err != nil {
		return []Short_profile_db{}, fmt.Errorf("error encountered during iteration: %w", err)
	}
	// successfully retrieved banned users
	return banned_users, nil

}

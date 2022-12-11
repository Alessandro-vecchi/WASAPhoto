package database

import "fmt"

// get the list of the users that follows a user
func (db *appdbimpl) GetFollowing(follower_id string) ([]string, error) {

	const query = `
SELECT followed_id
FROM follow
WHERE follower_id =?`

	var followed []string

	rows, err := db.c.Query(query, follower_id)
	if err != nil {
		return nil, fmt.Errorf("error fetching following: %v", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all followers in the result set
	for rows.Next() {
		var followed_id string
		err = rows.Scan(&followed_id)
		if err != nil {
			return nil, fmt.Errorf("error scanning following: %v", err)
		}
		followed = append(followed, follower_id)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during iteration: %v", err)
	}
	// successfully retrieved following
	return followed, nil

}

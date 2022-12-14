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
		return []string{}, fmt.Errorf("error fetching following: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all followers in the result set
	for rows.Next() {
		var followed_id string
		err = rows.Scan(&followed_id)
		if err != nil {
			return []string{}, fmt.Errorf("error scanning following: %w", err)
		}
		followed_name, err := db.GetNameById(followed_id)
		if err != nil {
			return []string{}, err
		}
		followed = append(followed, followed_name)
	}
	if err = rows.Err(); err != nil {
		return []string{}, fmt.Errorf("error encountered during iteration: %w", err)
	}
	// successfully retrieved following
	return followed, nil

}

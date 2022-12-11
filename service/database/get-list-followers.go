package database

import "fmt"

// get the list of the users that follows a user
func (db *appdbimpl) GetFollowers(followed_id string) ([]string, error) {

	const query = `
SELECT follower_id
FROM follow
WHERE followed_id =?`

	var followers []string

	rows, err := db.c.Query(query, followed_id)
	if err != nil {
		return nil, fmt.Errorf("error fetching followers: %v", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all followers in the result set
	for rows.Next() {
		var follower_id string
		err = rows.Scan(&follower_id)
		if err != nil {
			return nil, fmt.Errorf("error scanning followers: %v", err)
		}
		follower_name, err := db.GetNameById(follower_id)
		if err != nil {
			return []string{}, err
		}
		followers = append(followers, follower_name)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during iteration: %v", err)
	}
	// successfully retrieved followers
	return followers, nil

}

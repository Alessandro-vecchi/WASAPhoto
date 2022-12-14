package database

import "fmt"

// get list of the users that are banned from another user
func (db *appdbimpl) GetBannedUsers(banner_id string) ([]string, error) {

	const query = `
		SELECT banned_id
		FROM ban
		WHERE banner_id = ?`

	var banned_users []string

	rows, err := db.c.Query(query, banner_id)
	if err != nil {
		return []string{}, fmt.Errorf("error fetching banned users: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all banned users in the result set
	for rows.Next() {
		var banned_id string
		err = rows.Scan(&banned_id)
		if err != nil {
			return []string{}, fmt.Errorf("error scanning banned users: %w", err)
		}
		banned_name, err := db.GetNameById(banned_id)
		if err != nil {
			return []string{}, err
		}
		banned_users = append(banned_users, banned_name)
	}
	if err = rows.Err(); err != nil {
		return []string{}, fmt.Errorf("error encountered during iteration: %w", err)
	}
	// successfully retrieved banned users
	return banned_users, nil

}

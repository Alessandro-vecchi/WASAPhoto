package database

import "fmt"

// get the list of the users that liked a photo
func (db *appdbimpl) GetLikes(photoId string) ([]string, error) {

	const query = `
	SELECT user_id
	FROM likes
	WHERE photo_id =?`

	var likes []string

	rows, err := db.c.Query(query, photoId)
	if err != nil {
		return nil, fmt.Errorf("error fetching likes: %v", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all followers in the result set
	for rows.Next() {
		var user_like_id string
		err = rows.Scan(&user_like_id)
		if err != nil {
			return nil, fmt.Errorf("error scanning likes: %v", err)
		}
		like_name, err := db.GetNameById(user_like_id)
		if err != nil {
			return []string{}, err
		}
		likes = append(likes, like_name)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during iteration: %v", err)
	}
	// successfully retrieved followers
	return likes, nil

}

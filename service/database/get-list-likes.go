package database

import "fmt"

// get the list of the users that liked a photo
func (db *appdbimpl) GetLikes(photoId string) ([]string, error) {

	const query = `
	SELECT liker_id
	FROM likes
	WHERE photo_id =?`

	var likes []string

	rows, err := db.c.Query(query, photoId)
	if err != nil {
		return []string{}, fmt.Errorf("error fetching likes: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all followers in the result set
	for rows.Next() {
		var user_like_id string
		err = rows.Scan(&user_like_id)
		if err != nil {
			return []string{}, fmt.Errorf("error scanning likes: %w", err)
		}
		like_name, err := db.GetNameById(user_like_id)
		if err != nil {
			return []string{}, err
		}
		likes = append(likes, like_name)
	}
	if err = rows.Err(); err != nil {
		return []string{}, fmt.Errorf("error encountered during iteration: %w", err)
	}
	// successfully retrieved followers
	return likes, nil

}

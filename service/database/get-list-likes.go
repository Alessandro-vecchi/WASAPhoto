package database

import "fmt"

// get the list of the users that liked a photo
func (db *appdbimpl) GetLikes(photoId string) ([]Short_profile_db, error) {

	const query = `
	SELECT profilePictureUrl, user_id 
	FROM profile, ( SELECT liker_id FROM likes WHERE photo_id =? ) as likers 
	WHERE profile.user_id = likers.liker_id;`

	var likes []Short_profile_db

	rows, err := db.c.Query(query, photoId)
	if err != nil {
		return []Short_profile_db{}, fmt.Errorf("error fetching likes: %w", err)
	}

	defer func() { _ = rows.Close() }()

	// Read all the users with the respective profile photo that liked the post
	for rows.Next() {
		var short_profile Short_profile_db
		var user_like_id string
		err = rows.Scan(&short_profile.ProfilePictureUrl, &user_like_id)
		if err != nil {
			return []Short_profile_db{}, fmt.Errorf("error scanning likes: %w", err)
		}
		short_profile.Username, err = db.GetNameById(user_like_id)
		if err != nil {
			return []Short_profile_db{}, err
		}
		likes = append(likes, short_profile)
	}
	if err = rows.Err(); err != nil {
		return []Short_profile_db{}, fmt.Errorf("error encountered during iteration: %w", err)
	}
	// successfully retrieved list of likers and profile pic
	return likes, nil

}

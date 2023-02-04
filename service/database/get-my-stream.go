package database

import "fmt"

// get the list of the users that liked a photo
func (db *appdbimpl) GetMyStream(user_id string, offset string) ([]Photo_db, error) {
	/* Here we want to get the list of photos ordered in reverse
	   chronological order of the users we are following.
	   1. Get the list of the users followed by us;
	   2. Filter all the photos by using this list of users:
	       a. If the photo belongs to one of these users include it in the stream;
		   b. Otherwise, don't consider it.
	   3. Order in reverse chronological order

	*/
	query := ` 
	SELECT photos.photo_id, photos.timestamp, photos.image, photos.caption, photos.user_id
	FROM photos, ( SELECT followed_id FROM follow WHERE follower_id = ? ) as followed_users
    WHERE photos.user_id = followed_users.followed_id
	ORDER BY photos.timestamp DESC
	LIMIT 5
	OFFSET ?;`

	rows, err := db.c.Query(query, user_id, offset)
	if err != nil {
		return []Photo_db{}, fmt.Errorf("error encountered while querying: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var photos []Photo_db
	for rows.Next() {
		var photo Photo_db
		err = rows.Scan(&photo.PhotoId, &photo.Timestamp, &photo.Image, &photo.Caption, &photo.UserId)
		if err != nil {
			return []Photo_db{}, err
		}
		photos = append(photos, photo)
	}
	if err = rows.Err(); err != nil {
		return []Photo_db{}, err
	}
	return photos, nil
}

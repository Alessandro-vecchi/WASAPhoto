package database

func (db *appdbimpl) GetListUserPhotos(user_id string) ([]Photo_db, error) {

	const query = `
		SELECT *
		FROM photos
		WHERE user_id = ?
		ORDER BY timestamp DESC;`

	var photos []Photo_db

	// Issue the query, using the user_id as filter
	rows, err := db.c.Query(query,
		user_id)
	if err != nil {
		return []Photo_db{}, err
	}
	defer func() { _ = rows.Close() }()

	// Read all photos in the result set
	for rows.Next() {
		var p Photo_db
		err = rows.Scan(&p.UserId, &p.PhotoId, &p.Timestamp, &p.Caption, &p.Image)
		if err != nil {
			return []Photo_db{}, err
		}
		photos = append(photos, p)
	}
	if err = rows.Err(); err != nil {
		return []Photo_db{}, err
	}

	return photos, nil

}

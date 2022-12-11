package database

func (db *appdbimpl) GetListUserPhotos(user_id string) ([]Photo_db, error) {

	const query = `
SELECT *
FROM photos
WHERE user_id =?`

	var ret []Photo_db

	// Issue the query, using the user_id as filter
	rows, err := db.c.Query(query,
		user_id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the result set
	for rows.Next() {
		var p Photo_db
		err = rows.Scan(&p.UserId, &p.PhotoId, &p.Timestamp, &p.Caption, &p.Image)
		if err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil

}

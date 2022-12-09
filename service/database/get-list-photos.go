package database

import "fmt"

func (db *appdbimpl) GetListUserPhotos(user_id string) ([]Photo_db, error) {
	// Here we need to get all fountains inside a given range. One simple solution is to rely on GIS/Spatial functions
	// from the DB itself. GIS/Spatial functions are those dedicated to geometry/geography/space computation.
	//
	// However, some databases (like SQLite) do not support these functions. So, we use a naive approach: instead of
	// drawing a circle for a given range, we get slightly more fountains by retrieving a square area, and then we will
	// filter the result later ("cutting the corner").
	//
	// Steps are:
	// 1. We compute a square ("bounding box") that contains the circle. The square will have edges with the same length
	//    of the range of the circle.
	// 2. For each resulting fountain, we will check (using Go and some math) if it's inside the range or not.

	const query = `
SELECT *
FROM photos
WHERE user_id =?`

	var ret []Photo_db

	// Issue the query, using the bounding box as filter
	rows, err := db.c.Query(query,
		user_id)
	fmt.Println("1", rows)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the result set
	fmt.Println("2", rows)
	for rows.Next() {
		var p Photo_db
		err = rows.Scan(&p.UserId, &p.PhotoId, &p.Timestamp, &p.LikesCount, &p.CommentsCount, &p.Caption, &p.Image)
		if err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	fmt.Println("3", ret, err)
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil

}

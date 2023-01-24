package database

func (db *appdbimpl) GetComments(photoId string) ([]Comment_db, error) {

	const query = `
        SELECT user_id, comment_id, created_in, body, modified_in, is_reply_comment
        FROM comments
        WHERE photo_id =?
		ORDER BY created_in DESC;`

	/* const query = `WITH RECURSIVE comments_tree AS (
		-- Get the root comments (comments with no parent)
		SELECT user_id, comment_id, created_in, body, modified_in, is_reply_comment, 0 as depth
		FROM comments
		WHERE photo_id =? AND parent_id = ""
		ORDER BY created_in DESC

		UNION ALL

		-- Get the child comments
		SELECT c.user_id, c.comment_id, c.created_in, c.body, c.modified_in, c.is_reply_comment, ct.depth + 1
		FROM comments c
		JOIN comments_tree ct ON ct.comment_id = c.parent_id
		WHERE ct.depth < 2
		ORDER BY c.created_in DESC
	)

	SELECT * FROM comments_tree;
		` */

	var comments []Comment_db

	// Issue the query, using the id of the photo as filter
	rows, err := db.c.Query(query, photoId)
	if err != nil {
		return []Comment_db{}, err
	}
	defer func() { _ = rows.Close() }()

	// Read all comments in the result set
	for rows.Next() {
		var comment Comment_db
		err = rows.Scan(&comment.UserId, &comment.CommentId, &comment.Created_in, &comment.Body, &comment.Modified_in, &comment.IsReplyComment)
		if err != nil {
			return []Comment_db{}, err
		}
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return []Comment_db{}, err
	}

	return comments, nil

}

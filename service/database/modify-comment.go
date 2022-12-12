package database

func (db *appdbimpl) ModifyComment(c Comment_db) (Comment_db, error) {

	const query = `
		SELECT body, modified_in
		FROM comments
		WHERE comment_id = ?`

	var old Comment_db
	comment := db.c.QueryRow(query, c.CommentId)

	err := comment.Scan(&old.Body, &old.Modified_in)
	if err != nil {
		return Comment_db{}, err
	}
	_, err = db.c.Exec(`UPDATE comments SET body = ? AND modified_in = ? WHERE comment_id = ?`, c.Body, c.Modified_in, c.CommentId)

	if err != nil {
		return old, err
	}

	return c, nil
}

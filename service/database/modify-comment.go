package database

import "time"

func (db *appdbimpl) ModifyComment(c Comment_db) (Comment_db, error) {
	c.Modified_in = time.Now().Format(time.RFC3339)
	_, err := db.c.Exec(`UPDATE comments SET body = ?, modified_in = ? WHERE comment_id = ?`, c.Body, c.Modified_in, c.CommentId)

	if err != nil {
		return Comment_db{}, err
	}

	return c, nil
}

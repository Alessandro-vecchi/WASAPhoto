package database

func (db *appdbimpl) GetSingleComment(commentId string) (Comment_db, error) {

	var comment Comment_db

	const query = `
        SELECT *
        FROM comments
        WHERE comment_id =?`

	// Issue the query, using the id of the photo as filter
	err := db.c.QueryRow(query, commentId).Scan(&comment.UserId, &comment.CommentId, &comment.Created_in, &comment.Body, &comment.PhotoId, &comment.Modified_in, &comment.IsReplyComment, &comment.ParentId)
	if err != nil {
		return Comment_db{}, ErrCommentNotExists
	}

	return comment, nil

}

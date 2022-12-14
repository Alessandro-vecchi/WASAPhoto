package database

import (
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

func (db *appdbimpl) CommentPhoto(photoId string, c Comment_db) (Comment_db, error) {

	// photo does not exist, creating a new username
	rawCommentId, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to get UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID for the comment: %v", rawCommentId)
	commentId := rawCommentId.String()
	c.Created_in = time.Now().Format(time.RFC3339)
	c.Modified_in = c.Created_in
	if c.IsReplyComment {
		_, err = db.c.Exec(`INSERT INTO comments (user_id, comment_id, created_in, body, photo_id, modified_in, is_reply_comment, parent_id) VALUES (?,?,?,?,?,?,?,?)`, c.UserId, commentId, c.Created_in, c.Body, photoId, c.Modified_in, c.IsReplyComment, "")
	} else {
		_, err = db.c.Exec(`INSERT INTO comments (user_id, comment_id, created_in, body, photo_id, modified_in, is_reply_comment, parent_id) VALUES (?,?,?,?,?,?,?,?)`, c.UserId, commentId, c.Created_in, c.Body, photoId, c.Modified_in, c.IsReplyComment, c.ParentId)
	}
	if err != nil {
		return Comment_db{}, fmt.Errorf("comment could not be created: %w", err)
	}
	return c, nil

}

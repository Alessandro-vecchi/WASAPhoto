package models

import (
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
)

var ()

// Attributes of a comment
type Comment struct {
	CommentId string `json:"commentId,omitempty"`
	// Date and time of creation of the comment following RFC3339
	Created_in string `json:"created_in,omitempty"`
	// Content of the comment
	Body string `json:"body,omitempty"`
	// Id of the photo under which the comments are being written
	PhotoId string `json:"photoId,omitempty"`
	// Username of the user that created the comment
	Author string `json:"author,omitempty"`
	// Date and time of when the comment was modified following RFC3339
	Modified_in string `json:"modified_in,omitempty"`
	// States if a comment is a reply to another comment or not
	IsReplyComment bool `json:"isReplyComment,omitempty"`
}

func (c *Comment) FromDatabase(comment database.Comment_db, db database.AppDatabase) {
	c.CommentId = comment.CommentId
	c.PhotoId = comment.PhotoId
	c.Created_in = comment.Created_in
	c.Modified_in = comment.Modified_in
	c.IsReplyComment = comment.IsReplyComment
	c.Body = comment.Body
	c.Author, _ = db.GetNameById(comment.UserId)
}

// ToDatabase returns the profile in a database-compatible representation
func (comment *Comment) ToDatabase(db database.AppDatabase) database.Comment_db {
	id, _ := db.GetIdByName(comment.Author)
	return database.Comment_db{
		UserId:         id,
		CommentId:      comment.CommentId,
		Created_in:     comment.Created_in,
		Body:           comment.Body,
		PhotoId:        comment.PhotoId,
		Modified_in:    comment.Modified_in,
		IsReplyComment: comment.IsReplyComment,
	}
}

func (c *Comment) IsValid() bool {
	return len(c.Body) <= 300 && captionRx.MatchString(c.Body) &&
		usernameRx.MatchString(c.Author)

}

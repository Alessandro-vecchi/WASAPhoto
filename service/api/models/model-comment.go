package models

import (
	"time"

	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
)

// Attributes of a comment
type Comment struct {
	// Id of the comment
	CommentId string `json:"commentId,omitempty"`
	// Date and time of creation of the comment following RFC3339
	Created_in string `json:"created_in,omitempty"`
	// Content of the comment
	Body string `json:"body,omitempty"`
	// Username of the user that created the comment
	Author string `json:"author,omitempty"`
	// Date and time of when the comment was modified following RFC3339
	Modified_in string `json:"modified_in,omitempty"`
	// States if a comment is a reply to another comment or not
	IsReplyComment bool `json:"isReplyComment,omitempty"`
	// Id of the parent comment, "" if top-level comment
	ParentId string `json:"parentId,omitempty"`
	// Profile picture of the author
	Profile_pic string `json:"profile_pic,omitempty"`
}

func (c *Comment) FromDatabase(comment database.Comment_db, db database.AppDatabase) {
	c.CommentId = comment.CommentId
	c.Created_in = comment.Created_in
	c.Modified_in = comment.Modified_in
	c.IsReplyComment = comment.IsReplyComment
	c.Body = comment.Body
	c.Author, _ = db.GetNameById(comment.UserId)
	c.ParentId = comment.ParentId
	c.Profile_pic, _ = db.GetProfilePic(comment.UserId)
}

// ToDatabase returns the profile in a database-compatible representation
func (comment *Comment) ToDatabase(db database.AppDatabase, photo_id string) database.Comment_db {
	id, _ := db.GetIdByName(comment.Author)
	return database.Comment_db{
		UserId:         id,
		CommentId:      comment.CommentId,
		Created_in:     comment.Created_in,
		Body:           comment.Body,
		PhotoId:        photo_id,
		Modified_in:    time.Now().Format(time.RFC3339),
		IsReplyComment: comment.IsReplyComment,
		ParentId:       comment.ParentId,
	}
}

func (c *Comment) IsValid() bool {
	return len(c.Body) <= 300 && captionRx.MatchString(c.Body) &&
		usernameRx.MatchString(c.Author)

}

package models

//"github.com/Alessandro-vecchi/WASAPhoto/service/database"

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

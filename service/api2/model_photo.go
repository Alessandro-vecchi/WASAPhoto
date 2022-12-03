package api

import (
	"time"
)

// Attributes of a photo
type Photo struct {
	PhotoId int64 `json:"photoId,omitempty"`
	// Date and time of creation following RFC3339
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Number of likes
	LikesCount int32 `json:"likes_count,omitempty"`
	// Number of comments
	CommentsCount int32 `json:"comments_count,omitempty"`
	// URL of the image just uploaded. | Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	Image string `json:"image,omitempty"`
	// A written description or explanation about a photo to provide more context
	Caption string `json:"caption,omitempty"`

	Username string `json:"username,omitempty"`
}

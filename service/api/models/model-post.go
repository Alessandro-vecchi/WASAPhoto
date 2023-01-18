package models

import (
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
)

// Attributes of a photo
type Post struct {
	PhotoId string `json:"photoId,omitempty"`
	// Date and time of creation following RFC3339
	Timestamp string `json:"timestamp,omitempty"`
	// Number of likes
	LikesCount uint32 `json:"likes_count"`
	// Number of comments
	CommentsCount uint32 `json:"comments_count"`
	// URL of the image just uploaded. | Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	Image string `json:"image,omitempty"`
	// A written description or explanation about a photo to provide more context
	Caption string `json:"caption,omitempty"`
	// Username of the user
	Username string `json:"username,omitempty"`
	// Profile picture of the author
	Profile_pic string `json:"profile_pic,omitempty"`
}

func (p *Post) FromDatabase(photo database.Photo_db, db database.AppDatabase) {
	p.PhotoId = photo.PhotoId
	p.Timestamp = photo.Timestamp
	p.LikesCount = db.CountStuffs("photo_id", "likes", p.PhotoId)
	p.CommentsCount = db.CountStuffs("photo_id", "comments", p.PhotoId)
	p.Image = photo.Image
	p.Caption = photo.Caption
	p.Username, _ = db.GetNameById(photo.UserId)
	p.Profile_pic, _ = db.GetProfilePic(photo.UserId)
}

// ToDatabase returns the profile in a database-compatible representation
func (photo *Post) ToDatabase(db database.AppDatabase) database.Photo_db {
	id, _ := db.GetIdByName(photo.Username)
	return database.Photo_db{
		PhotoId:   photo.PhotoId,
		Timestamp: photo.Timestamp,
		Image:     photo.Image,
		Caption:   photo.Caption,
		UserId:    id,
	}
}

// IsValid checks the validity of the content. In particular, the caption should have a max length of 150 and should match its regex. Same for the image, and the
// username should be smaller than 16 characters and greater or equal than 3.
func (p *Post) IsValid() bool {
	return len(p.Caption) <= 150 && captionRx.MatchString(p.Caption) &&
		len(p.Image) <= 150 && imageUrlRx.MatchString(p.Image) &&
		3 <= len(p.Username) && len(p.Username) <= 16 && usernameRx.MatchString(p.Username)

}

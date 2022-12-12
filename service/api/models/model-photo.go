package models

import (
	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
)

var (
	imageUrlRx = ppURLRx
	captionRx  = bioRx
)

// Attributes of a photo
type Photo struct {
	PhotoId string `json:"photoId,omitempty"`
	// Date and time of creation following RFC3339
	Timestamp string `json:"timestamp,omitempty"`
	// Number of likes
	LikesCount uint32 `json:"likes_count,omitempty"`
	// Number of comments
	CommentsCount uint32 `json:"comments_count,omitempty"`
	// URL of the image just uploaded. | Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	Image string `json:"image,omitempty"`
	// A written description or explanation about a photo to provide more context
	Caption string `json:"caption,omitempty"`

	Username string `json:"username,omitempty"`
}

func (p *Photo) FromDatabase(photo database.Photo_db, db database.AppDatabase) {
	p.PhotoId = photo.PhotoId
	p.Timestamp = photo.Timestamp
	p.LikesCount = db.CountStuffs("photo_id", "likes", p.PhotoId)
	p.CommentsCount = db.CountStuffs("photo_id", "comments", p.PhotoId)
	p.Image = photo.Image
	p.Caption = photo.Caption
	p.Username, _ = db.GetNameById(photo.UserId)
}

// ToDatabase returns the profile in a database-compatible representation
func (photo *Photo) ToDatabase(db database.AppDatabase) database.Photo_db {
	id, _ := db.GetIdByName(photo.Username)
	return database.Photo_db{
		PhotoId:   photo.PhotoId,
		Timestamp: photo.Timestamp,
		Image:     photo.Image,
		Caption:   photo.Caption,
		UserId:    id,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either FountainStatusGood or FountainStatusFaulty. Note that the ID is not checked, as fountains
// read from requests have zero IDs as the user won't send us the ID in that way.
func (p *Photo) IsValid() bool {
	return len(p.Caption) <= 100 && captionRx.MatchString(p.Caption) &&
		len(p.Image) <= 150 && imageUrlRx.MatchString(p.Image) &&
		usernameRx.MatchString(p.Username)

}

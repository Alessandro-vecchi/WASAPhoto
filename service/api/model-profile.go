package api

import (
	"regexp"

	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
)

var (
	usernameRx = regexp.MustCompile(`^[a-zA-Z0-9-_]*$`)
	ppURLRx    = regexp.MustCompile(`^(https?:\/\/.*\.(?:png|jpg|jpeg))$`)
	bioRx      = regexp.MustCompile(`^[a-zA-Z0-9,._:;?!\x27\- ]*$`)
)

// Represents the information seen in the Profile Page of a user
type Profile struct {

	// ID of the user
	ID string `json:"userID,omitempty"`
	// Name of the user
	Username string `json:"username,omitempty"`
	// Number of photos in the profile of the user
	PicturesCount int32 `json:"pictures_count,omitempty"`
	// Number of users that follow the profile
	FollowersCount int32 `json:"followers_count,omitempty"`
	// number of users that the user follows
	FollowsCount int32 `json:"follows_count,omitempty"`
	// URL of the profile picture. Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	ProfilePictureUrl string `json:"profile_picture_url,omitempty"`
	// Biography of the profile. Just allowing alphanumeric characters and basic punctuation.
	Bio string `json:"bio,omitempty"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (p *Profile) FromDatabase(profile database.DProfile) {
	p.ID = profile.ID
	p.Username = profile.Username
	p.PicturesCount = profile.PicturesCount
	p.FollowersCount = profile.FollowersCount
	p.FollowsCount = profile.FollowsCount
	p.ProfilePictureUrl = profile.ProfilePictureUrl
	p.Bio = profile.Bio
}

// ToDatabase returns the profile in a database-compatible representation
func (profile *Profile) ToDatabase() database.DProfile {
	return database.DProfile{
		ID:                profile.ID,
		Username:          profile.Username,
		PicturesCount:     profile.PicturesCount,
		FollowersCount:    profile.FollowersCount,
		FollowsCount:      profile.FollowsCount,
		ProfilePictureUrl: profile.ProfilePictureUrl,
		Bio:               profile.Bio,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either FountainStatusGood or FountainStatusFaulty. Note that the ID is not checked, as fountains
// read from requests have zero IDs as the user won't send us the ID in that way.
func (p *Profile) isValid() bool {
	return len(p.ID) >= 1 && len(p.ID) <= 20 && usernameRx.MatchString(p.ID) &&
		len(p.Username) >= 3 && len(p.Username) <= 16 && usernameRx.MatchString(p.Username) &&
		p.PicturesCount >= 0 &&
		p.FollowersCount >= 0 &&
		p.FollowsCount >= 0 &&
		len(p.ProfilePictureUrl) >= 0 && len(p.ProfilePictureUrl) <= 150 && ppURLRx.MatchString(p.ProfilePictureUrl) &&
		len(p.Bio) >= 0 && len(p.Bio) <= 150 && bioRx.MatchString(p.Bio)

}

package models

import "github.com/Alessandro-vecchi/WASAPhoto/service/database"

// It's shown when viewing list of followers, likers...
type Short_profile struct {
	// list of username + profile picture
	S_p []Short_profile_api `json:"short_profile"`
	// condition that assess whether the logged user follow the profile
	Cond bool `json:"cond"`
}

// It's shown when viewing list of followers, likers...
type Short_profile_api struct {

	// Name of the user
	Username string `json:"username"`
	// URL of the profile picture. Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	ProfilePictureUrl string `json:"profilePictureUrl"`
}

func (p *Short_profile) FromDatabase(profile []database.Short_profile_db, my_name string) {
	var sp []Short_profile_api
	p.Cond = false
	for _, v := range profile {
		if v.Username == my_name {
			p.Cond = true
		}
		var s Short_profile_api
		s.Username = v.Username
		s.ProfilePictureUrl = v.ProfilePictureUrl
		sp = append(sp, s)
	}
	p.S_p = sp
}

/*
// ToDatabase returns the profile in a database-compatible representation
func (profile *Short_profile) ToDatabase() database.Short_profile_db {
	for _, v := range profile.s_p {
		return v
	}
	return database.Short_profile_db{}

} */

func Conversion(id_A string, id_B string, db database.AppDatabase) (string, string, error) {
	name_A, err := db.GetNameById(id_A)
	if err != nil {
		return "", "", err
	}
	name_B, err := db.GetNameById(id_B)
	if err != nil {
		return "", "", err
	}
	return name_A, name_B, nil
}

func AreTheSame(follower_id string, followed_id string) bool {
	return follower_id == followed_id
}

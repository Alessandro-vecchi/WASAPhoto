package models

import "github.com/Alessandro-vecchi/WASAPhoto/service/database"

// It's shown when viewing list of followers, likers...
type Short_profile struct {
	// list of username + profile picture
	S_p []database.Short_profile_db `json:"short_profile"`
	// condition that assess whether the logged user follow the profile
	Cond bool `json:"cond"`
}

func (p *Short_profile) FromDatabase(profile []database.Short_profile_db, my_name string) {
	p.S_p = profile
	p.Cond = false
	for _, v := range profile {
		if v.Username == my_name {
			p.Cond = true
		}
	}
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

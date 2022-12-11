package models

import "github.com/Alessandro-vecchi/WASAPhoto/service/database"

func IsLikingHimself(photo_id string, user_id string, db database.AppDatabase) bool {
	profile_owner, _ := db.GetProfileOwner(photo_id)
	return profile_owner == user_id
}

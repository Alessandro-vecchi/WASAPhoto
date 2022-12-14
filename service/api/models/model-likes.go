package models

import (
	"errors"

	"github.com/Alessandro-vecchi/WASAPhoto/service/database"
)

func IsLikingHimself(photo_id string, user_id string, db database.AppDatabase) (bool, error) {
	profile_owner, err := db.GetProfileOwner(photo_id)
	if errors.Is(err, database.ErrPhotoNotExists) {
		return false, database.ErrPhotoNotExists
	} else if err != nil {
		return false, err
	}
	return profile_owner == user_id, nil
}

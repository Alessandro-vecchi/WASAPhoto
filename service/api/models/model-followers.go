package models

import "github.com/Alessandro-vecchi/WASAPhoto/service/database"

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

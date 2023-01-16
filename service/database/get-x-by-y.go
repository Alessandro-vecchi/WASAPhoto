package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetNameById(user_id string) (string, error) {

	var name string
	const query = `
		SELECT username
		FROM profile
		WHERE user_id = ?`

	err := db.c.QueryRow(query, user_id).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {

		return "", ErrUserNotExists
	}
	return name, nil
}

func (db *appdbimpl) GetIdByName(username string) (string, error) {

	var id string
	const query = `
		SELECT user_id
		FROM profile
		WHERE username = ?`

	err := db.c.QueryRow(query, username).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {

		return "", ErrUserNotExists
	}
	return id, nil
}

func (db *appdbimpl) GetPhotoIdFromCommentId(comment_id string) (string, error) {

	var photo_id string
	const query = `
		SELECT photo_id
		FROM comments
		WHERE comment_id = ?`

	err := db.c.QueryRow(query, comment_id).Scan(&photo_id)
	if errors.Is(err, sql.ErrNoRows) {

		return "", ErrCommentNotExists
	}
	return photo_id, nil
}
func (db *appdbimpl) GetProfilePic(user_id string) (string, error) {

	var profile_pic string
	const query = `
		SELECT profilePictureUrl
		FROM profile
		WHERE user_id = ?`

	err := db.c.QueryRow(query, user_id).Scan(&profile_pic)
	if errors.Is(err, sql.ErrNoRows) {

		return "", ErrUserNotExists
	}
	return profile_pic, nil
}

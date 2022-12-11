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

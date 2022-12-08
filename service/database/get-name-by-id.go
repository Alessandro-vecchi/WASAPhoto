package database

func (db *appdbimpl) GetNameById(user_id string) (string, error) {

	var name string
	const query = `
SELECT username
FROM profile
WHERE user_id = ?`

	err := db.c.QueryRow(query, user_id).Scan(&name)
	if err != nil {

		return "", ErrUserNotExists
	}
	return name, nil
}

package database

import "fmt"

func (db *appdbimpl) UpdateUsername(p DProfile) (DProfile, error) {
	var p_empty DProfile
	res, err := db.c.Exec(`UPDATE Profile SET username =? WHERE user_id=?`, p.Username, p.ID)
	fmt.Println(res)
	if err != nil {
		return p_empty, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return p_empty, err
	} else if rowsAffected == 0 {
		return p_empty, ErrUserNotExists
	}
	return p, nil
}

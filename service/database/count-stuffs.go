package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CountStuffs(filter string, table_name string, filterVal string) uint32 {

	query := `
SELECT COUNT(*) FROM ` + table_name + ` WHERE ` + filter + ` = ?`

	var count uint32
	err := db.c.QueryRow(query, filterVal).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		// id not present
		return 0
	}
	return count
}

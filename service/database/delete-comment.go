package database

import "fmt"

func (db *appdbimpl) UncommentPhoto(commentId string) error {
	res, err := db.c.Exec("DELETE FROM comments WHERE comment_id =?", commentId)
	if err != nil {
		// error deleting the photo
		return fmt.Errorf("error deleting the comment: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the photo didn't exist
		return ErrCommentNotExists
	}
	return nil
}

package models

import (
	"github.com/gofrs/uuid"
)

type Username struct {

	// Name of the user
	Username string `json:"username"`
}

func (u Username) IsValid() bool {
	return len(u.Username) >= 3 && len(u.Username) <= 16 && usernameRx.MatchString(u.Username)
}

func IsValidUUID(id string) bool {
	_, err := uuid.FromString(id)
	return err == nil
}

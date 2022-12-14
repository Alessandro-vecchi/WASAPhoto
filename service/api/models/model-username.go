package models

import (
	"github.com/gofrs/uuid"
)

type Username struct {

	// Name of the user
	Name string `json:"username"`
}

func (u Username) IsValid() bool {
	return len(u.Name) >= 3 && len(u.Name) <= 16 && usernameRx.MatchString(u.Name)
}

func IsValidUUID(id string) bool {
	_, err := uuid.FromString(id)
	return err == nil
}

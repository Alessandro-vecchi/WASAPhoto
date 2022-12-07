package models

type Username struct {

	// Name of the user
	Name string `json:"username"`
}

func (u Username) IsValid() bool {
	return len(u.Name) >= 3 && len(u.Name) <= 16 && usernameRx.MatchString(u.Name)
}

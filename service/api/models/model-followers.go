package models

import "github.com/Alessandro-vecchi/WASAPhoto/service/database"

// Followers relationship
type Follow struct {
	// User_id of the followers
	Follower_name string `json:"followerId,omitempty"`
	// User_id of the user that is being followed
	Followed_name string `json:"followedId,omitempty"`
}

func (f *Follow) FromDatabase(follow database.Follow_db, db database.AppDatabase) error {
	var err error
	f.Follower_name, err = db.GetNameById(follow.Follower_id)
	if err != nil {
		return err
	}
	f.Followed_name, err = db.GetNameById(follow.Followed_id)
	if err != nil {
		return err
	}
	return nil
}

// ToDatabase returns the profile in a database-compatible representation
func (f *Follow) ToDatabase(db database.AppDatabase) (database.Follow_db, error) {
	fed_id, err := db.GetIdByName(f.Follower_name)
	if err != nil {
		return database.Follow_db{}, err
	}
	fer_id, err := db.GetIdByName(f.Followed_name)
	if err != nil {
		return database.Follow_db{}, err
	}
	return database.Follow_db{
		Follower_id: fer_id,
		Followed_id: fed_id,
	}, nil
}

func (f *Follow) AreTheSame() bool {
	return f.Follower_name == f.Followed_name
}

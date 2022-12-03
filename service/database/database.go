/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrUserNotExists = errors.New("user does not exists")
)

// Represents the information seen in the Profile Page of a user
type Profile struct {

	// ID of the user
	ID string `json:"userID,omitempty"`
	// Name of the user
	Username string `json:"username,omitempty"`
	// Number of photos in the profile of the user
	PicturesCount int32 `json:"pictures_count,omitempty"`
	// Number of users that follow the profile
	FollowersCount int32 `json:"followers_count,omitempty"`
	// number of users that the user follows
	FollowsCount int32 `json:"follows_count,omitempty"`
	// URL of the profile picture. Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	ProfilePictureUrl string `json:"profile_picture_url,omitempty"`
	// Biography of the profile. Just allowing alphanumeric characters and basic punctuation.
	Bio string `json:"bio,omitempty"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// CreateUserProfile creates a new user if he/she doesn't exist

	// Get the user's profile by username
	GetUserProfileByUsername(username string) (Profile, error)

	// Get the user's profile by ID
	//GetUserProfileByID(ID string) (Profile, error)

	// Update username of the user
	//UpdateUsername(username string) (Profile, error)

	//
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='profile';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE profile (
    id TEXT NOT NULL,
    username TEXT NOT NULL,
    picturesCount INTEGER NOT NULL,
    followersCount INTEGER NOT NULL,
	followsCount INTEGER NOT NULL,
    profilePictureUrl TEXT  NULL,
    bio TEXT NULL, PRIMARY KEY(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

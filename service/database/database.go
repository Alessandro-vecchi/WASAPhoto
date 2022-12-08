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
	// ErrDatabaseNotInitialized is returned when the database is not initialized.
	ErrUserNotExists = errors.New("user does not exists")
	ErrUserExists    = errors.New("user already exists")
)

// Represents the information seen in the Profile Page of a user
type Profile_db struct {

	// ID of the user
	ID string
	// Name of the user
	Username string
	// Number of photos in the profile of the user
	PicturesCount uint32
	// Number of users that follow the profile
	FollowersCount uint32
	// number of users that the user follows
	FollowsCount uint32
	// URL of the profile picture. Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	ProfilePictureUrl string
	// Biography of the profile. Just allowing alphanumeric characters and basic punctuation.
	Bio string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// CreateUserProfile creates a new user if he/she doesn't exist

	// Get the user's profile by username
	GetUserProfileByUsername(username string) (Profile_db, error)

	// Get the user's profile by ID
	//GetUserProfileByID(ID string) (Profile, error)

	// if the user does not exist, it will be created and an identifier will be returned.
	// If it does exist, the user identifier will be returned.
	DoLogin(username string) (string, error)

	// Update profile of the user
	UpdateUserProfile(p Profile_db) (Profile_db, error)

	// Delete user profile
	DeleteUserProfile(userID string) error

	// Convert id and name
	GetNameById(userId string) (string, error)
	// check availability
	Ping() error
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
	//f, _ := db.Exec(`DROP TABLE IF EXISTS profile;`)
	//fmt.Println(f)
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='profile';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE profile (
    user_id TEXT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    picturesCount INTEGER DEFAULT 0 NOT NULL,
    followersCount INTEGER DEFAULT 0 NOT NULL,
	followsCount INTEGER DEFAULT 0 NOT NULL,
    profilePictureUrl TEXT DEFAULT "" NOT NULL,
    bio TEXT DEFAULT "" NOT NULL);`
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

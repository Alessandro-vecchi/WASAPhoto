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
	ErrUserNotExists  = errors.New("user does not exist")
	ErrUserExists     = errors.New("user already exists")
	ErrPhotoNotExists = errors.New("photo does not exist")
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

// Attributes of a photo
type Photo_db struct {
	PhotoId string
	// Date and time of creation following RFC3339
	Timestamp string
	// Number of likes
	LikesCount uint32
	// Number of comments
	CommentsCount uint32
	// URL of the image just uploaded. | Accepting only http/https URLs and .png/.jpg/.jpeg extensions.
	Image string
	// A written description or explanation about a photo to provide more context
	Caption string

	UserId string
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

	// Retrieve collection of photos resources of a certain user
	GetListUserPhotos(userId string) ([]Photo_db, error)

	// Upload a photo on the profile of a specific user
	UploadPhoto(userId string, p Photo_db) (Photo_db, error)

	// Get a single photo from the profile of a user
	GetUserPhoto(userId string, photoId string) (Photo_db, error)

	// Delete photo from the profile of a specific user. It also removes likes and comments
	DeletePhoto(photoIs string) error

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
	var tableName string = "profile"
	sqlStmt := `CREATE TABLE profile (
    user_id TEXT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    picturesCount INTEGER DEFAULT 0 NOT NULL,
    followersCount INTEGER DEFAULT 0 NOT NULL,
	followsCount INTEGER DEFAULT 0 NOT NULL,
    profilePictureUrl TEXT DEFAULT "" NOT NULL,
    bio TEXT DEFAULT "" NOT NULL);`

	err := createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	tableName = "photos"
	sqlStmt = `CREATE TABLE photos (
    user_id TEXT NOT NULL,
    photoId TEXT NOT NULL PRIMARY KEY,
	timestamp TEXT DEFAULT "" NOT NULL,
    likesCount INTEGER DEFAULT 0 NOT NULL,
    commentsCount INTEGER DEFAULT 0 NOT NULL,
    caption TEXT DEFAULT "" NOT NULL,
    image TEXT DEFAULT "" NOT NULL);`

	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	return &appdbimpl{
		c: db,
	}, nil
}

func createTables(tableName string, sqlStmt string, db *sql.DB) error {
	var table string
	//f, _ := db.Exec(`DROP TABLE IF EXISTS profile;`)
	//fmt.Println(f)
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='` + tableName + `';`).Scan(&table)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return fmt.Errorf("error creating database structure: %w", err)
		}
	}
	return err
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

package user

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// User holds methods for interacting with users
type User struct{}

// GetUser h
func (u User) GetUser(username string) (*db.User, error) {
	dbUser := new(db.User)

	err := db.Db.Model(dbUser).Where("last_fm_username = ?", username).Limit(1).Select()

	if err != nil {
		lastFM := lastfm.CreateService()

		if !lastFM.ValidateUser(username) {
			return nil, gqlerror.Errorf("The user %s doesn't exist in Last.fm!", username)
		}

		dbUser = &db.User{
			LastFMUsername: username,
		}

		db.Db.Model(dbUser).Insert()
	}

	return dbUser, nil
}

// CreateService creates an instance of the webhook service object
func CreateService() *User {
	service := &User{}

	return service
}

package users

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
)

// Users holds methods for interacting with users
type Users struct{}

// FindUserByDiscordID finds a user by their discord id
func (u Users) FindUserByDiscordID(discordID string) *db.User {
	user := new(db.User)

	err := db.Db.Model(user).Where("discord_id = ?", discordID).Select()

	if err != nil {
		return nil
	}

	return user
}

// CreateUser creates a user if one doesn't already exist
func (u Users) CreateUser(username, discordID, userType string) (*db.User, error) {
	existingUser := u.FindUserByDiscordID(discordID)

	if existingUser != nil {
		return nil, customerrors.EntityAlreadyExists("user")
	}

	user := &db.User{
		Username:  username,
		DiscordID: discordID,
		UserType:  userType,
	}

	_, err := db.Db.Model(user).Insert()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return user, nil
}

// DeleteUser deletes a user if one exists
func (u Users) DeleteUser(discordID string) error {
	existingUser := u.FindUserByDiscordID(discordID)

	if existingUser == nil {
		return customerrors.EntityDoesntExist("user")
	}

	_, err := db.Db.Model(existingUser).WherePK().Delete()

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}

// CreateService creates an instance of the users service object
func CreateService() *Users {
	service := &Users{}

	return service
}

package controllers

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

// Login logs a user in
func Login(username, discordID, userType string) (*model.User, error) {
	usersService := users.CreateService()

	user, err := usersService.CreateUser(username, discordID, userType)

	if err != nil {
		if user != nil {
			if user.Username == username && user.UserType == userType {
				return presenters.PresentUser(user), nil
			}

			user, _ = usersService.ChangeUsername(user, username, userType)

			return presenters.PresentUser(user), nil
		}

		return nil, err
	}

	return presenters.PresentUser(user), nil
}

// Logout logs a user out and deletes all their data
func Logout(discordID string) (*string, error) {
	usersService := users.CreateService()

	err := usersService.DeleteUser(discordID)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

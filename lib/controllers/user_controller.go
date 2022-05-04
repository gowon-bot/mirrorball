package controllers

import (
	"strings"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

// Login logs a user in
func Login(username string, session *string, discordID string) (*model.User, error) {
	usersService := users.CreateService()

	user, err := usersService.CreateUser(username, discordID, session)

	if err != nil {
		if user != nil {
			if user.Username == username {
				return presenters.PresentUser(user), nil
			}

			user, _ = usersService.ChangeUsername(user, username, session)

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

func Users(inputs []*model.UserInput) ([]*model.User, error) {
	usersService := users.CreateService()

	users, err := usersService.GetUsers(inputs)

	if err != nil {
		return nil, err
	}

	return presenters.PresentUsers(users), nil
}

func UpdatePrivacy(userInput model.UserInput, privacy *model.Privacy) (*string, error) {
	if strings.ToUpper(privacy.String()) == db.ConvertPrivacyToString(db.DefaultPrivacy) {
		return nil, customerrors.CannotSetToUnset()
	}

	usersService := users.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	_, err := usersService.UpdatePrivacy(user, db.ConvertPrivacyFromString(privacy.String()))

	return nil, err
}

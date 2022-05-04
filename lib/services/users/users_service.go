package users

import (
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
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

// FindUserByInput finds a user from a generic graphql input
func (u Users) FindUserByInput(userInput model.UserInput) *db.User {
	user := new(db.User)

	query := db.Db.Model(user)

	if userInput.DiscordID != nil {
		query = query.Where("discord_id = ?", userInput.DiscordID)
	}
	if userInput.LastFMUsername != nil {
		query = query.Where("username = ?", userInput.LastFMUsername)
	}

	err := query.Select()

	if err != nil {
		return nil
	}

	return user
}

// CreateUser creates a user if one doesn't already exist
func (u Users) CreateUser(username, discordID string, session *string) (*db.User, error) {
	existingUser := u.FindUserByDiscordID(discordID)

	if existingUser != nil {
		return existingUser, customerrors.EntityAlreadyExistsError("user")
	}

	user := &db.User{
		Username:      username,
		DiscordID:     discordID,
		LastFMSession: session,
		Privacy:       db.DefaultPrivacy,
	}

	_, err := db.Db.Model(user).Insert()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return user, nil
}

// CreateUser creates a user from input if one doesn't already exist
func (u Users) CreateUserFromInput(input model.UserInput) (*db.User, error) {

	username := input.LastFMUsername

	if username == nil || input.DiscordID == nil {
		missingArguments := []string{}

		if username == nil {
			missingArguments = append(missingArguments, "username")
		}
		if input.DiscordID == nil {
			missingArguments = append(missingArguments, "discord ID")
		}

		return nil, customerrors.InsufficientArgumentsSupplied(strings.Join(missingArguments, ", "))
	}

	return u.CreateUser(*username, *input.DiscordID, nil)
}

// DeleteUser deletes a user if one exists
func (u Users) DeleteUser(discordID string) error {
	existingUser := u.FindUserByDiscordID(discordID)

	if existingUser == nil {
		return customerrors.EntityDoesntExistError("user")
	}

	_, err := db.Db.Model(existingUser).WherePK().Delete()

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}

// GetUsersByDiscordIDs returns a list of users that have one of the given discord ids
func (u Users) GetUsersByDiscordIDs(discordIDs []string) ([]*db.User, error) {
	var users []*db.User

	err := db.Db.Model(&users).Where("discord_id IN (?)", pg.In(discordIDs)).Select()

	if err != nil {
		return users, nil
	}

	return users, nil
}

// ChangeUsername changes a user's username
func (u Users) ChangeUsername(user *db.User, username string, session *string) (*db.User, error) {
	// TODO: delete data

	user.Username = username
	user.LastFMSession = session

	_, err := db.Db.Model(user).WherePK().Update()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return user, nil
}

func (u Users) UpdatePrivacy(user *db.User, privacy int64) (*db.User, error) {
	user.Privacy = privacy

	_, err := db.Db.Model(user).WherePK().Update()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return user, nil
}

func (u Users) GetUsers(inputs []*model.UserInput) ([]*db.User, error) {
	var users []*db.User

	parser := inputparser.CreateParser(
		db.Db.Model(&users)).ParseUserInputs(inputs, inputparser.InputSettings{UserRelation: "-"})

	err := parser.GetQuery().Select()

	if err != nil {
		return users, customerrors.DatabaseUnknownError()
	}

	return users, nil
}

// CreateService creates an instance of the users service object
func CreateService() *Users {
	service := &Users{}

	return service
}

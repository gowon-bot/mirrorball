package controllers

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/response"
	"github.com/jivison/gowon-indexer/lib/services/users"
	"github.com/jivison/gowon-indexer/lib/tasks"
)

// FullIndex downloads a users full data and caches it
func FullIndex(userInput model.UserInput, forceUserCreate *bool) (*model.TaskStartResponse, error) {
	usersService := users.CreateService()
	responseService := response.CreateService()

	user := usersService.FindUserByInput(userInput)
	token := responseService.GenerateToken()

	if user == nil && forceUserCreate != nil && *forceUserCreate {
		newUser, err := usersService.CreateUserFromInput(userInput)

		if err != nil {
			return nil, err
		}

		user = newUser
	}

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	tasks.TaskServer.SendIndexUserTask(user, token)

	return responseService.BuildTaskStartResponse("index_user", token, true), nil
}

// Update downloads the latest data and updates the cache
func Update(userInput model.UserInput, forceUserCreate *bool) (*model.TaskStartResponse, error) {
	usersService := users.CreateService()
	responseService := response.CreateService()

	user := usersService.FindUserByInput(userInput)
	token := responseService.GenerateToken()

	if user == nil && forceUserCreate != nil && *forceUserCreate {
		newUser, err := usersService.CreateUserFromInput(userInput)

		if err != nil {
			return nil, err
		}

		user = newUser
	}

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	if user.LastIndexed.IsZero() {
		indexForceUserCreate := false

		return FullIndex(userInput, &indexForceUserCreate)
	}

	tasks.TaskServer.SendUpdateUserTask(user, token)

	return responseService.BuildTaskStartResponse("update_user", token, true), nil
}

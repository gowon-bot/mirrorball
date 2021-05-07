package controllers

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

func Plays(userInput model.UserInput, pageInput *model.PageInput) ([]*model.Play, error) {
	usersService := users.CreateService()
	indexingService := indexing.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	plays, err := indexingService.GetPlays(*user, pageInput)

	if err != nil {
		return nil, err
	}

	var builtPlays []*model.Play

	for _, play := range plays {
		builtPlays = append(builtPlays, presenters.PresentPlay(&play))
	}

	return builtPlays, nil
}

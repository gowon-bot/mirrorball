package controllers

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

func Plays(playsInput model.PlaysInput, pageInput *model.PageInput) (*model.PlaysResponse, error) {
	indexingService := indexing.CreateService()

	plays, err := indexingService.GetPlays(playsInput, pageInput)

	if err != nil {
		return nil, err
	}

	count, err := indexingService.CountPlays(playsInput)

	if err != nil {
		return nil, err
	}

	return &model.PlaysResponse{Plays: presenters.PresentPlays(plays), PageInfo: presenters.PresentPageInfo(count)}, nil
}

func ArtistPlays(userInput model.UserInput, settings *model.ArtistPlaysSettings) ([]*model.ArtistCount, error) {
	usersService := users.CreateService()
	indexingService := indexing.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	plays, err := indexingService.GetArtistPlays(*user, settings)

	if err != nil {
		return nil, err
	}

	return presenters.PresentArtistCounts(plays), nil
}

func AlbumPlays(userInput model.UserInput, settings *model.AlbumPlaysSettings) ([]*model.AlbumCount, error) {
	usersService := users.CreateService()
	indexingService := indexing.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	plays, err := indexingService.GetAlbumPlays(*user, settings)

	if err != nil {
		return nil, err
	}

	return presenters.PresentAlbumCounts(plays), nil
}

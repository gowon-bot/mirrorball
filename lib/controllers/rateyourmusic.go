package controllers

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/rateyourmusic"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

func ImportRatings(csvString string, userInput model.UserInput) (*string, error) {
	usersService := users.CreateService()
	rymsService := rateyourmusic.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	rawRatings, err := rymsService.ParseRYMSExport(csvString)

	if err != nil {
		return nil, err
	}

	rymsService.ResetRatings(*user)

	rymsAlbumsMap, err := rymsService.ConvertRateYourMusicAlbums(rawRatings)

	if err != nil {
		return nil, err
	}

	_, err = rymsService.SaveRatings(rawRatings, rymsAlbumsMap, *user)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func Ratings(settings *model.RatingsSettings) ([]*model.Rating, error) {
	rymsService := rateyourmusic.CreateService()

	ratings, err := rymsService.GetRatings(settings)

	if err != nil {
		return nil, err
	}

	return presenters.PresentRatings(ratings), nil
}

func RateYourMusicArtist(keywords string) (*model.RateYourMusicArtist, error) {
	rymsService := rateyourmusic.CreateService()

	album, err := rymsService.GetArtist(keywords)

	if err != nil {
		return nil, nil
	}

	return presenters.PresentRateYourMusicArtist(*album), nil
}

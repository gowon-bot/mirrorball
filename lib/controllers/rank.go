package controllers

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

func ArtistRank(artistInput model.ArtistInput, userInput model.UserInput, serverID *string) (*model.ArtistRankResponse, error) {
	usersService := users.CreateService()
	indexingService := indexing.CreateService()
	analysisService := analysis.CreateService()

	artist, err := indexingService.GetArtist(artistInput, false)

	if err != nil {
		return nil, err
	}

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	settings := &model.WhoKnowsSettings{GuildID: serverID}

	whoKnows, err := analysisService.WhoKnowsArtist(artist, settings)

	if err != nil {
		return nil, err
	}

	var plays *db.ArtistCount
	rank := -1

	for idx, row := range whoKnows {
		if row.UserID == user.ID {
			plays = &row
			rank = idx
			break
		}
	}

	return presenters.PresentArtistRank(artist, whoKnows, plays, rank), nil
}

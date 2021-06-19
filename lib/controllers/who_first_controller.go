package controllers

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

func WhoFirstArtist(artistInput model.ArtistInput, settings *model.WhoKnowsSettings, whoLast *bool) (*model.WhoFirstArtistResponse, error) {
	indexingService := indexing.CreateService()
	analysisService := analysis.CreateService()

	artist, err := indexingService.GetArtist(artistInput, false)

	if err != nil {
		return nil, err
	}

	var whoLastArg bool

	if whoLast == nil {
		whoLastArg = false
	} else {
		whoLastArg = *whoLast
	}

	whoFirst, err := analysisService.WhoFirstArtist(artist, settings, whoLastArg)

	if err != nil {
		return nil, err
	}

	return presenters.PresentWhoFirstArtistResponse(artist, whoFirst), nil
}

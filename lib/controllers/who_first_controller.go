package controllers

import (
	"github.com/jivison/gowon-indexer/lib/db"
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

	var excludeIDs []int64
	var undated []db.Scrobble

	if whoLast == nil || !*whoLast {
		undated, err = analysisService.WhoHasUndatedArtist(artist, settings, whoLastArg)

		if err != nil {
			return nil, err
		}

		for _, excludedUser := range undated {
			excludeIDs = append(excludeIDs, excludedUser.UserID)
		}
	}

	whoFirst, err := analysisService.WhoFirstArtist(artist, settings, whoLastArg, excludeIDs)

	if err != nil {
		return nil, err
	}

	return presenters.PresentWhoFirstArtistResponse(artist, whoFirst, undated), nil
}

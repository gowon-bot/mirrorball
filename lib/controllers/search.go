package controllers

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
)

func SearchArtist(criteria model.ArtistSearchCriteria, settings *model.SearchSettings) (*model.ArtistSearchResults, error) {
	analysisService := analysis.CreateService()

	artists, err := analysisService.SearchArtist(criteria, settings)

	if err != nil {
		return nil, err
	}

	return presenters.PresentArtistSearchResults(artists), nil
}

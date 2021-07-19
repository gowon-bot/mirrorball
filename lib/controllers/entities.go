package controllers

import (
	"context"

	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

func Artists(ctx context.Context, inputs []*model.ArtistInput, tag *model.TagInput, requireTagsForMissing *bool) ([]*model.Artist, error) {
	indexingService := indexing.CreateService()

	if inputs != nil {
		if requireTagsForMissing != nil && *requireTagsForMissing {
			analysisService := analysis.CreateService()

			analysisService.RequireTagsForMissing(inputs)
		}

		artists, err := indexingService.GetArtists(inputs, ctx)

		if err != nil {
			return nil, err
		}

		return presenters.PresentArtists(artists), nil
	}

	return []*model.Artist{}, nil
}

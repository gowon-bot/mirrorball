package controllers

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
)

type ArtistsTagsMap = map[int][]int

func TagArtists(artists []*model.ArtistInput, tags []*model.TagInput, markAsChecked *bool) (*string, error) {
	analysisService := analysis.CreateService()

	err := analysisService.TagArtists(artists, tags)

	if err != nil {
		return nil, err
	} else {
		var artistNames []string

		for _, artist := range artists {
			if artist.Name != nil {
				artistNames = append(artistNames, *artist.Name)
			}
		}

		if markAsChecked != nil && *markAsChecked {
			dbhelpers.UpdateManyArtistsToBeChecked(artistNames, constants.ChunkSize)
		}
	}

	return nil, nil
}

func Tags(settings *model.TagsSettings, requireTagsForMissing *bool) (*model.TagsResponse, error) {
	analysisService := analysis.CreateService()

	if requireTagsForMissing != nil && *requireTagsForMissing {
		analysisService.RequireTagsForMissing(settings.Artists)
	} else {
		go analysisService.RequireTagsForMissing(settings.Artists)
	}

	tags, err := analysisService.GetTags(settings)

	if err != nil {
		return nil, err
	}

	count, err := analysisService.CountTags(settings)

	if err != nil {
		return nil, err
	}

	return presenters.PresentTagsResponse(tags, count), nil
}

package controllers

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

func TagArtists(artists []*model.ArtistInput, tags []*model.TagInput) (*string, error) {
	indexingService := indexing.CreateService()

	var artistNames []string
	var tagNames []string

	for _, artist := range artists {
		if artist.Name != nil {
			artistNames = append(artistNames, *artist.Name)
		}
	}

	for _, tag := range tags {
		if tag.Name != nil {
			tagNames = append(tagNames, *tag.Name)
		}
	}

	if len(artistNames) == 0 || len(tagNames) == 0 {
		return nil, nil
	}

	artistsMap, err := indexingService.ConvertArtists(artistNames)

	if err != nil {
		return nil, err
	}

	tagsMap, err := indexingService.ConvertTags(tagNames)

	if err != nil {
		return nil, err
	}

}

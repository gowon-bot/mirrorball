package controllers

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

type ArtistsTagsMap = map[int][]int

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

	artistTags := generateArtistTagsToCreate(artistsMap, tagsMap)

	err = dbhelpers.InsertUniqueArtistTags(artistTags)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func generateArtistTagsToCreate(artistsMap indexing.ArtistsMap, tagsMap indexing.TagsMap) []db.ArtistTag {
	var artistTags []db.ArtistTag

	for _, artist := range artistsMap {
		for _, tag := range tagsMap {
			artistTags = append(artistTags, db.ArtistTag{ArtistID: artist.ID, TagID: tag.ID})
		}
	}

	return artistTags
}

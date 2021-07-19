package analysis

import (
	"log"
	"sync"

	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

func (a Analysis) TagArtists(artists []*model.ArtistInput, tags []*model.TagInput) error {
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
		return nil
	}

	artistsMap, err := a.indexingService.ConvertArtists(artistNames)

	if err != nil {
		return err
	}

	tagsMap, err := a.indexingService.ConvertTags(tagNames)

	if err != nil {
		return err
	}

	artistTags := a.generateArtistTagsToCreate(artistsMap, tagsMap)

	err = dbhelpers.InsertUniqueArtistTags(artistTags)

	if err != nil {
		return err
	}

	return nil
}

type TagResponse struct {
	Name        string
	Occurrences int
}

func (a Analysis) GetTags(settings *model.TagsSettings) ([]TagResponse, error) {
	var tags []TagResponse

	query := db.Db.Model((*db.Tag)(nil)).Column(`tag.name`).ColumnExpr("count(*) as occurrences").Group(`tag.name`).OrderExpr("2 desc, 1")

	query = inputparser.CreateParser(query).ParseTagsSettings(settings).GetQuery()

	err := query.Select(&tags)

	if err != nil {
		return tags, customerrors.DatabaseUnknownError()
	}

	return tags, nil
}

func (a Analysis) CountTags(settings *model.TagsSettings) (int, error) {
	query := db.Db.Model((*db.Tag)(nil))

	query = inputparser.CreateParser(query).ParseTagsSettings(settings).GetQuery()

	var count int

	err := query.ColumnExpr("count(distinct tag.name)").Select(&count)

	if err != nil {
		return 0, customerrors.DatabaseUnknownError()
	}

	return count, nil
}

func (a Analysis) generateArtistTagsToCreate(artistsMap indexing.ArtistsMap, tagsMap indexing.TagsMap) []db.ArtistTag {
	var artistTags []db.ArtistTag

	for _, artist := range artistsMap {
		for _, tag := range tagsMap {
			artistTags = append(artistTags, db.ArtistTag{ArtistID: artist.ID, TagID: tag.ID})
		}
	}

	return artistTags
}

func (a Analysis) RequireTagsForMissing(artistInputs []*model.ArtistInput) error {
	var artistNames []string

	for _, artist := range artistInputs {
		if artist.Name != nil {
			artistNames = append(artistNames, *artist.Name)
		}
	}

	artists, err := dbhelpers.SelectArtistsWhereInMany(artistNames, constants.ChunkSize)

	if err != nil {
		return err
	}

	var artistsThatNeedTags []string

	for _, artist := range artists {
		if !artist.CheckedForTags {

			artistsThatNeedTags = append(artistsThatNeedTags, artist.Name)
		}
	}

	parallelization := 5

	log.Print("Caching...", len(artistsThatNeedTags))
	if len(artistsThatNeedTags) == 0 {
		return nil
	} else if len(artistsThatNeedTags) < parallelization {
		for _, artist := range artistsThatNeedTags {
			a.CacheTagsForArtist(artist)
		}
	} else {
		artistChannel := make(chan string)
		var wg sync.WaitGroup
		wg.Add(parallelization)

		for ii := 0; ii < parallelization; ii++ {
			go func(c chan string) {
				for {
					artist, more := <-c
					log.Print("Processing " + artist)
					if !more {
						wg.Done()
						return
					}

					log.Print("Caching for " + artist)
					a.CacheTagsForArtist(artist)
				}
			}(artistChannel)
		}

		for _, artist := range artistsThatNeedTags {
			log.Print("Adding " + artist + " to the channel")
			artistChannel <- artist
		}

		close(artistChannel)
		wg.Wait()
	}

	dbhelpers.UpdateManyArtistsToBeChecked(artistNames, constants.ChunkSize)

	return nil
}

func (a Analysis) CacheTagsForArtist(artistName string) {
	_, artistInfo := a.lastFMService.ArtistInfo(artistName)

	inputArtistName := artistName

	artistInput := []*model.ArtistInput{{Name: &inputArtistName}}

	var tagNames []*model.TagInput

	for _, tag := range artistInfo.Artist.Tags.Tag {
		tagName := tag.Name

		tagNames = append(tagNames, &model.TagInput{Name: &tagName})
	}

	a.TagArtists(artistInput, tagNames)
}

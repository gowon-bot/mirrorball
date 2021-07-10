package rateyourmusic

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

// RateYourMusic holds methods for interacting with the cached rateyourmusic data
type RateYourMusic struct {
	indexingService *indexing.Indexing
}

func (rym RateYourMusic) GetRatings(settings *model.RatingsSettings) ([]db.Rating, error) {
	var ratings []db.Rating

	query := db.Db.Model(&ratings).Relation("RateYourMusicAlbum").Order("rating DESC", "release_year ASC", "title DESC")

	parser := inputparser.CreateParser(query).ParseRatingsSettings(settings)

	err := parser.GetQuery().Select()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return ratings, nil
}

func (rym RateYourMusic) CountRatings(settings *model.RatingsSettings) (int, error) {
	query := db.Db.Model((*db.Rating)(nil))

	parser := inputparser.CreateParser(query).ParseRatingsSettings(settings)

	count, err := parser.GetQuery().Count()

	if err != nil {
		return 0, customerrors.DatabaseUnknownError()
	}

	return count, nil
}

func (rym RateYourMusic) GetArtist(keywords string) (*db.RateYourMusicAlbum, error) {
	album := new(db.RateYourMusicAlbum)

	err := db.Db.Model(album).Where("artist_name ILIKE ?", keywords).WhereOr("artist_native_name ILIKE ?", keywords).Limit(1).Select()

	if err != nil {
		return album, customerrors.EntityDoesntExistError("artist")
	}

	return album, nil
}

// CreateService creates an instance of the indexing service object
func CreateService() *RateYourMusic {
	service := &RateYourMusic{
		indexingService: indexing.CreateService(),
	}

	return service
}

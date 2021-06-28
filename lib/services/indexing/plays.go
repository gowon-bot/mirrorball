package indexing

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
)

func (i Indexing) GetPlays(playsInput model.PlaysInput, pageInput *model.PageInput) ([]db.Play, error) {
	var plays []db.Play

	query := db.Db.Model(&plays).Relation("User").Relation("Track").Relation("Track.Artist").Relation("Track.Album")

	query = inputparser.CreateParser(query).ParsePlaysInput(&playsInput, &inputparser.InputSettings{
		AlbumPath:   "track__album",
		ArtistPath:  "track__artist",
		DefaultSort: "playcount desc",
	}).ParsePageInput(pageInput).GetQuery()

	err := query.Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

func (i Indexing) GetArtistPlays(user db.User, settings *model.ArtistPlaysSettings) ([]db.ArtistCount, error) {
	var plays []db.ArtistCount

	query := db.Db.Model(&plays).Relation("User").Relation("Artist").Where("user_id = ?", user.ID)
	parser := inputparser.CreateParser(query)

	parser.ParseArtistPlaysSettings(settings, inputparser.InputSettings{DefaultSort: "playcount desc"})

	err := query.Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

func (i Indexing) GetAlbumPlays(user db.User, settings *model.AlbumPlaysSettings) ([]db.AlbumCount, error) {
	var plays []db.AlbumCount

	query := db.Db.Model(&plays).Relation("User").Relation("Album").Relation("Album.Artist").Where("user_id = ?", user.ID)

	parser := inputparser.CreateParser(query).ParseAlbumPlaysSettings(settings, inputparser.InputSettings{
		DefaultSort: "playcount desc",
		ArtistPath:  "\"album__artist\"",
	})

	err := parser.GetQuery().Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

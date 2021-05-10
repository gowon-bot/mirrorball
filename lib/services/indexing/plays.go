package indexing

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func (i Indexing) GetPlays(user db.User, pageInput *model.PageInput) ([]db.Play, error) {
	var plays []db.Play

	query := db.Db.Model(&plays).Relation("User").Relation("Track").Relation("Track.Artist").Relation("Track.Album").Where("user_id = ?", user.ID).Order("scrobbled_at desc")
	query = ParsePageInput(query, pageInput)

	err := query.Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

func (i Indexing) GetArtistPlays(user db.User, settings *model.ArtistPlaysSettings) ([]db.ArtistCount, error) {
	var plays []db.ArtistCount

	query := db.Db.Model(&plays).Relation("User").Relation("Artist").Where("user_id = ?", user.ID)

	query = ParseArtistPlaysSettings(query, settings)

	err := query.Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

func (i Indexing) GetAlbumPlays(user db.User, settings *model.AlbumPlaysSettings) ([]db.AlbumCount, error) {
	var plays []db.AlbumCount

	query := db.Db.Model(&plays).Relation("User").Relation("Album").Relation("Album.Artist").Where("user_id = ?", user.ID)

	query = ParseAlbumPlaysSettings(query, settings)

	err := query.Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

func (i Indexing) GetTrackPlays(user db.User, settings *model.AlbumPlaysSettings) ([]db.TrackCount, error) {
	var plays []db.TrackCount

	query := db.Db.Model(&plays).Relation("User").Relation("Album").Relation("Album.Artist").Where("user_id = ?", user.ID)

	query = ParseAlbumPlaysSettings(query, settings)

	err := query.Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

package analysis

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// WhoKnowsArtist returns a list of who has listened to an artist
func (a Analysis) WhoKnowsArtist(artist *db.Artist, settings *model.WhoKnowsSettings) ([]db.ArtistCount, error) {
	var whoKnows []db.ArtistCount

	query := db.Db.Model(&whoKnows).
		Relation("Artist").
		Relation("User").
		Where("artist_id = ?", artist.ID).
		Order("playcount desc", "username desc")

	err := ParseWhoKnowsSettings(query, settings).Select()

	if err != nil {
		return whoKnows, customerrors.DatabaseUnknownError()
	}

	return whoKnows, nil
}

// WhoKnowsAlbum returns a list of who has listened to an album
func (a Analysis) WhoKnowsAlbum(album *db.Album, settings *model.WhoKnowsSettings) ([]db.AlbumCount, error) {
	var whoKnows []db.AlbumCount

	query := db.Db.Model(&whoKnows).
		Relation("Album").
		Relation("User").
		Where("album_id = ?", album.ID).
		Order("playcount desc", "username desc")

	err := ParseWhoKnowsSettings(query, settings).Select()

	if err != nil {
		return whoKnows, customerrors.DatabaseUnknownError()
	}

	return whoKnows, nil
}

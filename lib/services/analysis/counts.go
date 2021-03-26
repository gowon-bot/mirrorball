package analysis

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
)

func (a Analysis) ArtistTopAlbums(userID int64, artistID int64) ([]db.AlbumCount, error) {
	var topAlbums []db.AlbumCount

	err := db.Db.Model(&topAlbums).Relation("Album").Where("user_id = ?", userID).Where("artist_id = ?", artistID).Order("playcount desc").Select()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return topAlbums, nil
}

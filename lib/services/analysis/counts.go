package analysis

import (
	"github.com/go-pg/pg/v10"
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

type AmbiguousTrackCount struct {
	Playcount int
	Name      string
}

func (a Analysis) ArtistTopTracks(userID int64, artistID int64) ([]AmbiguousTrackCount, error) {
	var topTracks []AmbiguousTrackCount

	err := db.Db.Model((*db.TrackCount)(nil)).
		Relation("Track._").
		ColumnExpr("sum(playcount) as playcount").
		Column("name").
		Where("user_id = ?", userID).
		Where("artist_id = ?", artistID).
		Order("playcount desc").
		Group("name").
		Select(&topTracks)

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return topTracks, nil
}

func (a Analysis) AlbumTopTracks(userID int64, albumID int64) ([]AmbiguousTrackCount, error) {
	var topTracks []AmbiguousTrackCount

	err := db.Db.Model((*db.TrackCount)(nil)).
		Relation("Track._").
		ColumnExpr("sum(playcount) as playcount").
		Column("name").
		Where("user_id = ?", userID).
		Where("album_id = ?", albumID).
		Order("playcount desc").
		Group("name").
		Select(&topTracks)

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return topTracks, nil
}

func (a Analysis) TrackTopAlbums(userID int64, trackIDs []int64) ([]db.TrackCount, error) {
	var topTracks []db.TrackCount

	err := db.Db.Model(&topTracks).
		Relation("Track").
		Relation("Track.Album").
		Where("user_id = ?", userID).
		Where("track_id IN (?)", pg.In(trackIDs)).
		Order("playcount desc").
		Select()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return topTracks, nil
}

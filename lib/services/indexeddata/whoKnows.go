package indexeddata

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/jivison/gowon-indexer/lib/db"
)

// WhoKnowsArtist returns a list of users who have scrobbled an artist
func (id IndexedQuery) WhoKnowsArtist(artist *db.Artist) []db.ArtistCount {
	var whoKnows []db.ArtistCount

	db.Db.Model(&whoKnows).
		Relation("Artist").
		Relation("User").
		Where("artist_id=?", artist.ID).
		Order("playcount desc", "last_fm_username desc").
		Select()

	return whoKnows
}

// WhoKnowsAlbum returns a list of users who have scrobbled an album
func (id IndexedQuery) WhoKnowsAlbum(album *db.Album) []db.AlbumCount {
	var whoKnows []db.AlbumCount

	db.Db.Model(&whoKnows).
		Relation("Album").
		Relation("User").
		Where("album_id=?", album.ID).
		Order("playcount desc", "last_fm_username desc").
		Select()

	return whoKnows
}

// WhoKnowsTrack is a struct type containing a user and their playcounts
type WhoKnowsTrack struct {
	User      db.User
	Playcount int64
}

// WhoKnowsTrack returns a list of users who have scrobbled an track
func (id IndexedQuery) WhoKnowsTrack(tracks []db.Track) []WhoKnowsTrack {
	var whoKnows []struct {
		UserID         int64
		LastFMUsername string
		Playcount      int64
	}

	var ids []int64

	for _, track := range tracks {
		ids = append(ids, track.ID)
	}

	trackCounts := db.Db.Model((*db.TrackCount)(nil)).
		Join("JOIN tracks ON tracks.id = track_id").
		Column("name", "artist_id", "user_id", "tracks.name").
		ColumnExpr("sum(playcount) as total_playcount").
		Where("track_id IN (?)", pg.In(ids)).
		Group("name", "artist_id", "user_id")

	err := db.Db.Model(&whoKnows).
		TableExpr("(?) as track_counts", trackCounts).
		ColumnExpr("total_playcount as playcount").
		ColumnExpr("users.id as user_id").
		ColumnExpr("users.last_fm_username as last_fm_username").
		Join("JOIN users ON users.id = user_id").
		Order("total_playcount desc", "last_fm_username desc").
		Select(&whoKnows)

	if err != nil {
		log.Panic(err)
	}

	var whoKnowsTracks []WhoKnowsTrack

	for _, track := range whoKnows {
		whoKnowsTracks = append(whoKnowsTracks, WhoKnowsTrack{
			Playcount: track.Playcount,
			User: db.User{
				LastFMUsername: track.LastFMUsername,
				ID:             track.UserID,
			},
		})
	}

	return whoKnowsTracks
}

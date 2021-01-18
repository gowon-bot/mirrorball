package indexeddata

import (
	"github.com/go-pg/pg/v10"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/services/user"
)

// IndexedQuery holds methods for analysing the cached Last.fm data
type IndexedQuery struct {
	userService *user.User
}

// UserTopArtists returns a user's top artists
func (id IndexedQuery) UserTopArtists(username string) int {
	_, err := id.userService.GetUser(username)

	if err != nil {
		// return nil, err
	}

	var count int

	_, err =
		db.Db.Model((*db.Scrobble)(nil)).
			QueryOne(pg.Scan(&count), `
			SELECT count(*)
			FROM ?TableName as ?TableAlias
		`)

	return count
}

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

// CreateIndexedQueryService creates an instance of the lastfm indexed data service object
func CreateIndexedQueryService() *IndexedQuery {
	service := &IndexedQuery{
		userService: user.CreateService(),
	}

	return service
}

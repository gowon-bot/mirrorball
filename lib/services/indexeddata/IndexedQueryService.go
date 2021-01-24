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

// CreateIndexedQueryService creates an instance of the lastfm indexed data service object
func CreateIndexedQueryService() *IndexedQuery {
	service := &IndexedQuery{
		userService: user.CreateService(),
	}

	return service
}

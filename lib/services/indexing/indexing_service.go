package indexing

import (
	"time"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"

	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// Indexing holds methods for indexing users
type Indexing struct {
	lastFMService *lastfm.API
}

// AddPlay saves a play to the database
func (i Indexing) AddPlay(user *db.User, track *db.Track, scrobbledAt time.Time) (*db.Scrobble, error) {
	scrobble := &db.Scrobble{
		UserID: user.ID,
		User:   user,

		TrackID: track.ID,
		Track:   track,

		ScrobbledAt: scrobbledAt,
	}

	_, err := db.Db.Model(scrobble).Insert()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return scrobble, nil
}

// CreateService creates an instance of the indexing service object
func CreateService() *Indexing {
	service := &Indexing{
		lastFMService: lastfm.CreateAPIService(),
	}

	return service
}

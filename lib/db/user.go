package db

import (
	"log"
)

// AddScrobble saves a scrobble to the database
func (u User) AddScrobble(track *Track) *Scrobble {
	scrobble := &Scrobble{
		UserID: u.ID,
		User:   &u,

		TrackID: track.ID,
		Track:   track,
	}

	_, err := Db.Model(scrobble).Insert()

	if err != nil {
		log.Panic(err)
	}

	return scrobble
}

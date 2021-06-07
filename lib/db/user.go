package db

import (
	"time"

	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// SetLastIndexed sets a user's last indexed time
func (u User) SetLastIndexed(to time.Time) {
	_, err := Db.Model(&u).Set("last_indexed = ?", to).WherePK().Update()
	if err != nil {
		panic(err)
	}
}

// IsWavyUser returns whether a user has a Wavy user type
func (u User) IsWavyUser() bool {
	return u.UserType == "Wavy"
}

func (u User) AsRequestable() lastfm.Requestable {
	return lastfm.Requestable{Username: u.Username, Session: u.LastFMSession}
}

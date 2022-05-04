package db

import (
	"strings"
	"time"

	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// Possible values
// - PRIVATE (1)
// - DISCORD (2)
// - FMUSERNAME (3)
// - BOTH (4)
// - UNSET (5) - used for displaying help messages for people who haven't set it
const DefaultPrivacy = 5

func ConvertPrivacyToString(privacy int64) string {
	switch privacy {
	case 1:
		return "PRIVATE"
	case 2:
		return "DISCORD"
	case 3:
		return "FMUSERNAME"
	case 4:
		return "BOTH"
	default:
		return "UNSET"
	}
}

func ConvertPrivacyFromString(privacy string) int64 {
	switch strings.ToUpper(privacy) {
	case "PRIVATE":
		return 1
	case "DISCORD":
		return 2
	case "FMUSERNAME":
		return 3
	case "BOTH":
		return 4
	default:
		return 5
	}
}

// SetLastIndexed sets a user's last indexed time
func (u User) SetLastIndexed(to time.Time) {
	_, err := Db.Model(&u).Set("last_indexed = ?", to).WherePK().Update()
	if err != nil {
		panic(err)
	}
}

func (u User) AsRequestable() lastfm.Requestable {
	return lastfm.Requestable{Username: u.Username, Session: u.LastFMSession}
}

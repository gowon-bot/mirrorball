package indexing

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/services/indexeddata"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// Indexing holds methods for indexing users
type Indexing struct {
	lastFMService  *lastfm.API
	indexedService *indexeddata.IndexedMutation
}

// FullIndex indexes a user for the first time
func (i Indexing) FullIndex(user *db.User) {

	_, recentTracks := i.lastFMService.RecentTracks(lastfm.RecentTracksParams{
		Username: user.LastFMUsername,
		Period:   "overall",
		Limit:    1000,
	})

	for _, track := range recentTracks.RecentTracks.Tracks {
		var album *string

		if track.Album.Text != "" {
			album = &track.Album.Text
		}

		savedTrack, _ := i.indexedService.GetTrack(track.Name, track.Artist.Text, album, true)

		user.AddScrobble(savedTrack)
	}
}

// CreateService creates an instance of the response service object
func CreateService() *Indexing {

	service := &Indexing{
		lastFMService:  lastfm.CreateAPIService(),
		indexedService: indexeddata.CreateIndexedMutationService(),
	}

	return service
}

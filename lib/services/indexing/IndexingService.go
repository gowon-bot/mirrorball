package indexing

import (
	"strconv"

	"github.com/jivison/gowon-indexer/lib/db"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/api"
	"github.com/jivison/gowon-indexer/lib/services/indexeddata"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// Indexing holds methods for indexing users
type Indexing struct {
	lastFMService  *lastfm.API
	indexedService *indexeddata.IndexedMutation
}

// Update updates the index with new scrobbles
// func (i Indexing) Update(user *db.User) {

// 	_, recentTracks := i.lastFMService.RecentTracks(lastfm.RecentTracksParams{
// 		Username: user.LastFMUsername,
// 		Period:   "overall",
// 		Limit:    1000,
// 	})

// 	for _, track := range recentTracks.RecentTracks.Tracks {
// 		var album *string

// 		if track.Album.Text != "" {
// 			album = &track.Album.Text
// 		}

// 		savedTrack, _ := i.indexedService.GetTrack(track.Name, track.Artist.Text, album, true)

// 		user.AddScrobble(savedTrack)
// 	}
// }

// FullIndex indexes a user for the first time
func (i Indexing) FullIndex(user *db.User) {
	i.FullArtistCountIndex(user)
}

// FullArtistCountIndex fully indexes a users topartists
func (i Indexing) FullArtistCountIndex(user *db.User) {
	i.ResetArtistCounts(user)

	params := lastfm.TopArtistParams{
		Username: user.LastFMUsername,
		Limit:    1000,
		Page:     1,
	}

	var topArtists []lastfm.TopArtist

	_, firstPage := i.lastFMService.TopArtists(params)

	topArtists = append(topArtists, firstPage.TopArtists.Artists...)

	totalPages, _ := strconv.Atoi(firstPage.TopArtists.Attributes.TotalPages)

	paginator := helpers.Paginator{
		PageSize:      params.Limit,
		TotalPages:    totalPages,
		SkipFirstPage: true,

		Function: func(pp helpers.PagedParams) {
			params.Page = pp.Page

			_, response := i.lastFMService.TopArtists(params)

			topArtists = append(topArtists, response.TopArtists.Artists...)
		},
	}

	paginator.GetAll()

	for _, topArtist := range topArtists {
		artist, _ := i.indexedService.GetArtist(topArtist.Name, true)

		playcount, _ := strconv.Atoi(topArtist.Playcount)

		i.indexedService.IncrementArtistCount(artist, user, int32(playcount))
	}
}

// ResetArtistCounts deletes all of a user's artist counts
func (i Indexing) ResetArtistCounts(user *db.User) {
	db.Db.Model((*db.ArtistCount)(nil)).Where("user_id=?", user.ID).Delete()
}

// CreateService creates an instance of the response service object
func CreateService() *Indexing {

	service := &Indexing{
		lastFMService:  lastfm.CreateAPIService(),
		indexedService: indexeddata.CreateIndexedMutationService(),
	}

	return service
}

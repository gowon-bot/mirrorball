package indexing

import (
	"strconv"
	"time"

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
func (i Indexing) Update(user *db.User) {
	var tracks []lastfm.RecentTrack

	params := lastfm.RecentTracksParams{
		Username: user.LastFMUsername,
		Period:   "overall",
		Limit:    1000,
		From:     strconv.FormatInt(user.LastIndexed.Unix(), 10),
	}

	_, recentTracks := i.lastFMService.RecentTracks(params)

	tracks = append(tracks, recentTracks.RecentTracks.Tracks...)

	if recentTracks.RecentTracks.Attributes.Total == "0" {
		return
	}

	if totalPages, _ := strconv.Atoi(recentTracks.RecentTracks.Attributes.TotalPages); totalPages > 1 {
		paginator := helpers.Paginator{
			PageSize:      1000,
			TotalPages:    totalPages,
			SkipFirstPage: true,
			Function: func(pp helpers.PagedParams) {
				params.Page = pp.Page
				_, response := i.lastFMService.RecentTracks(params)

				tracks = append(tracks, response.RecentTracks.Tracks...)
			},
		}

		paginator.GetAll()
	}

	for _, track := range tracks {
		if track.Attributes.IsNowPlaying == "true" {
			continue
		}

		var album *string

		if track.Album.Text != "" {
			album = &track.Album.Text
		}

		savedTrack, _ := i.indexedService.GetTrack(track.Name, track.Artist.Text, album, true)

		timestamp, _ := helpers.ParseUnix(track.Timestamp.UTS)

		i.AddScrobble(user, savedTrack, timestamp)
		i.indexedService.IncrementArtistCount(savedTrack.Artist, user, 1)
		i.indexedService.IncrementAlbumCount(savedTrack.Album, user, 1)
	}

	lastTrack := tracks[len(tracks)-1]
	lastTimestamp, _ := helpers.ParseUnix(lastTrack.Timestamp.UTS)

	user.SetLastIndexed(lastTimestamp)
}

// FullIndex indexes a user for the first time
func (i Indexing) FullIndex(user *db.User) {
	i.FullArtistCountIndex(user)
	i.FullAlbumCountIndex(user)

	lastScrobbled := i.lastFMService.LastScrobbledTimestamp(user.LastFMUsername)

	user.SetLastIndexed(lastScrobbled.Add(time.Second))
}

// FullArtistCountIndex fully indexes a users top artists
func (i Indexing) FullArtistCountIndex(user *db.User) {
	i.ResetArtistCounts(user)

	params := lastfm.TopEntityParams{
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

// FullAlbumCountIndex fully indexes a users top albums
func (i Indexing) FullAlbumCountIndex(user *db.User) {
	i.ResetAlbumCounts(user)

	params := lastfm.TopEntityParams{
		Username: user.LastFMUsername,
		Limit:    1000,
		Page:     1,
	}

	var topAlbums []lastfm.TopAlbum

	_, firstPage := i.lastFMService.TopAlbums(params)

	topAlbums = append(topAlbums, firstPage.TopAlbums.Albums...)

	totalPages, _ := strconv.Atoi(firstPage.TopAlbums.Attributes.TotalPages)

	paginator := helpers.Paginator{
		PageSize:      params.Limit,
		TotalPages:    totalPages,
		SkipFirstPage: true,

		Function: func(pp helpers.PagedParams) {
			params.Page = pp.Page

			_, response := i.lastFMService.TopAlbums(params)

			topAlbums = append(topAlbums, response.TopAlbums.Albums...)
		},
	}

	paginator.GetAll()

	for _, topAlbum := range topAlbums {
		album, _ := i.indexedService.GetAlbum(topAlbum.Name, topAlbum.Artist.Name, true)

		playcount, _ := strconv.Atoi(topAlbum.Playcount)

		i.indexedService.IncrementAlbumCount(album, user, int32(playcount))
	}
}

// ResetArtistCounts deletes all of a user's artist counts
func (i Indexing) ResetArtistCounts(user *db.User) {
	db.Db.Model((*db.ArtistCount)(nil)).Where("user_id=?", user.ID).Delete()
}

// ResetAlbumCounts deletes all of a user's album counts
func (i Indexing) ResetAlbumCounts(user *db.User) {
	db.Db.Model((*db.AlbumCount)(nil)).Where("user_id=?", user.ID).Delete()
}

// AddScrobble saves a scrobble to the database
func (i Indexing) AddScrobble(user *db.User, track *db.Track, timestamp time.Time) (*db.Scrobble, error) {
	scrobble := &db.Scrobble{
		UserID: user.ID,
		User:   user,

		TrackID: track.ID,
		Track:   track,

		Timestamp: timestamp,
	}

	_, err := db.Db.Model(scrobble).Insert()

	if err != nil {
		return nil, err
	}

	return scrobble, nil
}

// CreateService creates an instance of the response service object
func CreateService() *Indexing {

	service := &Indexing{
		lastFMService:  lastfm.CreateAPIService(),
		indexedService: indexeddata.CreateIndexedMutationService(),
	}

	return service
}

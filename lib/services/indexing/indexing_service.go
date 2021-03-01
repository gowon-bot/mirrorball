package indexing

import (
	"strconv"
	"time"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/api"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// Indexing holds methods for indexing users
type Indexing struct {
	lastFMService *lastfm.API
}

// FullIndex downloads all of a users data and caches it
func (i Indexing) FullIndex(user *db.User) error {
	startTime := time.Now()
	err := i.fullArtistCountIndex(user)

	if err != nil {
		return err
	}

	err = i.fullAlbumCountIndex(user)

	if err != nil {
		return err
	}

	err = i.fullTrackCountIndex(user)

	if err != nil {
		return err
	}

	user.SetLastIndexed(startTime)

	return nil
}

// Update updates the cache with the newest data
func (i Indexing) Update(user *db.User) error {
	err := i.updateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (i Indexing) fullArtistCountIndex(user *db.User) error {
	i.resetArtistCounts(user)

	topArtists, err := i.lastFMService.AllTopArtists(user.Username)

	if err != nil {
		return err
	}

	for _, topArtist := range topArtists {
		artist, _ := i.GetArtist(model.ArtistInput{Name: &topArtist.Name}, true)

		playcount, _ := strconv.Atoi(topArtist.Playcount)

		i.IncrementArtistCount(artist, user, int32(playcount))
	}

	return nil
}

func (i Indexing) fullAlbumCountIndex(user *db.User) error {
	i.resetAlbumCounts(user)

	topAlbums, err := i.lastFMService.AllTopAlbums(user.Username)

	if err != nil {
		return err
	}

	for _, topAlbum := range topAlbums {
		album, _ := i.GetAlbum(model.AlbumInput{
			Name:   &topAlbum.Name,
			Artist: &model.ArtistInput{Name: &topAlbum.Artist.Name},
		}, true)

		playcount, _ := strconv.Atoi(topAlbum.Playcount)

		i.IncrementAlbumCount(album, user, int32(playcount))
	}

	return nil
}

func (i Indexing) fullTrackCountIndex(user *db.User) error {
	i.resetTrackCounts(user)

	topTracks, err := i.lastFMService.AllTopTracks(user.Username)

	if err != nil {
		return err
	}

	for _, topTrack := range topTracks {
		track, _ := i.GetTrack(model.TrackInput{
			Name:   &topTrack.Name,
			Artist: &model.ArtistInput{Name: &topTrack.Artist.Name},
		}, true)

		playcount, _ := strconv.Atoi(topTrack.Playcount)

		i.IncrementTrackCount(track, user, int32(playcount))
	}

	return nil
}

func (i Indexing) updateUser(user *db.User) error {
	tracks, err := i.lastFMService.AllScrobblesSince(user.Username, user.LastIndexed)

	if err != nil {
		return nil
	}

	for _, track := range tracks {
		if track.Attributes.IsNowPlaying == "true" {
			continue
		}

		var album *string

		if track.Album.Text != "" {
			album = &track.Album.Text
		}

		cachedTrack, _ := i.GetTrack(model.TrackInput{
			Name:   &track.Name,
			Artist: &model.ArtistInput{Name: &track.Artist.Text},
			Album:  &model.AlbumInput{Name: album},
		}, true)

		timestamp, _ := helpers.ParseUnix(track.Timestamp.UTS)

		_, err := i.AddPlay(user, cachedTrack, timestamp)

		if err == nil {
			i.IncrementArtistCount(cachedTrack.Artist, user, 1)
			i.IncrementAlbumCount(cachedTrack.Album, user, 1)
			i.IncrementTrackCount(cachedTrack, user, 1)
		}
	}

	lastTrack := tracks[len(tracks)-1]
	lastTimestamp, _ := helpers.ParseUnix(lastTrack.Timestamp.UTS)

	user.SetLastIndexed(lastTimestamp)

	return nil
}

// IncrementArtistCount increments an artist's aggregated playcount by a given amount
func (i Indexing) IncrementArtistCount(artist *db.Artist, user *db.User, playcount int32) (*db.ArtistCount, error) {
	artistCount, err := i.GetArtistCount(artist, user, true)

	if err != nil {
		return nil, err
	}

	var newPlaycount int32

	_, err = db.Db.Model(artistCount).
		Set("playcount=?", playcount+artistCount.Playcount).
		Where("artist_id=?", artist.ID).
		Where("user_id=?", user.ID).
		Returning("playcount").
		Update(&newPlaycount)

	artistCount.Artist = artist
	artistCount.Playcount = newPlaycount

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return artistCount, nil
}

// IncrementAlbumCount increments an album's aggregated playcount by a given amount
func (i Indexing) IncrementAlbumCount(album *db.Album, user *db.User, count int32) (*db.AlbumCount, error) {

	albumCount, err := i.GetAlbumCount(album, user, true)

	if err != nil {
		return nil, err
	}

	var newPlaycount int32

	_, err = db.Db.Model(albumCount).
		Set("playcount=?", count+albumCount.Playcount).
		Where("album_id=?", album.ID).
		Where("user_id=?", user.ID).
		Returning("playcount").
		Update(&newPlaycount)

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	albumCount.Album = album
	albumCount.Playcount = newPlaycount

	return albumCount, nil
}

// IncrementTrackCount increments an track's aggregated playcount by a given amount
func (i Indexing) IncrementTrackCount(track *db.Track, user *db.User, count int32) (*db.TrackCount, error) {

	trackCount, err := i.GetTrackCount(track, user, true)

	if err != nil {
		return nil, err
	}

	var newPlaycount int32

	_, err = db.Db.Model(trackCount).
		Set("playcount=?", count+trackCount.Playcount).
		Where("track_id=?", track.ID).
		Where("user_id=?", user.ID).
		Returning("playcount").
		Update(&newPlaycount)

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	trackCount.Track = track
	trackCount.Playcount = newPlaycount

	return trackCount, nil
}

func (i Indexing) resetArtistCounts(user *db.User) {
	db.Db.Model((*db.ArtistCount)(nil)).Where("user_id=?", user.ID).Delete()
}

func (i Indexing) resetAlbumCounts(user *db.User) {
	db.Db.Model((*db.AlbumCount)(nil)).Where("user_id=?", user.ID).Delete()
}

func (i Indexing) resetTrackCounts(user *db.User) {
	db.Db.Model((*db.TrackCount)(nil)).Where("user_id=?", user.ID).Delete()
}

// AddPlay saves a play to the database
func (i Indexing) AddPlay(user *db.User, track *db.Track, scrobbledAt time.Time) (*db.Play, error) {
	scrobble := &db.Play{
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

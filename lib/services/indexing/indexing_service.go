package indexing

import (
	"strconv"
	"time"

	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	apihelpers "github.com/jivison/gowon-indexer/lib/helpers/api"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"

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

	i.resetPlays(user)
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

	var topArtistNames []string

	for _, artist := range topArtists {
		topArtistNames = append(topArtistNames, artist.Name)
	}

	artistMap, err := i.ConvertArtists(topArtistNames)

	if err != nil {
		return err
	}

	var artistCounts []*db.ArtistCount

	for _, topArtist := range topArtists {
		artist, _ := artistMap[topArtist.Name]
		playcount, _ := strconv.Atoi(topArtist.Playcount)

		artistCounts = append(artistCounts, &db.ArtistCount{Artist: &artist, ArtistID: artist.ID, User: user, UserID: user.ID, Playcount: int32(playcount)})
	}

	_, err = db.Db.Model(&artistCounts).Insert()

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}

func (i Indexing) fullAlbumCountIndex(user *db.User) error {
	i.resetAlbumCounts(user)

	topAlbums, err := i.lastFMService.AllTopAlbums(user.Username)

	if err != nil {
		return err
	}

	var topAlbumNames []AlbumToConvert

	for _, topAlbum := range topAlbums {
		topAlbumNames = append(topAlbumNames, AlbumToConvert{
			ArtistName: topAlbum.Artist.Name,
			AlbumName:  topAlbum.Name,
		})
	}

	albumMap, err := i.ConvertAlbums(topAlbumNames)

	if err != nil {
		return err
	}

	var albumCounts []*db.AlbumCount

	for _, topAlbum := range topAlbums {
		album, _ := albumMap[topAlbum.Artist.Name][topAlbum.Name]
		playcount, _ := strconv.Atoi(topAlbum.Playcount)

		albumCounts = append(albumCounts, &db.AlbumCount{
			Album:     &album,
			AlbumID:   album.ID,
			User:      user,
			UserID:    user.ID,
			Playcount: int32(playcount),
		})
	}

	_, err = db.Db.Model(&albumCounts).Insert()

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}

func (i Indexing) fullTrackCountIndex(user *db.User) error {
	i.resetTrackCounts(user)

	topTracks, err := i.lastFMService.AllTopTracks(user.Username)

	if err != nil {
		return err
	}

	var topTrackNames []TrackToConvert

	for _, topTrack := range topTracks {
		topTrackNames = append(topTrackNames, TrackToConvert{
			ArtistName: topTrack.Artist.Name,
			TrackName:  topTrack.Name,
		})
	}

	tracksMap, err := i.ConvertTracks(topTrackNames)

	if err != nil {
		return nil
	}

	var trackCounts []*db.TrackCount

	for _, topTrack := range topTracks {
		albumName := ""

		track, _ := tracksMap[topTrack.Artist.Name][albumName][topTrack.Name]
		playcount, _ := strconv.Atoi(topTrack.Playcount)

		trackCounts = append(trackCounts, &db.TrackCount{
			Track:     &track,
			TrackID:   track.ID,
			User:      user,
			UserID:    user.ID,
			Playcount: int32(playcount),
		})
	}

	_, err = db.Db.Model(&trackCounts).Insert()

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}

func (i Indexing) updateUser(user *db.User) error {
	recentTracks, err := i.lastFMService.AllScrobblesSince(user.Username, user.LastIndexed)

	if err != nil {
		return err
	} else if len(recentTracks) == 0 {
		return nil
	}

	var scrobbles []lastfm.RecentTrack

	if recentTracks[0].Attributes.IsNowPlaying == "true" {
		scrobbles = append(scrobbles, recentTracks[1:]...)
	} else {
		scrobbles = recentTracks
	}

	if len(scrobbles) < 1 {
		return nil
	}

	artistCounts, albumCounts, trackCounts, plays, err := i.GenerateCountsFromScrobbles(scrobbles, *user)

	err = i.createOrUpdateCounts(*user, artistCounts, albumCounts, trackCounts)

	if err != nil {
		return err
	}

	dbhelpers.InsertManyPlays(plays, constants.ChunkSize)

	firstTrack := scrobbles[0]
	lastTimestamp, _ := apihelpers.ParseUnix(firstTrack.Timestamp.UTS)

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

func (i Indexing) resetPlays(user *db.User) {
	db.Db.Model((*db.Play)(nil)).Where("user_id=?", user.ID).Delete()
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

func (i Indexing) createOrUpdateCounts(user db.User, artistCounts []db.ArtistCount, albumCounts []db.AlbumCount, trackCounts []db.TrackCount) error {
	_, err := dbhelpers.UpdateOrCreateManyArtistCounts(artistCounts, user.ID)

	if err != nil {
		return err
	}

	_, err = dbhelpers.UpdateOrCreateManyAlbumCounts(albumCounts, user.ID)

	if err != nil {
		return err
	}

	_, err = dbhelpers.UpdateOrCreateManyTrackCounts(trackCounts, user.ID)

	if err != nil {
		return err
	}

	return nil
}

// CreateService creates an instance of the indexing service object
func CreateService() *Indexing {
	service := &Indexing{
		lastFMService: lastfm.CreateAPIService(),
	}

	return service
}

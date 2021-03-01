package indexing

import (
	"strconv"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// Indexing holds methods for indexing users
type Indexing struct {
	lastFMService *lastfm.API
}

// FullIndex downloads all of a users data and caches it
func (i Indexing) FullIndex(user *db.User) error {
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

// CreateService creates an instance of the indexing service object
func CreateService() *Indexing {
	service := &Indexing{
		lastFMService: lastfm.CreateAPIService(),
	}

	return service
}

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
	return i.fullArtistCountIndex(user)
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

		i.incrementArtistCount(artist, user, int32(playcount))
	}

	return nil
}

func (i Indexing) incrementArtistCount(artist *db.Artist, user *db.User, playcount int32) (*db.ArtistCount, error) {
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

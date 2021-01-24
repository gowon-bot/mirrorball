package indexeddata

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// IndexedMutation holds methods for interacting with the cached Last.fm data
type IndexedMutation struct{}

// GetArtist returns (and optionally creates) an artist from the database
func (i IndexedMutation) GetArtist(artist string, create bool) (*db.Artist, error) {
	dbArtist := new(db.Artist)

	err := db.Db.Model(dbArtist).Where("name ILIKE ?", artist).Limit(1).Select()

	if err != nil && create == true {
		dbArtist = &db.Artist{
			Name: artist,
		}

		db.Db.Model(dbArtist).Insert()
	} else if err != nil {
		return nil, err
	}

	return dbArtist, nil
}

// GetAlbum returns (and optionally creates) an album from the database
func (i IndexedMutation) GetAlbum(album, artist string, create bool) (*db.Album, error) {
	dbAlbum := new(db.Album)

	err := db.Db.Model(dbAlbum).
		Relation("Artist").
		Where("album.name ILIKE ?", album).
		Where("artist.name ILIKE ?", artist).
		Limit(1).
		Select()

	if err != nil && create == true {
		dbArtist, _ := i.GetArtist(artist, true)

		dbAlbum = &db.Album{
			Name: album,

			ArtistID: dbArtist.ID,
			Artist:   dbArtist,
		}

		db.Db.Model(dbAlbum).Insert()
	} else if err != nil {
		return nil, err
	}

	return dbAlbum, nil
}

// GetTrack returns (and optionally creates) a track from the database
func (i IndexedMutation) GetTrack(track, artist string, album *string, create bool) (*db.Track, error) {

	dbTrack := new(db.Track)

	query := db.Db.Model(dbTrack).
		Relation("Artist").
		Relation("Album").
		Where("track.name ILIKE ?", track).
		Where("artist.name ILIKE ?", artist).
		Limit(1)

	if album != nil {
		query = query.Where("album.name ILIKE ?", album)
	}

	err := query.Select()

	if err != nil && create == true {
		dbTrack = i.SaveTrack(track, artist, album)
	}

	return dbTrack, nil
}

// GetTracks returns a list of tracks from the database
func (i IndexedMutation) GetTracks(track, artist string) ([]db.Track, error) {
	var dbTracks []db.Track

	query := db.Db.Model(&dbTracks).
		Relation("Artist").
		Where("track.name ILIKE ?", track).
		Where("artist.name ILIKE ?", artist)

	err := query.Select()

	if err != nil {
		return nil, nil
	}

	return dbTracks, nil
}

// SaveTrack saves a track in the database
func (i IndexedMutation) SaveTrack(trackName, artistName string, albumName *string) *db.Track {

	artist, _ := i.GetArtist(artistName, true)
	var album *db.Album = nil

	if albumName != nil {
		album, _ = i.GetAlbum(*albumName, artistName, true)
	}

	track := &db.Track{
		Name: trackName,
	}

	if album != nil {
		track.Album = album
		track.AlbumID = album.ID
	}

	if artist != nil {
		track.Artist = artist
		track.ArtistID = artist.ID
	}

	db.Db.Model(track).Insert()

	return track
}

// GetArtistCount gets and optionally creates an artist count
func (i IndexedMutation) GetArtistCount(artist *db.Artist, user *db.User, create bool) (*db.ArtistCount, error) {
	artistCount := new(db.ArtistCount)

	err := db.Db.Model(artistCount).Where("user_id=?", user.ID).Where("artist_id=?", artist.ID).Limit(1).Select()

	if err != nil && create == true {
		artistCount = &db.ArtistCount{
			UserID: user.ID,
			User:   user,

			ArtistID: artist.ID,
			Artist:   artist,
		}

		db.Db.Model(artistCount).Insert()
	} else if err != nil {
		return nil, err
	}

	return artistCount, nil
}

// GetAlbumCount gets and optionally creates an album count
func (i IndexedMutation) GetAlbumCount(album *db.Album, user *db.User, create bool) (*db.AlbumCount, error) {
	albumCount := new(db.AlbumCount)

	err := db.Db.Model(albumCount).Where("user_id=?", user.ID).Where("album_id=?", album.ID).Limit(1).Select()

	if err != nil && create == true {
		albumCount = &db.AlbumCount{
			UserID: user.ID,
			User:   user,

			AlbumID: album.ID,
			Album:   album,
		}

		db.Db.Model(albumCount).Insert()
	} else if err != nil {
		return nil, err
	}

	return albumCount, nil
}

// GetTrackCount gets and optionally creates an track count
func (i IndexedMutation) GetTrackCount(track *db.Track, user *db.User, create bool) (*db.TrackCount, error) {
	trackCount := new(db.TrackCount)

	err := db.Db.Model(trackCount).Where("user_id=?", user.ID).Where("track_id=?", track.ID).Limit(1).Select()

	if err != nil && create == true {
		trackCount = &db.TrackCount{
			UserID: user.ID,
			User:   user,

			TrackID: track.ID,
			Track:   track,
		}

		db.Db.Model(trackCount).Insert()
	} else if err != nil {
		return nil, err
	}

	return trackCount, nil
}

// IncrementArtistCount increments an artist's aggregated playcount by a given amount
func (i IndexedMutation) IncrementArtistCount(artist *db.Artist, user *db.User, count int32) *db.ArtistCount {

	artistCount, _ := i.GetArtistCount(artist, user, true)

	var newPlaycount int32

	db.Db.Model(artistCount).
		Set("playcount=?", count+artistCount.Playcount).
		Where("artist_id=?", artist.ID).
		Where("user_id=?", user.ID).
		Returning("playcount").
		Update(&newPlaycount)

	artistCount.Artist = artist
	artistCount.Playcount = newPlaycount

	return artistCount
}

// IncrementAlbumCount increments an album's aggregated playcount by a given amount
func (i IndexedMutation) IncrementAlbumCount(album *db.Album, user *db.User, count int32) *db.AlbumCount {

	albumCount, _ := i.GetAlbumCount(album, user, true)

	var newPlaycount int32

	db.Db.Model(albumCount).
		Set("playcount=?", count+albumCount.Playcount).
		Where("album_id=?", album.ID).
		Where("user_id=?", user.ID).
		Returning("playcount").
		Update(&newPlaycount)

	albumCount.Album = album
	albumCount.Playcount = newPlaycount

	return albumCount
}

// IncrementTrackCount increments an track's aggregated playcount by a given amount
func (i IndexedMutation) IncrementTrackCount(track *db.Track, user *db.User, count int32) *db.TrackCount {

	trackCount, _ := i.GetTrackCount(track, user, true)

	var newPlaycount int32

	_, err := db.Db.Model(trackCount).
		Set("playcount=?", count+trackCount.Playcount).
		Where("track_id=?", track.ID).
		Where("user_id=?", user.ID).
		Returning("playcount").
		Update(&newPlaycount)

	if err != nil {
		panic(err)
	}

	trackCount.Track = track
	trackCount.Playcount = newPlaycount

	return trackCount
}

// ConvertToGraphQLArtist converts a db artist to a gql artist
func ConvertToGraphQLArtist(artist *db.Artist) *model.Artist {
	if artist == nil {
		return nil
	}

	return &model.Artist{
		ID:   int(artist.ID),
		Name: artist.Name,
	}
}

// ConvertToGraphQLAlbum converts a db album to a gql album
func ConvertToGraphQLAlbum(album *db.Album) *model.Album {

	convertedAlbum := &model.Album{
		ID:   int(album.ID),
		Name: album.Name,
	}

	if album.Artist != nil {
		convertedAlbum.Artist = ConvertToGraphQLArtist(album.Artist)
	}

	return convertedAlbum
}

// ConvertToGraphQLTrack converts a db track to a gql track
func ConvertToGraphQLTrack(track *db.Track) *model.Track {
	if track == nil {
		return nil
	}

	convertedTrack := &model.Track{
		ID:   int(track.ID),
		Name: track.Name,
	}

	if track.Artist != nil {
		convertedTrack.Artist = ConvertToGraphQLArtist(track.Artist)
	}

	if track.Album != nil {
		convertedTrack.Album = ConvertToGraphQLAlbum(track.Album)
	}

	return convertedTrack
}

// AmbiguousTrack is the type for a track with no id or album
// this is because a track can belong to multiple albums
// and in cases where we need to aggregate those tracks,
// we can use AmbiguousTrack
type AmbiguousTrack struct {
	Name   string
	Artist *db.Artist
}

// ConvertToAmbiguousTrack converts a db track to a gql ambiguous track
func ConvertToAmbiguousTrack(track *AmbiguousTrack) *model.AmbiguousTrack {
	if track == nil {
		return nil
	}

	convertedTrack := &model.AmbiguousTrack{
		Name: track.Name,
	}

	if track.Artist != nil {
		convertedTrack.Artist = ConvertToGraphQLArtist(track.Artist)
	}

	return convertedTrack
}

// CreateIndexedMutationService creates an instance of the lastfm indexed service object
func CreateIndexedMutationService() *IndexedMutation {
	service := &IndexedMutation{}

	return service
}

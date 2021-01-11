package lastfm

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// Indexed holds methods for interacting with the cached Last.fm data
type Indexed struct{}

// GetArtist returns (and optionally creates) an artist from the database
func (i Indexed) GetArtist(artist string, create bool) (*db.Artist, error) {
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
func (i Indexed) GetAlbum(album, artist string, create bool) (*db.Album, error) {
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
func (i Indexed) GetTrack(track, artist string, album *string, create bool) (*db.Track, error) {
	dbTrack := new(db.Track)

	query := db.Db.Model(dbTrack).
		Relation("Artist").
		Where("track.name ILIKE ?", track).
		Where("artist.name ILIKE ?", artist).
		Limit(1)

	if album != nil {
		query = query.Where("album.name ILIKE ?", album)
	}

	err := query.Select()

	if err != nil && create == true {
		dbTrack = i.SaveTrack(track, artist, album)
	} else {
		return nil, err
	}

	return dbTrack, nil
}

// SaveTrack saves a track in the database
func (i Indexed) SaveTrack(artistName, trackName string, albumName *string) *db.Track {

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

// CreateIndexedService creates an instance of the lastfm indexed service object
func CreateIndexedService() *Indexed {
	service := &Indexed{}

	return service
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
	// var artist *model.Artist
	// if album.Artist != nil {
	// 	artist = ConvertToGraphQLArtist(album.Artist)
	// }

	return &model.Album{
		ID:   int(album.ID),
		Name: album.Name,
	}
}

// ConvertToGraphQLTrack converts a db track to a gql track
func ConvertToGraphQLTrack(track *db.Track) *model.Track {
	if track == nil {
		return nil
	}

	var artist *model.Artist
	if track.Artist != nil {
		artist = ConvertToGraphQLArtist(track.Artist)
	}

	var album *model.Album
	if track.Album != nil {
		album = ConvertToGraphQLAlbum(track.Album)
	}

	return &model.Track{
		ID:   int(track.ID),
		Name: track.Name,

		Artist: artist,
		Album:  album,
	}
}

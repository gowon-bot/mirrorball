package indexing

import "github.com/jivison/gowon-indexer/lib/db"

type ArtistsMap = map[string]db.Artist
type AlbumsMap = map[string]map[string]db.Album
type TracksMap = map[string]map[string]map[string]db.Track

type AlbumToCreate struct {
	ArtistID  int64
	AlbumName string
	Artist    *db.Artist
}

type ArtistToConvert = string

type TrackToConvert struct {
	ArtistName string
	TrackName  string
	AlbumName  *string
}

type AlbumToConvert struct {
	ArtistName string
	AlbumName  string
}

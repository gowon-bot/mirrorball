package indexing

import "github.com/jivison/gowon-indexer/lib/db"

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

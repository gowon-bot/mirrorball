package db

import "time"

// User is the database model for a last.fm user
type User struct {
	ID          int64
	DiscordID   string
	Username    string
	UserType    string
	LastIndexed time.Time
}

/*
* Indexing Models
 */

// Artist represents a cached artist
type Artist struct {
	ID   int64
	Name string
}

// Album represents a cached artist
type Album struct {
	ID   int64
	Name string

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`
}

// Track represents a cached artist
type Track struct {
	ID   int64
	Name string

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`

	AlbumID *int64
	Album   *Album `pg:"rel:has-one"`
}

// Play represents a single play of a song (eg. a Last.fm scrobble)
type Play struct {
	ScrobbledAt time.Time

	UserID int64
	User   *User `pg:"rel:has-one"`

	TrackID int64
	Track   *Track `pg:"rel:has-one"`
}

/*
* Aggregated Structures
 */

// ArtistCount represents aggregated artist plays
type ArtistCount struct {
	Playcount int32

	UserID int64
	User   *User `pg:"rel:has-one"`

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`
}

// AlbumCount represents aggregated album plays
type AlbumCount struct {
	Playcount int32

	UserID int64
	User   *User `pg:"rel:has-one"`

	AlbumID int64
	Album   *Album `pg:"rel:has-one"`
}

// TrackCount represents aggregated track plays
type TrackCount struct {
	Playcount int32

	UserID int64
	User   *User `pg:"rel:has-one"`

	TrackID int64
	Track   *Track `pg:"rel:has-one"`
}

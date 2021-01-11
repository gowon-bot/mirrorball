package db

// User is the database model for a last.fm user
type User struct {
	ID             int64  `json:"id"`
	LastFMUsername string `json:"lastFMUsername"`
}

/**
* Last.fm Structures
 */

// Artist represents an artist in Last.fm
type Artist struct {
	ID   int64
	Name string

	Albums *[]Album `pg:"rel:has-many"`
	Tracks *[]Track `pg:"rel:has-many"`
}

// Album represents an album in Last.fm
type Album struct {
	ID   int64
	Name string

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`

	Tracks *[]Track `pg:"rel:has-many"`
}

// Track represents a track in Last.fm
type Track struct {
	ID   int64
	Name string

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`

	AlbumID int64
	Album   *Album `pg:"rel:has-one"`
}

package db

import "time"

// User is the database model for a last.fm user
type User struct {
	ID            int64 `pg:",pk"`
	DiscordID     string
	Username      string
	UserType      string
	LastIndexed   time.Time
	LastFMSession *string

	GuildMembers *[]GuildMember `pg:"rel:has-many"`
}

/*
* Indexing Models
 */

// Artist represents a cached artist
type Artist struct {
	ID   int64 `pg:",pk"`
	Name string
}

// Album represents a cached artist
type Album struct {
	ID   int64 `pg:",pk"`
	Name string

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`
}

// Track represents a cached artist
type Track struct {
	ID   int64 `pg:",pk"`
	Name string

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`

	AlbumID *int64
	Album   *Album `pg:"rel:has-one"`
}

// Play represents a single play of a song (eg. a Last.fm scrobble)
type Play struct {
	ID          int64 `pg:",pk"`
	ScrobbledAt time.Time

	UserID int64
	User   *User `pg:"rel:has-one"`

	TrackID int64
	Track   *Track `pg:"rel:has-one"`
}

type RateYourMusicAlbum struct {
	ID              int64 `pg:",pk"`
	RateYourMusicID string
	ReleaseYear     *int

	Title            string
	ArtistName       string
	ArtistNativeName *string

	Albums []Album `pg:"many2many:rate_your_music_album_albums"`
}

type RateYourMusicAlbumAlbum struct {
	RateYourMusicAlbumID int64
	RateYourMusicAlbum   *RateYourMusicAlbum `pg:"rel:has-one"`

	AlbumID int64
	Album   *Album `pg:"rel:has-one"`
}

// Rating represents a single user's rating from rateyourmusic
type Rating struct {
	ID     int `pg:",pk"`
	Rating int

	UserID int64
	User   *User `pg:"rel:has-one"`

	RateYourMusicAlbumID int64
	RateYourMusicAlbum   *RateYourMusicAlbum `pg:"rel:has-one"`
}

/*
* Aggregated Structures
 */

// ArtistCount represents aggregated artist plays
type ArtistCount struct {
	ID        int64 `pg:",pk"`
	Playcount int32

	UserID int64
	User   *User `pg:"rel:has-one"`

	ArtistID int64
	Artist   *Artist `pg:"rel:has-one"`
}

// AlbumCount represents aggregated album plays
type AlbumCount struct {
	ID        int64 `pg:",pk"`
	Playcount int32

	UserID int64
	User   *User `pg:"rel:has-one"`

	AlbumID int64
	Album   *Album `pg:"rel:has-one"`
}

// TrackCount represents aggregated track plays
type TrackCount struct {
	ID        int64 `pg:",pk"`
	Playcount int32

	UserID int64
	User   *User `pg:"rel:has-one"`

	TrackID int64
	Track   *Track `pg:"rel:has-one"`
}

/*
* Guild sync structures
 */

// GuildMember represents a discord user in a guild
type GuildMember struct {
	GuildID string

	User   *User `pg:"rel:has-one"`
	UserID int64
}

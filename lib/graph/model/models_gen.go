// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Album struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Artist *Artist  `json:"artist"`
	Tracks []*Track `json:"tracks"`
}

type AlbumCount struct {
	Album     *Album `json:"album"`
	Playcount int    `json:"playcount"`
}

type AlbumInput struct {
	Artist *ArtistInput `json:"artist"`
	Name   *string      `json:"name"`
}

type AlbumPlaysSettings struct {
	PageInput *PageInput  `json:"pageInput"`
	Album     *AlbumInput `json:"album"`
	Sort      *string     `json:"sort"`
}

type AlbumTopTracksResponse struct {
	Album     *Album                 `json:"album"`
	TopTracks []*AmbiguousTrackCount `json:"topTracks"`
}

type AmbiguousTrack struct {
	Name   string   `json:"name"`
	Artist string   `json:"artist"`
	Albums []*Album `json:"albums"`
}

type AmbiguousTrackCount struct {
	Name      string `json:"name"`
	Playcount int    `json:"playcount"`
}

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ArtistCount struct {
	Artist    *Artist `json:"artist"`
	Playcount int     `json:"playcount"`
}

type ArtistInput struct {
	Name *string `json:"name"`
}

type ArtistPlaysSettings struct {
	PageInput *PageInput   `json:"pageInput"`
	Artist    *ArtistInput `json:"artist"`
	Sort      *string      `json:"sort"`
}

type ArtistSearchCriteria struct {
	Keywords *string `json:"keywords"`
}

type ArtistSearchResult struct {
	ArtistID        int    `json:"artistID"`
	ArtistName      string `json:"artistName"`
	ListenerCount   int    `json:"listenerCount"`
	GlobalPlaycount int    `json:"globalPlaycount"`
}

type ArtistSearchResults struct {
	Artists []*ArtistSearchResult `json:"artists"`
}

type ArtistTopAlbumsResponse struct {
	Artist    *Artist       `json:"artist"`
	TopAlbums []*AlbumCount `json:"topAlbums"`
}

type ArtistTopTracksResponse struct {
	Artist    *Artist                `json:"artist"`
	TopTracks []*AmbiguousTrackCount `json:"topTracks"`
}

type GuildMember struct {
	UserID  int    `json:"userID"`
	GuildID string `json:"guildID"`
	User    *User  `json:"user"`
}

type PageInput struct {
	Limit *int `json:"limit"`
}

type Play struct {
	ID          int    `json:"id"`
	ScrobbledAt int    `json:"scrobbledAt"`
	User        *User  `json:"user"`
	Track       *Track `json:"track"`
}

type RateYourMusicAlbum struct {
	RateYourMusicID  string  `json:"rateYourMusicID"`
	Title            string  `json:"title"`
	ReleaseYear      *int    `json:"releaseYear"`
	ArtistName       string  `json:"artistName"`
	ArtistNativeName *string `json:"artistNativeName"`
}

type RateYourMusicArtist struct {
	ArtistName       string  `json:"artistName"`
	ArtistNativeName *string `json:"artistNativeName"`
}

type Rating struct {
	RateYourMusicAlbum *RateYourMusicAlbum `json:"rateYourMusicAlbum"`
	Rating             int                 `json:"rating"`
}

type RatingsSettings struct {
	User      *UserInput  `json:"user"`
	Album     *AlbumInput `json:"album"`
	PageInput *PageInput  `json:"pageInput"`
}

type SearchSettings struct {
	Exact *bool      `json:"exact"`
	User  *UserInput `json:"user"`
}

type TaskStartResponse struct {
	TaskName string `json:"taskName"`
	Success  bool   `json:"success"`
	Token    string `json:"token"`
}

type Track struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Artist *Artist `json:"artist"`
	Album  *Album  `json:"album"`
}

type TrackCount struct {
	Track     *Track `json:"track"`
	Playcount int    `json:"playcount"`
}

type TrackInput struct {
	Artist *ArtistInput `json:"artist"`
	Album  *AlbumInput  `json:"album"`
	Name   *string      `json:"name"`
}

type TrackPlaysSettings struct {
	PageInput *PageInput  `json:"pageInput"`
	Track     *TrackInput `json:"track"`
	Sort      *string     `json:"sort"`
}

type TrackTopAlbumsResponse struct {
	Track     *AmbiguousTrack `json:"track"`
	TopAlbums []*TrackCount   `json:"topAlbums"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	DiscordID string    `json:"discordID"`
	UserType  *UserType `json:"userType"`
}

type UserInput struct {
	DiscordID      *string `json:"discordID"`
	LastFMUsername *string `json:"lastFMUsername"`
	WavyUsername   *string `json:"wavyUsername"`
}

type WhoKnowsAlbumResponse struct {
	Rows  []*WhoKnowsRow `json:"rows"`
	Album *Album         `json:"album"`
}

type WhoKnowsArtistResponse struct {
	Rows   []*WhoKnowsRow `json:"rows"`
	Artist *Artist        `json:"artist"`
}

type WhoKnowsRow struct {
	User      *User `json:"user"`
	Playcount int   `json:"playcount"`
}

type WhoKnowsSettings struct {
	GuildID *string `json:"guildID"`
	Limit   *int    `json:"limit"`
}

type WhoKnowsTrackResponse struct {
	Rows  []*WhoKnowsRow  `json:"rows"`
	Track *AmbiguousTrack `json:"track"`
}

type UserType string

const (
	UserTypeWavy   UserType = "Wavy"
	UserTypeLastfm UserType = "Lastfm"
)

var AllUserType = []UserType{
	UserTypeWavy,
	UserTypeLastfm,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeWavy, UserTypeLastfm:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

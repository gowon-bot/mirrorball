package lastfm

import (
	"net/url"
)

const ApiSigReplace = "ApiSigReplace"

type Requestable struct {
	Username string
	Session  *string
}

func (r Requestable) EncodeValues(key string, v *url.Values) error {
	v.Set("username", r.Username)

	if r.Session != nil {
		v.Set("sk", *r.Session)
		v.Set("api_sig", ApiSigReplace)
	}

	return nil
}

// Image is the struct type for a Last.fm image
type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

// RankAttributes is the struct type for
type RankAttributes struct {
	Rank string `json:"rank"`
}

// ErrorResponse is the struct type for a last.fm error response
type ErrorResponse struct {
	Error   int16  `json:"error"`
	Message string `json:"message"`
}

// UserInfoResponse is the struct type for a user.getInfo response from last.fm
type UserInfoResponse struct {
	User struct {
		Playlist   string  `json:"playlists"`
		Playcount  string  `json:"playcount"`
		Gender     string  `json:"gender"`
		Name       string  `json:"name"`
		Subscriber string  `json:"subscriber"`
		URL        string  `json:"url"`
		Country    string  `json:"country"`
		Image      []Image `json:"images"`
		Registered struct {
			Unixtime string `json:"unixtime"`
			AsNumber int64  `json:"#text"`
		} `json:"registered"`
		Type      string `json:"type"`
		Age       string `json:"age"`
		Bootstrap string `json:"bootstrap"`
		Realname  string `json:"realname"`
	} `json:"user"`
}

// UserInfoParams is the parameters for a user.getInfo call
type UserInfoParams struct {
	Username Requestable `url:"username"`
}

// RecentTrack is a struct containing a recent track from last.fm
type RecentTrack struct {
	Artist struct {
		MBID string `json:"mbid"`
		Text string `json:"#text"`
	} `json:"artist"`
	Attributes struct {
		IsNowPlaying string `json:"nowplaying"`
	} `json:"@attr"`
	MBID  string `json:"mbid"`
	Album struct {
		MBID string `json:"mbid"`
		Text string `json:"#text"`
	} `json:"album"`
	Images     []Image `json:"image"`
	Streamable string  `json:"streamable"`
	URL        string  `json:"url"`
	Name       string  `json:"name"`
	Timestamp  struct {
		UTS  string `json:"uts"`
		Text string `json:"#text"`
	} `json:"date"`
}

// RecentTracksResponse is the struct type for a user.getRecentTracks response from last.fm
type RecentTracksResponse struct {
	RecentTracks struct {
		Attributes struct {
			Page       string `json:"page"`
			Total      string `json:"total"`
			User       string `json:"user"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
		} `json:"@attr"`

		Tracks []RecentTrack `json:"track"`
	} `json:"recenttracks"`
}

// RecentTracksParams is the parameters for a user.recentTracks call
type RecentTracksParams struct {
	Username Requestable `url:"username"`
	Limit    int         `url:"limit"`
	Page     int         `url:"page"`
	Period   string      `url:"period"`
	From     string      `url:"from"`
}

// TopArtist is the struct type for a last.fm top artist
type TopArtist struct {
	Attributes RankAttributes `json:"@attr"`
	MBID       string         `json:"mbid"`
	URL        string         `json:"url"`
	Playcount  string         `json:"playcount"`
	Images     []Image        `json:"image"`
	Name       string         `json:"name"`
	Streamable string         `json:"streamable"`
}

// TopArtistsResponse is the struct type for a user.getTopArtists response from last.fm
type TopArtistsResponse struct {
	TopArtists struct {
		Artists    []TopArtist `json:"artist"`
		Attributes struct {
			Page       string `json:"page"`
			Total      string `json:"total"`
			User       string `json:"user"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
		} `json:"@attr"`
	} `json:"topartists"`
}

// TopEntityParams is the parameters for a user.getInfo response
type TopEntityParams struct {
	Username Requestable `url:"username"`
	Limit    int         `url:"limit"`
	Page     int         `url:"page"`
	Period   string      `url:"period"`
}

// TopAlbum is the struct type for a last.fm top album
type TopAlbum struct {
	Artist struct {
		URL  string `json:"url"`
		Name string `json:"name"`
		MBID string `json:"mbid"`
	} `json:"artist"`
	Attributes RankAttributes `json:"@attr"`
	Images     []Image        `json:"image"`
	Playcount  string         `json:"playcount"`
	URL        string         `json:"url"`
	Name       string         `json:"name"`
	MBID       string         `json:"mbid"`
}

// TopAlbumsResponse is the struct type for a user.getTopAlbums response from last.fm
type TopAlbumsResponse struct {
	TopAlbums struct {
		Albums     []TopAlbum `json:"album"`
		Attributes struct {
			Page       string `json:"page"`
			Total      string `json:"total"`
			User       string `json:"user"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
		} `json:"@attr"`
	} `json:"topalbums"`
}

// TopTrack is the struct type for a last.fm top track
type TopTrack struct {
	MBID      string  `json:"mbid"`
	Name      string  `json:"name"`
	URL       string  `json:"url"`
	Duration  string  `json:"duration"`
	Playcount string  `json:"playcount"`
	Images    []Image `json:"image"`

	Attributes struct {
		Rank string `json:"rank"`
	} `json:"@attr"`

	Artist struct {
		URL  string `json:"url"`
		Name string `json:"name"`
		Mbid string `json:"mbid"`
	} `json:"artist"`

	Streamable struct {
		Fulltrack string `json:"fulltrack"`
		Text      string `json:"#text"`
	} `json:"streamable"`
}

// TopTracksResponse is the struct type for a user.getTopTracks response from last.fm
type TopTracksResponse struct {
	TopTracks struct {
		Attributes struct {
			Page       string `json:"page"`
			Total      string `json:"total"`
			User       string `json:"user"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
		} `json:"@attr"`

		Tracks []TopTrack `json:"track"`
	} `json:"toptracks"`
}

package lastfm

// Image is the struct type for a Last.fm image
type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
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
	Username string `url:"username"`
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

		Tracks [](struct {
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
			date       struct {
				UTS  string `json:"uts"`
				Text string `json:"#text"`
			}
		}) `json:"track"`
	} `json:"recenttracks"`
}

// RecentTracksParams is the parameters for a user.recentTracks call
type RecentTracksParams struct {
	Username string `url:"username"`
	Limit    int    `url:"limit"`
	Page     int    `url:"page"`
	Period   string `url:"period"`
}

// // TopTracksResponse is the struct type for a user.getTopTracks response from last.fm
// type TopTracksResponse struct {
// 	TopTracks struct {
// 		Attributes struct {
// 			Page       string `json:"page"`
// 			Total      string `json:"total"`
// 			User       string `json:"user"`
// 			PerPage    string `json:"perPage"`
// 			TotalPages string `json:"totalPages"`
// 		} `json:"@attr"`

// 		track [](struct {
// 			MBID      string  `json:"mbid"`
// 			Name      string  `json:"name"`
// 			URL       string  `json:"url"`
// 			Duration  string  `json:"duration"`
// 			Playcount string  `json:"playcount"`
// 			Images    []Image `json:"image"`

// 			Attributes struct {
// 				Rank string `json:"rank"`
// 			} `json:"@attr"`

// 			Artist struct {
// 				URL  string `json:"url"`
// 				Name string `json:"name"`
// 				Mbid string `json:"mbid"`
// 			} `json:"artist"`

// 			Streamable struct {
// 				Fulltrack string `json:"fulltrack"`
// 				Text      string `json:"#text"`
// 			} `json:"streamable"`
// 		})
// 	} `json:"toptracks"`
// }

// // TopTracksParams is the parameters for a user.getInfo response
// type TopTracksParams struct {
// 	Username string `url:"username"`
// }

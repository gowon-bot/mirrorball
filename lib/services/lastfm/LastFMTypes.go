package lastfm

// ErrorResponse is the struct type for a last.fm error response
type ErrorResponse struct {
	Error   int16  `json:"error"`
	Message string `json:"message"`
}

// UserInfoResponse is the struct type for a user.getInfo response from last.fm
type UserInfoResponse struct {
	User struct {
		Playlist   string `json:"playlists"`
		Playcount  string `json:"playcount"`
		Gender     string `json:"gender"`
		Name       string `json:"name"`
		Subscriber string `json:"subscriber"`
		URL        string `json:"url"`
		Country    string `json:"country"`
		Image      [](struct {
			Size string `json:"size"`
			URL  string `json:"#text"`
		})
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

// UserInfoParams is the parameters for a user.getInfo response
type UserInfoParams struct {
	Username string `url:"username"`
}

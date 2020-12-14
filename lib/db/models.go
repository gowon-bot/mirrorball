package db

// User is the database model for a last.fm user
type User struct {
	ID             int64  `json:"id"`
	LastFMUsername string `json:"lastFMUsername"`
}

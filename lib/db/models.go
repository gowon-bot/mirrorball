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

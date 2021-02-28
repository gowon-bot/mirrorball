package db

import (
	"time"
)

// SetLastIndexed sets a user's last indexed time
func (u User) SetLastIndexed(to time.Time) {
	Db.Model(&u).Set("last_indexed = ?", to).WherePK().Update()
}

// IsWavyUser returns whether a user has a Wavy user type
func (u User) IsWavyUser() bool {
	return u.UserType == "Wavy"
}

package indexing

import "github.com/jivison/gowon-indexer/lib/db"

// Indexing holds methods for indexing users
type Indexing struct{}

// FullIndex indexes a user for the first time
func (i Indexing) FullIndex(user *db.User) {}

// CreateService creates an instance of the response service object
func CreateService() *Indexing {
	service := &Indexing{}

	return service
}

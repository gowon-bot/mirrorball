package analysis

import (
	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

// Analysis holds methods for generating API responses
type Analysis struct {
	indexingService *indexing.Indexing
	lastFMService   *lastfm.API
}

// CreateService creates an instance of the analysis service object
func CreateService() *Analysis {
	service := &Analysis{
		indexingService: indexing.CreateService(),
		lastFMService:   lastfm.CreateAPIService(),
	}

	return service
}

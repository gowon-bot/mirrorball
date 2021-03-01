package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// PresentArtist converts a database artist into a graphql artist
func PresentArtist(artist *db.Artist) *model.Artist {
	return &model.Artist{
		ID:   int(artist.ID),
		Name: artist.Name,
	}
}

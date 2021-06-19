package presenters

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
)

func PresentArtistSearchResults(artists []analysis.SearchArtistResult) *model.ArtistSearchResults {
	results := &model.ArtistSearchResults{
		Artists: []*model.ArtistSearchResult{},
	}

	for _, artist := range artists {
		results.Artists = append(results.Artists, &model.ArtistSearchResult{
			ArtistID:        int(artist.ID),
			ArtistName:      artist.Name,
			ListenerCount:   artist.ListenerCount,
			GlobalPlaycount: artist.GlobalPlaycount,
		})
	}

	return results
}

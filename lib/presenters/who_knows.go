package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// PresentWhoKnowsArtistResponse builds a graphql who knows artist response from an artist and db artist counts
func PresentWhoKnowsArtistResponse(artist *db.Artist, artistCounts []db.ArtistCount) *model.WhoKnowsArtistResponse {
	var whoKnowsRows []*model.WhoKnowsRow

	for _, artistCount := range artistCounts {
		whoKnowsRows = append(whoKnowsRows, &model.WhoKnowsRow{
			User:      PresentUser(artistCount.User),
			Playcount: int(artistCount.Playcount),
		})
	}

	return &model.WhoKnowsArtistResponse{
		Rows:   whoKnowsRows,
		Artist: PresentArtist(artist),
	}
}

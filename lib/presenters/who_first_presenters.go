package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func PresentWhoFirstArtistResponse(artist *db.Artist, plays []db.Play) *model.WhoFirstArtistResponse {
	var whoKnowsRows []*model.WhoFirstRow

	for _, play := range plays {
		whoKnowsRows = append(whoKnowsRows, &model.WhoFirstRow{
			User:        PresentUser(play.User),
			ScrobbledAt: int(play.ScrobbledAt.Unix()),
		})
	}

	return &model.WhoFirstArtistResponse{
		Rows:   whoKnowsRows,
		Artist: PresentArtist(artist),
	}
}

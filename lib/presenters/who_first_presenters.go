package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func PresentWhoFirstRow(play db.Play) *model.WhoFirstRow {
	return &model.WhoFirstRow{
		User:        PresentUser(play.User),
		ScrobbledAt: int(play.ScrobbledAt.Unix()),
	}
}

func PresentWhoFirstArtistResponse(artist *db.Artist, plays []db.Play, undated []db.Play) *model.WhoFirstArtistResponse {
	var whoKnowsRows []*model.WhoFirstRow
	var undatedRows []*model.WhoFirstRow

	for _, play := range plays {
		whoKnowsRows = append(whoKnowsRows, PresentWhoFirstRow(play))
	}

	for _, play := range undated {
		undatedRows = append(undatedRows, PresentWhoFirstRow(play))
	}

	return &model.WhoFirstArtistResponse{
		Rows:    whoKnowsRows,
		Undated: undatedRows,
		Artist:  PresentArtist(artist),
	}
}

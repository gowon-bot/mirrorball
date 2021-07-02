package analysis

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
)

// WhoFirstArtist returns a list of who first listened to an artist
func (a Analysis) WhoFirstArtist(artist *db.Artist, settings *model.WhoKnowsSettings, whoLast bool) ([]db.Play, error) {
	var whoFirst []db.Play

	var sort string
	aggFunc := "min"

	if whoLast {
		sort = "desc"
	} else {
		sort = "asc"
	}

	if whoLast {
		aggFunc = "max"
	}

	query := db.Db.Model(&whoFirst).
		Relation("Track._").
		ColumnExpr(aggFunc+"(scrobbled_at) as scrobbled_at").
		Column("user_id").
		Where("artist_id = ?", artist.ID).
		Group("user_id").OrderExpr("1 " + sort)

	err := inputparser.CreateParser(query).ParseWhoKnowsSettings(settings).GetQuery().Select()

	if err != nil {
		return whoFirst, customerrors.DatabaseUnknownError()
	}

	whoFirst, err = a.AddUserToPlays(whoFirst)

	if err != nil {
		return whoFirst, nil
	}

	return whoFirst, nil
}

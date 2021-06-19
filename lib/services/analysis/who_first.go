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

	subquery := db.Db.Model((*db.Play)(nil)).
		Relation("Track._").
		ColumnExpr(aggFunc+"(scrobbled_at) as scrobbled_at").
		Column("user_id").
		Where("artist_id = ?", artist.ID).
		Group("user_id")

	query := db.Db.Model(&whoFirst).
		Relation("User").
		Where("(scrobbled_at, play.user_id) IN (?)", subquery).
		Order("scrobbled_at "+sort, "username desc")

	err := inputparser.CreateParser(query).ParseWhoKnowsSettings(settings).GetQuery().Select()

	if err != nil {
		return whoFirst, customerrors.DatabaseUnknownError()
	}

	return whoFirst, nil
}

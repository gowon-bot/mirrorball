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

	if whoLast {
		sort = "desc"
	} else {
		sort = "asc"
	}

	subquery := db.Db.Model((*db.Play)(nil)).
		Relation("Track._").
		ColumnExpr("Row_number() over(PARTITION BY play.user_id ORDER BY scrobbled_at "+sort+", play.id) AS _rn").
		Column("play.user_id", "scrobbled_at").
		Where("artist_id = ?", artist.ID)

	subquery = inputparser.CreateParser(subquery).ParseWhoKnowsSettings(settings).GetQuery()

	query := db.Db.Model().TableExpr("(?) as play", subquery).
		Column("play.*").
		ColumnExpr("u.id as user__id, u.discord_id as user__discord_id, u.username as user__username").
		Join("JOIN users u ON u.id = play.user_id").
		Where("play._rn = 1").
		Order("username desc")

	err := query.Select(&whoFirst)

	if err != nil {
		return whoFirst, customerrors.DatabaseUnknownError()
	}

	return whoFirst, nil
}

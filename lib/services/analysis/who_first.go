package analysis

import (
	"github.com/go-pg/pg/v10"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
)

// WhoFirstArtist returns a list of who first listened to an artist
func (a Analysis) WhoFirstArtist(artist *db.Artist, settings *model.WhoKnowsSettings, whoLast bool, excludeIDs []int64) ([]db.Play, error) {
	var whoFirst []db.Play

	sort := "asc"
	aggFunc := "min"

	if whoLast {
		sort = "desc"
	}

	if whoLast {
		aggFunc = "max"
	}

	query := db.Db.Model(&whoFirst).
		Relation("Track._").
		ColumnExpr(aggFunc+"(scrobbled_at) as scrobbled_at").
		Column("play.user_id").
		Where("artist_id = ?", artist.ID).
		Where("scrobbled_at > '1970-01-01 00:00:00'::date").
		Where("play.user_id NOT IN (?)", pg.In(excludeIDs)).
		Group("play.user_id").OrderExpr("1 " + sort)

	err := inputparser.CreateParser(query).ParseWhoKnowsSettings(settings, &inputparser.InputSettings{
		UserIDPath: `play"."user_id`,
	}).GetQuery().Select()

	if err != nil {
		return whoFirst, customerrors.DatabaseUnknownError()
	}

	whoFirst, err = a.AddUserToPlays(whoFirst)

	if err != nil {
		return whoFirst, nil
	}

	return whoFirst, nil
}

func (a Analysis) WhoHasUndatedArtist(artist *db.Artist, settings *model.WhoKnowsSettings, whoLast bool) ([]db.Play, error) {
	var whoFirst []db.Play

	sort := "asc"

	if whoLast {
		sort = "desc"
	}

	query := db.Db.Model(&whoFirst).
		Relation("Track._").
		Column("play.user_id").
		Where("artist_id = ?", artist.ID).
		Where("scrobbled_at < '1970-01-01 00:00:00'::date").
		Group("play.user_id").OrderExpr("1 " + sort)

	err := inputparser.CreateParser(query).ParseWhoKnowsSettings(settings, &inputparser.InputSettings{
		UserIDPath: `play"."user_id`,
	}).GetQuery().Select()

	if err != nil {
		return whoFirst, customerrors.DatabaseUnknownError()
	}

	whoFirst, err = a.AddUserToPlays(whoFirst)

	if err != nil {
		return whoFirst, nil
	}

	return whoFirst, nil
}

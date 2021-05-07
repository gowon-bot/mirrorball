package indexing

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func (i Indexing) GetPlays(user db.User, pageInput *model.PageInput) ([]db.Play, error) {
	var plays []db.Play

	query := db.Db.Model(&plays).Relation("User").Relation("Track").Relation("Track.Artist").Relation("Track.Album").Where("user_id = ?", user.ID).Order("scrobbled_at desc")
	query = ParsePageInput(query, pageInput)

	err := query.Select()

	if err != nil {
		return plays, customerrors.DatabaseUnknownError()
	}

	return plays, nil
}

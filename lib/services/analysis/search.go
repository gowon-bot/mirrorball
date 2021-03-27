package analysis

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

type SearchArtistResult struct {
	ID              int64
	Name            string
	ListenerCount   int
	GlobalPlaycount int
}

func (a Analysis) SearchArtist(criteria model.ArtistSearchCriteria, settings *model.SearchSettings) ([]SearchArtistResult, error) {
	var artists []SearchArtistResult

	var query = db.Db.Model((*db.Artist)(nil)).
		Join("JOIN artist_counts ac ON ac.artist_id = artist.id").
		Column("name").
		ColumnExpr("artist.id").
		ColumnExpr("COALESCE(count(playcount), 0) as listener_count").
		ColumnExpr("COALESCE(sum(playcount), 0) as global_playcount").
		Group("name", "artist.id").
		OrderExpr("COALESCE(count(playcount), 0) desc")

	query = ParseArtistSearchCriteria(query, criteria, settings)

	err := query.Select(&artists)

	if err != nil {
		return artists, customerrors.DatabaseUnknownError()
	}

	return artists, nil
}

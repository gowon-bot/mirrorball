package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func PresentArtistRank(artist *db.Artist, whoKnows []db.ArtistCount, plays *db.ArtistCount, rank int) *model.ArtistRankResponse {

	builtResponse := &model.ArtistRankResponse{
		Artist: PresentArtist(artist),

		Listeners: len(whoKnows),
	}

	if rank != -1 {
		builtResponse.Playcount = int(plays.Playcount)
		builtResponse.Rank = rank + 1
	} else {
		builtResponse.Playcount = 0
		builtResponse.Rank = -1
		builtResponse.Above = PresentArtistCount(&whoKnows[len(whoKnows)-1])
	}

	if rank != -1 && rank-1 >= 0 {
		builtResponse.Above = PresentArtistCount(&whoKnows[rank-1])
	}

	if rank != -1 && rank+1 < len(whoKnows) {
		builtResponse.Below = PresentArtistCount(&whoKnows[rank+1])
	}

	return builtResponse
}

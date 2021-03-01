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

// PresentWhoKnowsAlbumResponse builds a graphql who knows album response from an album and db album counts
func PresentWhoKnowsAlbumResponse(album *db.Album, albumCounts []db.AlbumCount) *model.WhoKnowsAlbumResponse {
	var whoKnowsRows []*model.WhoKnowsRow

	for _, albumCount := range albumCounts {
		whoKnowsRows = append(whoKnowsRows, &model.WhoKnowsRow{
			User:      PresentUser(albumCount.User),
			Playcount: int(albumCount.Playcount),
		})
	}

	return &model.WhoKnowsAlbumResponse{
		Rows:  whoKnowsRows,
		Album: PresentAlbum(album),
	}
}

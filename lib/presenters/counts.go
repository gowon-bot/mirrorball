package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func PresentArtistTopAlbums(dbArtist *db.Artist, albumCounts []db.AlbumCount) *model.ArtistTopAlbumsResponse {
	artist := PresentArtist(dbArtist)

	topAlbumsResponse := &model.ArtistTopAlbumsResponse{
		Artist:    artist,
		TopAlbums: []*model.AlbumCount{},
	}

	for _, albumCount := range albumCounts {
		topAlbumsResponse.TopAlbums = append(topAlbumsResponse.TopAlbums, PresentAlbumCount(&albumCount))
	}

	return topAlbumsResponse
}

func PresentAlbumCount(albumCount *db.AlbumCount) *model.AlbumCount {
	presentedCount := &model.AlbumCount{
		Playcount: int(albumCount.Playcount),
	}

	if albumCount.Album != nil {
		presentedCount.Album = PresentAlbum(albumCount.Album)
	}

	return presentedCount
}

package rateyourmusic

import (
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

type RYMAlbumGenerator struct {
	albumsToConvert []indexing.AlbumToConvert
	indexingService *indexing.Indexing
}

func CreateAlbumGenerator() *RYMAlbumGenerator {
	return &RYMAlbumGenerator{
		albumsToConvert: []indexing.AlbumToConvert{},
		indexingService: indexing.CreateService(),
	}
}

func (ag *RYMAlbumGenerator) AddAlbumToEnsureExists(artistName, albumName string) {
	ag.albumsToConvert = append(ag.albumsToConvert, indexing.AlbumToConvert{
		ArtistName: artistName, AlbumName: albumName,
	})
}

func (ag *RYMAlbumGenerator) EnsureAlbumsExist() {
	ag.indexingService.ConvertAlbums(ag.albumsToConvert, nil)
}

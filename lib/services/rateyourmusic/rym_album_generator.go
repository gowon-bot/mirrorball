package rateyourmusic

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/meta"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

type RYMAlbumGenerator struct {
	albumsToConvert []indexing.AlbumToConvert
	allCombinations []indexing.AlbumToConvert
	indexingService *indexing.Indexing
	conversionMap   AlbumCombinationConversionMap
}

func CreateAlbumGenerator() *RYMAlbumGenerator {
	return &RYMAlbumGenerator{
		albumsToConvert: []indexing.AlbumToConvert{},
		allCombinations: []indexing.AlbumToConvert{},
		indexingService: indexing.CreateService(),
		conversionMap:   CreateAlbumCombinationConversionMap(),
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

func (ag *RYMAlbumGenerator) AddCombinations(combinations []indexing.AlbumToConvert, row RawRateYourMusicRating) {
	ag.allCombinations = append(ag.allCombinations, combinations...)

	for _, combination := range combinations {
		ag.conversionMap.Append(combination.ArtistName, combination.AlbumName, &row)
	}
}

func (ag *RYMAlbumGenerator) SelectAllCombinations() ([]db.Album, error) {
	searchableAlbums := ag.indexingService.GenerateAlbumsToSearch(ag.allCombinations)

	return dbhelpers.SelectAlbumsWhereInMany(searchableAlbums, constants.ChunkSize)
}

func (ag *RYMAlbumGenerator) AttachAlbumCombinations(ratings []*RawRateYourMusicRating, albums []db.Album) []*RawRateYourMusicRating {

	for _, album := range albums {
		if rows, _, ok := ag.conversionMap.Get(album.Artist.Name, album.Name); ok {
			var rowIDs []string

			for _, conversionRow := range rows {
				rowIDs = append(rowIDs, conversionRow.RYMID)
			}

			for _, row := range ratings {
				for _, rowID := range rowIDs {
					if rowID == row.RYMID {
						row.AllAlbums = append(row.AllAlbums, indexing.AlbumToConvert{ArtistName: album.Artist.Name, AlbumName: album.Name})
					}
				}
			}
		}
	}

	return ratings
}

// To avoid circular imports, this has to be here...
type AlbumCombinationConversionMap struct{ *meta.ConversionMap }

func (lm AlbumCombinationConversionMap) Get(artistName, albumName string) ([]*RawRateYourMusicRating, string, bool) {
	if artist, ok := lm.PrivateGet(artistName); ok {
		artistMap := artist.Value.(meta.ConversionMap)

		albums, ok := artistMap.PrivateGet(albumName)

		if ok {
			return albums.Value.([]*RawRateYourMusicRating), albums.Key, ok
		}
	}

	return []*RawRateYourMusicRating{}, albumName, false
}

func (lm AlbumCombinationConversionMap) Append(artistName, albumName string, album *RawRateYourMusicRating) {
	if _, ok := lm.PrivateGet(artistName); !ok {
		lm.PrivateSet(artistName, *meta.CreateConversionMap())
	}

	artist, _ := lm.PrivateGet(artistName)

	artistMap := artist.Value.(meta.ConversionMap)

	var newAlbums []*RawRateYourMusicRating

	existingValue, _, ok := lm.Get(artistName, albumName)

	if ok {
		newAlbums = append(newAlbums, existingValue...)
	}

	newAlbums = append(newAlbums, album)

	artistMap.PrivateSet(albumName, newAlbums)
}

func CreateAlbumCombinationConversionMap() AlbumCombinationConversionMap {
	return AlbumCombinationConversionMap{ConversionMap: meta.CreateConversionMap()}
}

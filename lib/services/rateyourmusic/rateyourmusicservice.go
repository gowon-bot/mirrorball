package rateyourmusic

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

// RateYourMusic holds methods for interacting with the cached rateyourmusic data
type RateYourMusic struct {
	indexingService *indexing.Indexing
}

type RateYourMusicAlbumMap = map[string]db.RateYourMusicAlbum

func (rym RateYourMusic) ConvertRateYourMusicAlbums(rawAlbums []RawRateYourMusicRating) (RateYourMusicAlbumMap, error) {
	albumsMap := make(RateYourMusicAlbumMap)

	var rymsIDs []interface{}

	for _, album := range rawAlbums {
		rymsIDs = append(rymsIDs, album.RYMID)
	}

	rateYourMusicAlbums, err := dbhelpers.SelectRateYourMusicAlbumsWhereInMany(rymsIDs, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, album := range rateYourMusicAlbums {
		albumsMap[album.RateYourMusicID] = album
	}

	albumsToCreate := rym.generateRateYourMusicAlbumsToCreate(albumsMap, rawAlbums)

	createdAlbums, err := rym.createRateYourMusicAlbums(albumsToCreate)

	if err != nil {
		return nil, err
	}

	for _, album := range createdAlbums {
		albumsMap[album.RateYourMusicID] = album
	}

	err = rym.createRateYourMusicAlbumAlbums(albumsToCreate, albumsMap)

	return albumsMap, nil
}

func (rym RateYourMusic) SaveRatings(ratings []RawRateYourMusicRating, rymsAlbumsMap RateYourMusicAlbumMap, user db.User) ([]db.Rating, error) {
	var dbRatings []db.Rating

	for _, rating := range ratings {
		rymsAlbum := rymsAlbumsMap[rating.RYMID]

		dbRating := db.Rating{
			Rating:               rating.Rating,
			User:                 &user,
			UserID:               user.ID,
			RateYourMusicAlbum:   &rymsAlbum,
			RateYourMusicAlbumID: rymsAlbum.ID,
		}

		dbRatings = append(dbRatings, dbRating)
	}

	_, err := dbhelpers.UpdateOrCreateManyRatings(dbRatings, user.ID)

	if err != nil {
		return nil, err
	}

	return dbRatings, nil
}

func (rym RateYourMusic) createRateYourMusicAlbums(albumsToCreate []RawRateYourMusicRating) ([]db.RateYourMusicAlbum, error) {
	var dbAlbums []db.RateYourMusicAlbum

	for _, album := range albumsToCreate {
		dbAlbums = append(dbAlbums, db.RateYourMusicAlbum{RateYourMusicID: album.RYMID, ReleaseYear: &album.ReleaseYear})
	}

	albums, err := dbhelpers.InsertManyRateYourMusicAlbums(dbAlbums, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (rym RateYourMusic) generateRateYourMusicAlbumsToCreate(albumsMap RateYourMusicAlbumMap, rawAlbums []RawRateYourMusicRating) []RawRateYourMusicRating {
	var albumsToCreate []RawRateYourMusicRating

	for _, album := range rawAlbums {
		if _, ok := albumsMap[album.RYMID]; !ok {
			albumsToCreate = append(albumsToCreate, album)
		}
	}

	return albumsToCreate
}

func (rym RateYourMusic) convertAlbumsFromRatings(rawAlbums []RawRateYourMusicRating) (indexing.AlbumsMap, error) {
	var albumList []indexing.AlbumToConvert

	for _, album := range rawAlbums {
		albumList = append(albumList, rym.generateAlbumCombinations(album)...)
	}

	albumsMap, err := rym.indexingService.ConvertAlbums(albumList)

	return albumsMap, err
}

func (rym RateYourMusic) generateAlbumCombinations(rawAlbum RawRateYourMusicRating) []indexing.AlbumToConvert {
	var combinations []indexing.AlbumToConvert

	combinations = append(combinations, indexing.AlbumToConvert{
		ArtistName: rawAlbum.ArtistName,
		AlbumName:  rawAlbum.Title,
	})

	if rawAlbum.ArtistNativeName != nil {
		combinations = append(combinations, indexing.AlbumToConvert{
			ArtistName: *rawAlbum.ArtistNativeName,
			AlbumName:  rawAlbum.Title,
		})
	}

	return combinations
}

func (rym RateYourMusic) createRateYourMusicAlbumAlbums(albums []RawRateYourMusicRating, rymsAlbumsMap RateYourMusicAlbumMap) error {
	albumsMap, err := rym.convertAlbumsFromRatings(albums)

	if err != nil {
		return err
	}

	var albumAlbumsToCreate []db.RateYourMusicAlbumAlbum

	for _, album := range albums {
		for _, combination := range rym.generateAlbumCombinations(album) {

			dbAlbum := albumsMap[combination.ArtistName][combination.AlbumName]

			albumAlbumsToCreate = append(albumAlbumsToCreate, db.RateYourMusicAlbumAlbum{
				RateYourMusicAlbumID: rymsAlbumsMap[album.RYMID].ID,
				AlbumID:              dbAlbum.ID,
			})
		}
	}

	_, err = dbhelpers.InsertManyRateYourMusicAlbumAlbums(albumAlbumsToCreate, constants.ChunkSize)

	if err != nil {
		return err
	}

	return nil
}

func (rym RateYourMusic) ResetRatings(user db.User) {
	db.Db.Model((*db.Rating)(nil)).Where("user_id=?", user.ID).Delete()
}

// CreateService creates an instance of the indexing service object
func CreateService() *RateYourMusic {
	service := &RateYourMusic{
		indexingService: indexing.CreateService(),
	}

	return service
}

package rateyourmusic

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/meta"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

func (rym RateYourMusic) ConvertRateYourMusicAlbums(rawAlbums []RawRateYourMusicRating) (meta.RateYourMusicAlbumConversionMap, error) {
	albumsMap := meta.CreateRateYourMusicAlbumConversionMap()

	var rymsIDs []interface{}

	for _, album := range rawAlbums {
		rymsIDs = append(rymsIDs, album.RYMID)
	}

	rateYourMusicAlbums, err := dbhelpers.SelectRateYourMusicAlbumsWhereInMany(rymsIDs, constants.ChunkSize)

	if err != nil {
		return albumsMap, err
	}

	for _, album := range rateYourMusicAlbums {
		albumsMap.Set(album.RateYourMusicID, album)
	}

	albumsToCreate, albumsToUpdate := rym.generateRateYourMusicAlbumsToCreate(albumsMap, rawAlbums)

	go rym.updateRateYourMusicAlbums(albumsToUpdate, albumsMap)

	createdAlbums, err := rym.createRateYourMusicAlbums(albumsToCreate)

	if err != nil {
		return albumsMap, err
	}

	for _, album := range createdAlbums {
		albumsMap.Set(album.RateYourMusicID, album)
	}

	err = rym.createRateYourMusicAlbumAlbums(albumsToCreate, albumsMap)

	if err != nil {
		return albumsMap, err
	}

	return albumsMap, nil
}

func (rym RateYourMusic) SaveRatings(ratings []RawRateYourMusicRating, rymsAlbumsMap meta.RateYourMusicAlbumConversionMap, user db.User) ([]db.Rating, error) {
	var dbRatings []db.Rating

	for _, rating := range ratings {
		rymsAlbum, _, _ := rymsAlbumsMap.Get(rating.RYMID)

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
		dbAlbums = append(dbAlbums, db.RateYourMusicAlbum{
			RateYourMusicID:  album.RYMID,
			ReleaseYear:      &album.ReleaseYear,
			Title:            album.Title,
			ArtistName:       album.ArtistName,
			ArtistNativeName: album.ArtistNativeName,
		})
	}

	albums, err := dbhelpers.InsertManyRateYourMusicAlbums(dbAlbums, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (rym RateYourMusic) updateRateYourMusicAlbums(albumsToUpdate []RateYourMusicAlbumToUpdate, albumsMap meta.RateYourMusicAlbumConversionMap) ([]db.RateYourMusicAlbum, error) {
	var dbAlbums []db.RateYourMusicAlbum
	var rawAlbums []RawRateYourMusicRating

	for _, album := range albumsToUpdate {
		dbAlbums = append(dbAlbums, album.dbAlbum)
		rawAlbums = append(rawAlbums, album.rawAlbum)
	}

	albums, err := dbhelpers.UpdateManyRateYourMusicAlbums(dbAlbums, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	rym.updateRateYourMusicAlbumAlbums(rawAlbums, albumsMap)

	return albums, nil
}

type RateYourMusicAlbumToUpdate struct {
	dbAlbum  db.RateYourMusicAlbum
	rawAlbum RawRateYourMusicRating
}

func (rym RateYourMusic) generateRateYourMusicAlbumsToCreate(albumsMap meta.RateYourMusicAlbumConversionMap, rawAlbums []RawRateYourMusicRating) ([]RawRateYourMusicRating, []RateYourMusicAlbumToUpdate) {
	var albumsToCreate []RawRateYourMusicRating
	var albumsToUpdate []RateYourMusicAlbumToUpdate

	for _, album := range rawAlbums {
		if dbAlbum, _, ok := albumsMap.Get(album.RYMID); !ok {
			albumsToCreate = append(albumsToCreate, album)
		} else {
			albumsToUpdate = append(albumsToUpdate, RateYourMusicAlbumToUpdate{
				dbAlbum:  dbAlbum,
				rawAlbum: album,
			})
		}
	}

	return albumsToCreate, albumsToUpdate
}

func (rym RateYourMusic) convertAlbumsFromRatings(rawAlbums []RawRateYourMusicRating) (meta.AlbumConversionMap, error) {
	var albumList []indexing.AlbumToConvert

	for _, album := range rawAlbums {
		albumList = append(albumList, album.AllAlbums...)
	}

	albumsMap, err := rym.indexingService.ConvertAlbums(albumList, nil)

	return albumsMap, err
}

func (rym RateYourMusic) createRateYourMusicAlbumAlbums(albums []RawRateYourMusicRating, rymsAlbumsMap meta.RateYourMusicAlbumConversionMap) error {
	albumsMap, err := rym.convertAlbumsFromRatings(albums)

	if err != nil {
		return err
	}

	var albumAlbumsToCreate []db.RateYourMusicAlbumAlbum

	for _, album := range albums {
		for _, combination := range album.AllAlbums {

			dbAlbum, _, _ := albumsMap.Get(combination.ArtistName, combination.AlbumName)
			rateYourMusicAlbum, _, _ := rymsAlbumsMap.Get(album.RYMID)

			albumAlbumsToCreate = append(albumAlbumsToCreate, db.RateYourMusicAlbumAlbum{
				RateYourMusicAlbumID: rateYourMusicAlbum.ID,
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

func (rym RateYourMusic) updateRateYourMusicAlbumAlbums(albums []RawRateYourMusicRating, rymsAlbumsMap meta.RateYourMusicAlbumConversionMap) error {
	albumsMap, err := rym.convertAlbumsFromRatings(albums)

	if err != nil {
		return err
	}

	var albumAlbumsToCreate []db.RateYourMusicAlbumAlbum

	for _, album := range albums {
		var unfilteredAlbums []db.RateYourMusicAlbumAlbum

		for _, combination := range album.AllAlbums {
			dbAlbum, _, _ := albumsMap.Get(combination.ArtistName, combination.AlbumName)
			rateYourMusicAlbum, _, _ := rymsAlbumsMap.Get(album.RYMID)

			albumAlbumsToCreate = append(unfilteredAlbums, db.RateYourMusicAlbumAlbum{
				RateYourMusicAlbumID: rateYourMusicAlbum.ID,
				AlbumID:              dbAlbum.ID,
			})
		}

		albumAlbumsToCreate = append(albumAlbumsToCreate, rym.filterDuplicateAlbumAlbums(album.RYMID, unfilteredAlbums)...)
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

func (rym RateYourMusic) filterDuplicateAlbumAlbums(rateYourMusicAlbumID string, albumAlbums []db.RateYourMusicAlbumAlbum) []db.RateYourMusicAlbumAlbum {
	var dbAlbumAlbums []db.RateYourMusicAlbumAlbum
	var filtered []db.RateYourMusicAlbumAlbum

	err := db.Db.Model(&dbAlbumAlbums).Where("rate_your_music_album_id = ?", rateYourMusicAlbumID).Select()

	if err != nil {
		return albumAlbums
	}

	for _, album := range albumAlbums {
		if ok := checkForDuplicateAlbumAlbum(dbAlbumAlbums, album); ok {
			filtered = append(filtered, album)
		}
	}

	return filtered
}

func checkForDuplicateAlbumAlbum(in []db.RateYourMusicAlbumAlbum, check db.RateYourMusicAlbumAlbum) bool {
	for _, albumAlbum := range in {
		if albumAlbum.AlbumID == check.AlbumID {
			return false
		}
	}

	return true
}

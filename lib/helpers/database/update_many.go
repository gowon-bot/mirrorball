package dbhelpers

import (
	"math"

	"github.com/jinzhu/copier"
	"github.com/jivison/gowon-indexer/lib/db"
)

func UpdateManyArtistCounts(artistCounts []db.ArtistCount, itemsPerChunk float64) ([]db.ArtistCount, error) {
	if len(artistCounts) == 0 {
		return nil, nil
	}

	var chunks [][]interface{}
	var allArtistCounts []db.ArtistCount

	chunks = make([][]interface{}, int(math.Floor(float64(len(artistCounts))/(itemsPerChunk)))+1)

	for index, artistCount := range artistCounts {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		copiedArtistCount := db.ArtistCount{}
		copier.Copy(&copiedArtistCount, &artistCount)

		chunks[chunkIndex] = append(chunks[chunkIndex], &copiedArtistCount)
	}

	for _, chunk := range chunks {
		var updatedArtistCounts []db.ArtistCount

		_, err := db.Db.Model(chunk...).WherePK().Update(&updatedArtistCounts)

		if err != nil {
			return allArtistCounts, err
		}

		allArtistCounts = append(allArtistCounts, updatedArtistCounts...)
	}

	return allArtistCounts, nil
}

func UpdateManyAlbumCounts(albumCounts []db.AlbumCount, itemsPerChunk float64) ([]db.AlbumCount, error) {
	if len(albumCounts) == 0 {
		return nil, nil
	}

	var chunks [][]interface{}
	var allAlbumCounts []db.AlbumCount

	chunks = make([][]interface{}, int(math.Floor(float64(len(albumCounts))/(itemsPerChunk)))+1)

	for index, albumCount := range albumCounts {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		copiedAlbumCount := db.AlbumCount{}
		copier.Copy(&copiedAlbumCount, &albumCount)

		chunks[chunkIndex] = append(chunks[chunkIndex], &copiedAlbumCount)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(chunk...).WherePK().Update()

		if err != nil {
			return allAlbumCounts, err
		}

	}

	return allAlbumCounts, nil
}

func UpdateManyTrackCounts(trackCounts []db.TrackCount, itemsPerChunk float64) ([]db.TrackCount, error) {
	if len(trackCounts) == 0 {
		return nil, nil
	}

	var chunks [][]interface{}
	var allTrackCounts []db.TrackCount

	chunks = make([][]interface{}, int(math.Floor(float64(len(trackCounts))/(itemsPerChunk)))+1)

	for index, trackCount := range trackCounts {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		copiedTrackCount := db.TrackCount{}
		copier.Copy(&copiedTrackCount, &trackCount)

		chunks[chunkIndex] = append(chunks[chunkIndex], &copiedTrackCount)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(chunk...).WherePK().Update()

		if err != nil {
			return allTrackCounts, err
		}

	}

	return allTrackCounts, nil
}

func UpdateManyRatings(ratings []db.Rating, itemsPerChunk float64) ([]db.Rating, error) {
	if len(ratings) == 0 {
		return nil, nil
	}

	var chunks [][]interface{}
	var allRatings []db.Rating

	chunks = make([][]interface{}, int(math.Floor(float64(len(ratings))/(itemsPerChunk)))+1)

	for index, rating := range ratings {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		copiedRating := db.Rating{}
		copier.Copy(&copiedRating, &rating)

		chunks[chunkIndex] = append(chunks[chunkIndex], &copiedRating)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(chunk...).WherePK().Update()

		if err != nil {
			return allRatings, err
		}

	}

	return allRatings, nil
}

func UpdateManyRateYourMusicAlbums(albums []db.RateYourMusicAlbum, itemsPerChunk float64) ([]db.RateYourMusicAlbum, error) {
	if len(albums) == 0 {
		return nil, nil
	}

	var chunks [][]interface{}
	var allRatings []db.RateYourMusicAlbum

	chunks = make([][]interface{}, int(math.Floor(float64(len(albums))/(itemsPerChunk)))+1)

	for index, album := range albums {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		copiedAlbum := db.RateYourMusicAlbum{}
		copier.Copy(&copiedAlbum, &album)

		chunks[chunkIndex] = append(chunks[chunkIndex], &copiedAlbum)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(chunk...).WherePK().Update()

		if err != nil {
			return allRatings, err
		}

	}

	return allRatings, nil
}

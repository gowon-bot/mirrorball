package dbhelpers

import (
	"math"

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

		chunks[chunkIndex] = append(chunks[chunkIndex], &artistCount)
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

		chunks[chunkIndex] = append(chunks[chunkIndex], &albumCount)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(chunk...).WherePK().Update()

		if err != nil {
			return allAlbumCounts, err
		}

		allAlbumCounts = append(allAlbumCounts)
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

		chunks[chunkIndex] = append(chunks[chunkIndex], &trackCount)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(chunk...).WherePK().Update()

		if err != nil {
			return allTrackCounts, err
		}

		allTrackCounts = append(allTrackCounts)
	}

	return allTrackCounts, nil
}

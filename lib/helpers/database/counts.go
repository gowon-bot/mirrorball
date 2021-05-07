package dbhelpers

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
)

type ArtistCountsMap = map[int64]db.ArtistCount
type AlbumCountsMap = map[int64]db.AlbumCount
type TrackCountsMap = map[int64]db.TrackCount

func UpdateOrCreateManyArtistCounts(artistCounts []db.ArtistCount, userID int64) (ArtistCountsMap, error) {
	var artistIDs []interface{}

	for _, artist := range artistCounts {
		artistIDs = append(artistIDs, artist.ArtistID)
	}

	artistCountsMap := make(ArtistCountsMap)
	savedArtistCounts, err := SelectArtistCountsWhereInMany(artistIDs, userID, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, artistCount := range savedArtistCounts {
		artistCountsMap[artistCount.ArtistID] = artistCount
	}

	var artistCountsToCreate []db.ArtistCount
	var artistCountsToUpdate []db.ArtistCount

	for _, artistCount := range artistCounts {
		if savedArtistCount, ok := artistCountsMap[artistCount.ArtistID]; !ok {
			artistCountsToCreate = append(artistCountsToCreate, artistCount)
		} else {
			newPlayCount := artistCount.Playcount + savedArtistCount.Playcount

			updatedArtistCount := savedArtistCount
			updatedArtistCount.Playcount = newPlayCount

			artistCountsToUpdate = append(artistCountsToUpdate, updatedArtistCount)
		}
	}

	createdArtistCounts, err := InsertManyArtistCounts(artistCountsToCreate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	updatedArtistCounts, err := UpdateManyArtistCounts(artistCountsToUpdate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, artistCount := range append(createdArtistCounts, updatedArtistCounts...) {
		artistCountsMap[artistCount.ArtistID] = artistCount
	}

	return artistCountsMap, nil
}

func UpdateOrCreateManyAlbumCounts(albumCounts []db.AlbumCount, userID int64) (AlbumCountsMap, error) {
	var albumIDs []interface{}

	for _, album := range albumCounts {
		albumIDs = append(albumIDs, album.AlbumID)
	}

	albumCountsMap := make(AlbumCountsMap)
	savedAlbumCounts, err := SelectAlbumCountsWhereInMany(albumIDs, userID, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, albumCount := range savedAlbumCounts {
		albumCountsMap[albumCount.AlbumID] = albumCount
	}

	var albumCountsToCreate []db.AlbumCount
	var albumCountsToUpdate []db.AlbumCount

	for _, albumCount := range albumCounts {
		if savedAlbumCount, ok := albumCountsMap[albumCount.AlbumID]; !ok {
			albumCountsToCreate = append(albumCountsToCreate, albumCount)
		} else {
			newPlayCount := albumCount.Playcount + savedAlbumCount.Playcount

			updatedAlbumCount := savedAlbumCount
			updatedAlbumCount.Playcount = newPlayCount

			albumCountsToUpdate = append(albumCountsToUpdate, updatedAlbumCount)
		}
	}

	createdAlbumCounts, err := InsertManyAlbumCounts(albumCountsToCreate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	updatedAlbumCounts, err := UpdateManyAlbumCounts(albumCountsToUpdate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, albumCount := range append(createdAlbumCounts, updatedAlbumCounts...) {
		albumCountsMap[albumCount.AlbumID] = albumCount
	}

	return albumCountsMap, nil
}

func UpdateOrCreateManyTrackCounts(trackCounts []db.TrackCount, userID int64) (TrackCountsMap, error) {
	var trackIDs []interface{}

	for _, track := range trackCounts {
		trackIDs = append(trackIDs, track.TrackID)
	}

	trackCountsMap := make(TrackCountsMap)
	savedTrackCounts, err := SelectTrackCountsWhereInMany(trackIDs, userID, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, trackCount := range savedTrackCounts {
		trackCountsMap[trackCount.TrackID] = trackCount
	}

	var trackCountsToCreate []db.TrackCount
	var trackCountsToUpdate []db.TrackCount

	for _, trackCount := range trackCounts {
		if savedTrackCount, ok := trackCountsMap[trackCount.TrackID]; !ok {
			trackCountsToCreate = append(trackCountsToCreate, trackCount)
		} else {
			newPlayCount := trackCount.Playcount + savedTrackCount.Playcount

			updatedTrackCount := savedTrackCount
			updatedTrackCount.Playcount = newPlayCount

			trackCountsToUpdate = append(trackCountsToUpdate, updatedTrackCount)
		}
	}

	createdTrackCounts, err := InsertManyTrackCounts(trackCountsToCreate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	updatedTrackCounts, err := UpdateManyTrackCounts(trackCountsToUpdate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, trackCount := range append(createdTrackCounts, updatedTrackCounts...) {
		trackCountsMap[trackCount.TrackID] = trackCount
	}

	return trackCountsMap, nil
}

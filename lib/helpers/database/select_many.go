package dbhelpers

import (
	"math"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
)

func SelectArtistsWhereInMany(artists []string, itemsPerChunk float64) ([]db.Artist, error) {
	var chunks [][]interface{}
	var allArtists []db.Artist

	if len(artists) == 0 {
		return allArtists, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(artists))/itemsPerChunk))+1)

	for index, artist := range artists {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], artist)
	}

	for _, chunk := range chunks {
		var selectedArtists []db.Artist

		err := db.Db.Model((*db.Artist)(nil)).
			Where(
				"artist.name IN (?)",
				pg.In(
					chunk,
				),
			).Select(&selectedArtists)

		if err != nil {
			return allArtists, customerrors.DatabaseUnknownError()
		}

		allArtists = append(allArtists, selectedArtists...)
	}

	return allArtists, nil
}

func SelectTagsWhereInMany(tags []string, itemsPerChunk float64) ([]db.Tag, error) {
	var chunks [][]interface{}
	var allTags []db.Tag

	if len(tags) == 0 {
		return allTags, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(tags))/itemsPerChunk))+1)

	for index, tag := range tags {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], strings.ToLower(tag))
	}

	for _, chunk := range chunks {
		var selectedTags []db.Tag

		err := db.Db.Model((*db.Tag)(nil)).
			Where(
				"lowercase(tag.name) IN (?)",
				pg.In(
					chunk,
				),
			).Select(&selectedTags)

		if err != nil {
			return allTags, customerrors.DatabaseUnknownError()
		}

		allTags = append(allTags, selectedTags...)
	}

	return allTags, nil
}

func SelectAlbumsWhereInMany(albums []interface{}, itemsPerChunk float64) ([]db.Album, error) {
	var chunks [][]interface{}
	var allAlbums []db.Album

	if len(albums) == 0 {
		return allAlbums, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(albums))/itemsPerChunk))+1)

	for index, album := range albums {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], album)
	}

	for _, chunk := range chunks {
		var selectedAlbums []db.Album

		err := db.Db.Model((*db.Album)(nil)).Relation("Artist").
			Where(
				"(artist.name, album.name) IN (?)",
				pg.InMulti(
					chunk...,
				),
			).Select(&selectedAlbums)

		if err != nil {
			return allAlbums, customerrors.DatabaseUnknownError()
		}

		allAlbums = append(allAlbums, selectedAlbums...)
	}

	return allAlbums, nil
}

func SelectTracksWhereInMany(tracks []interface{}, itemsPerChunk float64) ([]db.Track, error) {
	var chunks [][]interface{}
	var allTracks []db.Track

	if len(tracks) == 0 {
		return allTracks, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(tracks))/itemsPerChunk))+1)

	for index, track := range tracks {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], track)
	}

	for _, chunk := range chunks {
		var selectedTracks []db.Track

		err := db.Db.Model((*db.Track)(nil)).
			Relation("Artist").
			Relation("Album").
			Where(
				"(artist.name, track.name, album.name) IN (?)",
				pg.InMulti(
					chunk...,
				),
			).Select(&selectedTracks)

		if err != nil {
			return allTracks, customerrors.DatabaseUnknownError()
		}

		allTracks = append(allTracks, selectedTracks...)
	}

	return allTracks, nil
}

func SelectArtistCountsWhereInMany(artistIDs []interface{}, userID int64, itemsPerChunk float64) ([]db.ArtistCount, error) {
	var chunks [][]interface{}
	var allArtistCounts []db.ArtistCount

	if len(artistIDs) == 0 {
		return allArtistCounts, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(artistIDs))/itemsPerChunk))+1)

	for index, artistID := range artistIDs {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], artistID)
	}

	for _, chunk := range chunks {
		var selectedArtistCounts []db.ArtistCount

		err := db.Db.Model((*db.ArtistCount)(nil)).
			Relation("Artist").
			Where(
				"artist_id IN (?)",
				pg.In(chunk),
			).
			Where("user_id = ?", userID).
			Select(&selectedArtistCounts)

		if err != nil {
			return allArtistCounts, customerrors.DatabaseUnknownError()
		}

		allArtistCounts = append(allArtistCounts, selectedArtistCounts...)
	}

	return allArtistCounts, nil
}

func SelectAlbumCountsWhereInMany(albumIDs []interface{}, userID int64, itemsPerChunk float64) ([]db.AlbumCount, error) {
	var chunks [][]interface{}
	var allAlbumCounts []db.AlbumCount

	if len(albumIDs) == 0 {
		return allAlbumCounts, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(albumIDs))/itemsPerChunk))+1)

	for index, albumID := range albumIDs {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], albumID)
	}

	for _, chunk := range chunks {
		var selectedAlbumCounts []db.AlbumCount

		err := db.Db.Model((*db.AlbumCount)(nil)).
			Relation("Album").
			Relation("Album.Artist").
			Where(
				"album_id IN (?)",
				pg.In(chunk),
			).
			Where("user_id = ?", userID).
			Select(&selectedAlbumCounts)

		if err != nil {
			return allAlbumCounts, customerrors.DatabaseUnknownError()
		}

		allAlbumCounts = append(allAlbumCounts, selectedAlbumCounts...)
	}

	return allAlbumCounts, nil
}

func SelectTrackCountsWhereInMany(trackIDs []interface{}, userID int64, itemsPerChunk float64) ([]db.TrackCount, error) {
	var chunks [][]interface{}
	var allTrackCounts []db.TrackCount

	if len(trackIDs) == 0 {
		return allTrackCounts, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(trackIDs))/itemsPerChunk))+1)

	for index, trackID := range trackIDs {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], trackID)
	}

	for _, chunk := range chunks {
		var selectedTrackCounts []db.TrackCount

		err := db.Db.Model((*db.TrackCount)(nil)).
			Relation("Track").
			Relation("Track.Artist").
			Relation("Track.Album").
			Where(
				"track_id IN (?)",
				pg.In(chunk),
			).
			Where("user_id = ?", userID).
			Select(&selectedTrackCounts)

		if err != nil {
			return allTrackCounts, customerrors.DatabaseUnknownError()
		}

		allTrackCounts = append(allTrackCounts, selectedTrackCounts...)
	}

	return allTrackCounts, nil
}

func SelectRateYourMusicAlbumsWhereInMany(rymsAlbumIDs []interface{}, itemsPerChunk float64) ([]db.RateYourMusicAlbum, error) {
	var chunks [][]interface{}
	var allAlbums []db.RateYourMusicAlbum

	if len(rymsAlbumIDs) == 0 {
		return allAlbums, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(rymsAlbumIDs))/itemsPerChunk))+1)

	for index, rymsID := range rymsAlbumIDs {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], rymsID)
	}

	for _, chunk := range chunks {
		var selectedAlbums []db.RateYourMusicAlbum

		err := db.Db.Model((*db.RateYourMusicAlbum)(nil)).
			// Relation("Albums").
			Where(
				"rate_your_music_id IN (?)",
				pg.In(chunk),
			).
			Select(&selectedAlbums)

		if err != nil {
			return allAlbums, customerrors.DatabaseUnknownError()
		}

		allAlbums = append(allAlbums, selectedAlbums...)
	}

	return allAlbums, nil
}

func SelectRatingsWhereInMany(rymsAlbumIDs []interface{}, userID int64, itemsPerChunk float64) ([]db.Rating, error) {
	var chunks [][]interface{}
	var allRatings []db.Rating

	if len(rymsAlbumIDs) == 0 {
		return allRatings, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(rymsAlbumIDs))/itemsPerChunk))+1)

	for index, albumID := range rymsAlbumIDs {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], albumID)
	}

	for _, chunk := range chunks {
		var selectedRatings []db.Rating

		err := db.Db.Model((*db.Rating)(nil)).
			Relation("RateYourMusicAlbum").
			// Relation("RateYourMusicAlbum.Albums")
			Where(
				"rate_your_music_album_id IN (?)",
				pg.In(chunk),
			).
			Where("user_id = ?", userID).
			Select(&selectedRatings)

		if err != nil {
			return allRatings, customerrors.DatabaseUnknownError()
		}

		allRatings = append(allRatings, selectedRatings...)
	}

	return allRatings, nil
}

func SelectUsersWhereInMany(userIDs []int64, itemsPerChunk float64) ([]db.User, error) {
	var chunks [][]interface{}
	var allUsers []db.User

	if len(userIDs) == 0 {
		return allUsers, nil
	}

	chunks = make([][]interface{}, int(math.Floor(float64(len(userIDs))/itemsPerChunk))+1)

	for index, user := range userIDs {
		chunkIndex := int(math.Floor(float64(index+1) / itemsPerChunk))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]interface{}, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], user)
	}

	for _, chunk := range chunks {
		var selectedUsers []db.User

		err := db.Db.Model((*db.User)(nil)).
			Where(
				"id IN (?)",
				pg.In(
					chunk,
				),
			).Select(&selectedUsers)

		if err != nil {
			return allUsers, customerrors.DatabaseUnknownError()
		}

		allUsers = append(allUsers, selectedUsers...)
	}

	return allUsers, nil
}

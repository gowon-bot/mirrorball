package dbhelpers

import (
	"math"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
)

// InsertManyArtists takes an input of a slice of artists, and inserts them in chunks
// so as to not hit the postgres stack limit
func InsertManyArtists(artists []db.Artist, itemsPerChunk float64) ([]db.Artist, error) {
	var chunks [][]db.Artist
	var allArtists []db.Artist

	chunks = make([][]db.Artist, int(math.Floor(float64(len(artists))/(itemsPerChunk)))+1)

	for index, artist := range artists {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		chunks[chunkIndex] = append(chunks[chunkIndex], artist)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(&chunk).Insert()

		if err != nil {
			return allArtists, customerrors.DatabaseUnknownError()
		}

		allArtists = append(allArtists, chunk...)
	}

	return allArtists, nil
}

// InsertManyAlbums takes an input of a slice of albums, and inserts them in chunks
// so as to not hit the postgres stack limit
func InsertManyAlbums(albums []db.Album, itemsPerChunk float64) ([]db.Album, error) {
	var chunks [][]db.Album
	var allAlbums []db.Album

	chunks = make([][]db.Album, int(math.Floor(float64(len(albums))/(itemsPerChunk)))+1)

	for index, album := range albums {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]db.Album, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], album)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(&chunk).Insert()

		if err != nil {
			return allAlbums, customerrors.DatabaseUnknownError()
		}

		allAlbums = append(allAlbums, chunk...)
	}

	return allAlbums, nil
}

// InsertManyTracks takes an input of a slice of tracks, and inserts them in chunks
// so as to not hit the postgres stack limit
func InsertManyTracks(tracks []db.Track, itemsPerChunk float64) ([]db.Track, error) {
	var chunks [][]db.Track
	var allTracks []db.Track

	chunks = make([][]db.Track, int(math.Floor(float64(len(tracks))/(itemsPerChunk)))+1)

	for index, track := range tracks {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]db.Track, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], track)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(&chunk).Insert()

		if err != nil {
			return allTracks, err
		}

		allTracks = append(allTracks, chunk...)
	}

	return allTracks, nil
}

func InsertManyPlays(plays []db.Play, itemsPerChunk float64) ([]db.Play, error) {
	var chunks [][]db.Play
	var allPlays []db.Play

	chunks = make([][]db.Play, int(math.Floor(float64(len(plays))/(itemsPerChunk)))+1)

	for index, play := range plays {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]db.Play, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], play)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(&chunk).Insert()

		if err != nil {
			return allPlays, err
		}

		allPlays = append(allPlays, chunk...)
	}

	return allPlays, nil
}

// InsertManyArtistCounts takes an input of a slice of artistCounts, and inserts them in chunks
// so as to not hit the postgres stack limit
func InsertManyArtistCounts(artistCounts []db.ArtistCount, itemsPerChunk float64) ([]db.ArtistCount, error) {
	if len(artistCounts) < 1 {
		return nil, nil
	}

	var chunks [][]db.ArtistCount
	var allArtistCounts []db.ArtistCount

	chunks = make([][]db.ArtistCount, int(math.Floor(float64(len(artistCounts))/(itemsPerChunk)))+1)

	for index, artistCount := range artistCounts {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]db.ArtistCount, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], artistCount)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(&chunk).Insert()

		if err != nil {
			return allArtistCounts, err
		}

		allArtistCounts = append(allArtistCounts, chunk...)
	}

	return allArtistCounts, nil
}

func InsertManyAlbumCounts(albumCounts []db.AlbumCount, itemsPerChunk float64) ([]db.AlbumCount, error) {
	if len(albumCounts) < 1 {
		return nil, nil
	}

	var chunks [][]db.AlbumCount
	var allAlbumCounts []db.AlbumCount

	chunks = make([][]db.AlbumCount, int(math.Floor(float64(len(albumCounts))/(itemsPerChunk)))+1)

	for index, albumCount := range albumCounts {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]db.AlbumCount, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], albumCount)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(&chunk).Insert()

		if err != nil {
			return allAlbumCounts, err
		}

		allAlbumCounts = append(allAlbumCounts, chunk...)
	}

	return allAlbumCounts, nil
}

func InsertManyTrackCounts(trackCounts []db.TrackCount, itemsPerChunk float64) ([]db.TrackCount, error) {
	if len(trackCounts) < 1 {
		return nil, nil
	}

	var chunks [][]db.TrackCount
	var allTrackCounts []db.TrackCount

	chunks = make([][]db.TrackCount, int(math.Floor(float64(len(trackCounts))/(itemsPerChunk)))+1)

	for index, trackCount := range trackCounts {
		chunkIndex := int(math.Floor(float64(index+1) / (itemsPerChunk)))

		if chunks[chunkIndex] == nil {
			chunks[chunkIndex] = make([]db.TrackCount, 0)
		}

		chunks[chunkIndex] = append(chunks[chunkIndex], trackCount)
	}

	for _, chunk := range chunks {
		_, err := db.Db.Model(&chunk).Insert()

		if err != nil {
			return allTrackCounts, err
		}

		allTrackCounts = append(allTrackCounts, chunk...)
	}

	return allTrackCounts, nil
}

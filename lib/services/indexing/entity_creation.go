package indexing

import (
	"strings"

	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/database"
)

func (i Indexing) CreateArtists(artistNames []string) ([]db.Artist, error) {
	var values []db.Artist

	if len(artistNames) < 1 {
		return values, nil
	}

	for _, artistName := range artistNames {
		values = append(values, db.Artist{Name: artistName})
	}

	values, err := helpers.InsertManyArtists(values, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	return values, nil
}

func (i Indexing) CreateAlbums(albums []db.Album) ([]db.Album, error) {
	if len(albums) < 1 {
		return albums, nil
	}

	albums, err := helpers.InsertManyAlbums(albums, constants.ChunkSize)

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return albums, nil
}

func (i Indexing) CreateTracks(tracks []db.Track) ([]db.Track, error) {
	if len(tracks) < 1 {
		return tracks, nil
	}

	tracks, err := helpers.InsertManyTracks(tracks, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (i Indexing) generateAlbumsToCreate(albumNames []AlbumToConvert, albumsMap AlbumsMap, existingArtistsMap *ArtistsMap) ([]db.Album, error) {
	var albumsToCreate []db.Album

	var artistNames []string

	for _, album := range albumNames {
		artistNames = append(artistNames, album.ArtistName)
	}

	var artistsMap ArtistsMap

	if existingArtistsMap == nil {
		newArtistsMap, err := i.ConvertArtists(artistNames)
		artistsMap = newArtistsMap

		if err != nil {
			return albumsToCreate, err
		}
	} else {
		artistsMap = *existingArtistsMap
	}

	for _, album := range albumNames {
		if _, ok := albumsMap[strings.ToLower(album.ArtistName)]; !ok {
			albumsMap[strings.ToLower(album.ArtistName)] = make(map[string]db.Album)
		}

		if _, ok := albumsMap[strings.ToLower(album.ArtistName)][strings.ToLower(album.AlbumName)]; !ok {
			artist := artistsMap[strings.ToLower(album.ArtistName)]

			albumsToCreate = append(albumsToCreate, db.Album{
				ArtistID: artist.ID,
				Name:     album.AlbumName,
				Artist:   &artist,
			})
		}
	}

	return albumsToCreate, nil
}

func (i Indexing) generateTracksToCreate(trackNames []TrackToConvert, tracksMap TracksMap, existingArtistsMap *ArtistsMap, existingAlbumsMap *AlbumsMap) ([]db.Track, error) {
	var tracksToCreate []db.Track

	var artistNames []string
	var albumNames []AlbumToConvert

	for _, track := range trackNames {
		artistNames = append(artistNames, track.ArtistName)

		if track.AlbumName != nil {
			albumNames = append(albumNames, AlbumToConvert{
				ArtistName: track.ArtistName,
				AlbumName:  *track.AlbumName,
			})
		}
	}

	var artistsMap ArtistsMap
	var albumsMap AlbumsMap

	if existingArtistsMap == nil {
		newArtistsMap, err := i.ConvertArtists(artistNames)
		artistsMap = newArtistsMap

		if err != nil {
			return tracksToCreate, err
		}
	} else {
		artistsMap = *existingArtistsMap
	}

	if existingAlbumsMap == nil {
		newAlbumsMap, err := i.ConvertAlbums(albumNames, &artistsMap)
		albumsMap = newAlbumsMap

		if err != nil {
			return tracksToCreate, err
		}
	} else {
		albumsMap = *existingAlbumsMap
	}

	for _, track := range trackNames {
		albumName := ""

		if track.AlbumName != nil {
			albumName = *track.AlbumName
		}

		if _, ok := tracksMap[strings.ToLower(track.ArtistName)]; !ok {
			tracksMap[strings.ToLower(track.ArtistName)] = make(map[string]map[string]db.Track)
		}

		if _, ok := tracksMap[strings.ToLower(track.ArtistName)][strings.ToLower(albumName)]; !ok {
			tracksMap[strings.ToLower(track.ArtistName)][strings.ToLower(albumName)] = make(map[string]db.Track)
		}

		if _, ok := tracksMap[strings.ToLower(track.ArtistName)][strings.ToLower(albumName)][strings.ToLower(track.TrackName)]; !ok {
			artist := artistsMap[strings.ToLower(track.ArtistName)]

			trackToCreate := db.Track{
				Name:     track.TrackName,
				Artist:   &artist,
				ArtistID: artist.ID,
			}

			if track.AlbumName != nil {
				album := albumsMap[strings.ToLower(track.ArtistName)][strings.ToLower(*track.AlbumName)]

				trackToCreate.Album = &album
				trackToCreate.AlbumID = &album.ID
			}

			tracksToCreate = append(tracksToCreate, trackToCreate)
		}
	}

	return tracksToCreate, nil
}

func (i Indexing) CreateTags(tags []db.Tag) ([]db.Tag, error) {
	if len(tags) < 1 {
		return tags, nil
	}

	values, err := helpers.InsertManyTags(tags, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	return values, nil
}

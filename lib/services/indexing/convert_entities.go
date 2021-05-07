package indexing

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/database"
)

func (i Indexing) ConvertArtists(artistNames []string) (map[string]db.Artist, error) {
	artistsMap := map[string]db.Artist{}
	artistsToCreate := []string{}

	artists, err := helpers.SelectArtistsWhereInMany(artistNames, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, artist := range artists {
		artistsMap[artist.Name] = artist
	}

	for _, artistName := range artistNames {
		if _, ok := artistsMap[artistName]; !ok {
			artistsToCreate = append(artistsToCreate, artistName)
		}
	}

	createdArtists, err := i.CreateArtists(artistsToCreate)

	if err != nil {
		return nil, err
	}

	for _, createdArtist := range createdArtists {
		artistsMap[createdArtist.Name] = createdArtist
	}

	return artistsMap, nil
}

func (i Indexing) ConvertAlbums(albumNames []AlbumToConvert) (AlbumsMap, error) {
	var albums []db.Album
	albumsMap := AlbumsMap{}

	albumsToSearch := i.generateAlbumsToSearch(albumNames)

	albums, err := helpers.SelectAlbumsWhereInMany(albumsToSearch, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, album := range albums {
		if _, ok := albumsMap[album.Artist.Name]; !ok {
			albumsMap[album.Artist.Name] = make(map[string]db.Album)
		}

		albumsMap[album.Artist.Name][album.Name] = album
	}

	albumsToCreate, err := i.generateAlbumsToCreate(albumNames, albumsMap)

	if err != nil {
		return nil, err
	}

	createdAlbums, err := i.CreateAlbums(albumsToCreate)

	if err != nil {
		return nil, err
	}

	for _, createdAlbum := range createdAlbums {
		if _, ok := albumsMap[createdAlbum.Artist.Name]; !ok {
			albumsMap[createdAlbum.Name] = make(map[string]db.Album)
		}

		albumsMap[createdAlbum.Artist.Name][createdAlbum.Name] = createdAlbum
	}

	return albumsMap, nil
}

func (i Indexing) ConvertTracks(trackNames []TrackToConvert) (TracksMap, error) {
	var tracks []db.Track
	tracksMap := TracksMap{}

	tracksToSearch := i.generateTracksToSearch(trackNames)

	tracks, err := helpers.SelectTracksWhereInMany(tracksToSearch, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, track := range tracks {
		albumName := ""

		if track.Album != nil {
			albumName = track.Album.Name
		}

		if _, ok := tracksMap[track.Artist.Name]; !ok {
			tracksMap[track.Artist.Name] = make(map[string]map[string]db.Track)
		}

		if _, ok := tracksMap[track.Artist.Name][albumName]; !ok {
			tracksMap[track.Artist.Name][albumName] = make(map[string]db.Track)
		}

		tracksMap[track.Artist.Name][albumName][track.Name] = track
	}

	tracksToCreate, err := i.generateTracksToCreate(trackNames, tracksMap)

	if err != nil {
		return nil, err
	}

	createdTracks, err := i.CreateTracks(tracksToCreate)

	if err != nil {
		return nil, err
	}

	for _, createdTrack := range createdTracks {
		albumName := ""

		if createdTrack.Album != nil {
			albumName = createdTrack.Album.Name
		}

		if _, ok := tracksMap[createdTrack.Artist.Name]; !ok {
			tracksMap[createdTrack.Artist.Name] = make(map[string]map[string]db.Track)
		}

		if _, ok := tracksMap[createdTrack.Artist.Name][albumName]; !ok {
			tracksMap[createdTrack.Artist.Name][albumName] = make(map[string]db.Track)
		}

		tracksMap[createdTrack.Artist.Name][albumName][createdTrack.Name] = createdTrack
	}

	return tracksMap, nil
}

package indexing

import (
	"strings"

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
		artistsMap[strings.ToLower(artist.Name)] = artist
	}

	for _, artistName := range artistNames {
		if _, ok := artistsMap[strings.ToLower(artistName)]; !ok {
			artistsToCreate = append(artistsToCreate, artistName)
		}
	}

	createdArtists, err := i.CreateArtists(artistsToCreate)

	if err != nil {
		return nil, err
	}

	for _, createdArtist := range createdArtists {
		artistsMap[strings.ToLower(createdArtist.Name)] = createdArtist
	}

	return artistsMap, nil
}

func (i Indexing) ConvertAlbums(albumNames []AlbumToConvert, existingArtistsMap *ArtistsMap) (AlbumsMap, error) {
	var albums []db.Album
	albumsMap := AlbumsMap{}

	albumsToSearch := i.GenerateAlbumsToSearch(albumNames)

	albums, err := helpers.SelectAlbumsWhereInMany(albumsToSearch, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, album := range albums {
		if _, ok := albumsMap[strings.ToLower(album.Artist.Name)]; !ok {
			albumsMap[strings.ToLower(album.Artist.Name)] = make(map[string]db.Album)
		}

		albumsMap[strings.ToLower(album.Artist.Name)][strings.ToLower(album.Name)] = album
	}

	albumsToCreate, err := i.generateAlbumsToCreate(albumNames, albumsMap, existingArtistsMap)

	if err != nil {
		return nil, err
	}

	createdAlbums, err := i.CreateAlbums(albumsToCreate)

	if err != nil {
		return nil, err
	}

	for _, createdAlbum := range createdAlbums {
		if _, ok := albumsMap[strings.ToLower(createdAlbum.Artist.Name)]; !ok {
			albumsMap[strings.ToLower(createdAlbum.Artist.Name)] = make(map[string]db.Album)
		}

		albumsMap[strings.ToLower(createdAlbum.Artist.Name)][strings.ToLower(createdAlbum.Name)] = createdAlbum
	}

	return albumsMap, nil
}

func (i Indexing) ConvertTracks(trackNames []TrackToConvert, existingArtistsMap *ArtistsMap, existingAlbumsMap *AlbumsMap) (TracksMap, error) {
	var tracks []db.Track
	tracksMap := TracksMap{}

	tracksToSearch := i.GenerateTracksToSearch(trackNames)

	tracks, err := helpers.SelectTracksWhereInMany(tracksToSearch, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, track := range tracks {
		albumName := ""

		if track.Album != nil {
			albumName = track.Album.Name
		}

		if _, ok := tracksMap[strings.ToLower(track.Artist.Name)]; !ok {
			tracksMap[strings.ToLower(track.Artist.Name)] = make(map[string]map[string]db.Track)
		}

		if _, ok := tracksMap[strings.ToLower(track.Artist.Name)][strings.ToLower(albumName)]; !ok {
			tracksMap[strings.ToLower(track.Artist.Name)][strings.ToLower(albumName)] = make(map[string]db.Track)
		}

		tracksMap[strings.ToLower(track.Artist.Name)][strings.ToLower(albumName)][strings.ToLower(track.Name)] = track
	}

	tracksToCreate, err := i.generateTracksToCreate(trackNames, tracksMap, existingArtistsMap, existingAlbumsMap)

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

		if _, ok := tracksMap[strings.ToLower(createdTrack.Artist.Name)]; !ok {
			tracksMap[strings.ToLower(createdTrack.Artist.Name)] = make(map[string]map[string]db.Track)
		}

		if _, ok := tracksMap[strings.ToLower(createdTrack.Artist.Name)][strings.ToLower(albumName)]; !ok {
			tracksMap[strings.ToLower(createdTrack.Artist.Name)][strings.ToLower(albumName)] = make(map[string]db.Track)
		}

		tracksMap[strings.ToLower(createdTrack.Artist.Name)][strings.ToLower(albumName)][strings.ToLower(createdTrack.Name)] = createdTrack
	}

	return tracksMap, nil
}

type TagsMap = map[string]db.Tag

func (i Indexing) ConvertTags(tagNames []string) (TagsMap, error) {
	tagsMap := make(TagsMap)
	tagsToCreate := []db.Tag{}

	tags, err := helpers.SelectTagsWhereInMany(tagNames, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		tagsMap[strings.ToLower(tag.Name)] = tag
	}

	for _, tagName := range tagNames {
		if _, ok := tagsMap[strings.ToLower(tagName)]; !ok {
			tagsToCreate = append(tagsToCreate, db.Tag{Name: tagName})
		}
	}

	createdTags, err := i.CreateTags(tagsToCreate)

	if err != nil {
		return nil, err
	}

	for _, createdTag := range createdTags {
		tagsMap[strings.ToLower(createdTag.Name)] = createdTag
	}

	return tagsMap, nil
}

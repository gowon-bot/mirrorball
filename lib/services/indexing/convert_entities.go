package indexing

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/meta"
)

func (i Indexing) ConvertArtists(artistNames []string) (meta.ArtistConversionMap, error) {
	artistsMap := meta.CreateArtistConversionMap()
	artistsToCreate := []string{}

	artists, err := helpers.SelectArtistsWhereInMany(artistNames, constants.ChunkSize)

	if err != nil {
		return artistsMap, err
	}

	for _, artist := range artists {
		artistsMap.Set(artist.Name, artist)
	}

	for _, artistName := range artistNames {
		if _, _, ok := artistsMap.Get(artistName); !ok {
			artistsToCreate = append(artistsToCreate, artistName)
		}
	}

	createdArtists, err := i.CreateArtists(artistsToCreate)

	if err != nil {
		return artistsMap, err
	}

	for _, createdArtist := range createdArtists {
		artistsMap.Set(createdArtist.Name, createdArtist)
	}

	return artistsMap, nil
}

func (i Indexing) ConvertAlbums(albumNames []AlbumToConvert, existingArtistsMap *meta.ArtistConversionMap) (meta.AlbumConversionMap, error) {
	var albums []db.Album
	albumsMap := meta.CreateAlbumConversionMap()

	albumsToSearch := i.GenerateAlbumsToSearch(albumNames)

	albums, err := helpers.SelectAlbumsWhereInMany(albumsToSearch, constants.ChunkSize)

	if err != nil {
		return albumsMap, err
	}

	for _, album := range albums {
		albumsMap.Set(album.Artist.Name, album.Name, album)
	}

	albumsToCreate, err := i.generateAlbumsToCreate(albumNames, albumsMap, existingArtistsMap)

	if err != nil {
		return albumsMap, err
	}

	createdAlbums, err := i.CreateAlbums(albumsToCreate)

	if err != nil {
		return albumsMap, err
	}

	for _, createdAlbum := range createdAlbums {
		albumsMap.Set(createdAlbum.Artist.Name, createdAlbum.Name, createdAlbum)
	}

	return albumsMap, nil
}

func (i Indexing) ConvertTracks(trackNames []TrackToConvert, existingArtistsMap *meta.ArtistConversionMap, existingAlbumsMap *meta.AlbumConversionMap) (meta.TrackConversionMap, error) {
	var tracks []db.Track
	tracksMap := meta.CreateTrackConversionMap()

	tracksToSearch := i.GenerateTracksToSearch(trackNames)

	tracks, err := helpers.SelectTracksWhereInMany(tracksToSearch, constants.ChunkSize)

	if err != nil {
		return tracksMap, err
	}

	for _, track := range tracks {
		albumName := ""

		if track.Album != nil {
			albumName = track.Album.Name
		}

		tracksMap.Set(track.Artist.Name, albumName, track.Name, track)
	}

	tracksToCreate, err := i.generateTracksToCreate(trackNames, tracksMap, existingArtistsMap, existingAlbumsMap)

	if err != nil {
		return tracksMap, err
	}

	createdTracks, err := i.CreateTracks(tracksToCreate)

	if err != nil {
		return tracksMap, err
	}

	for _, createdTrack := range createdTracks {
		albumName := ""

		if createdTrack.Album != nil {
			albumName = createdTrack.Album.Name
		}

		tracksMap.Set(createdTrack.Artist.Name, albumName, createdTrack.Name, createdTrack)
	}

	return tracksMap, nil
}

func (i Indexing) ConvertTags(tagNames []string) (meta.TagConversionMap, error) {
	tagsMap := meta.CreateTagConversionMap()
	tagsToCreate := []db.Tag{}

	tags, err := helpers.SelectTagsWhereInMany(tagNames, constants.ChunkSize)

	if err != nil {
		return tagsMap, err
	}

	for _, tag := range tags {
		tagsMap.Set(tag.Name, tag)
	}

	for _, tagName := range tagNames {
		if _, _, ok := tagsMap.Get(tagName); !ok {
			tagsToCreate = append(tagsToCreate, db.Tag{Name: tagName})
		}
	}

	createdTags, err := i.CreateTags(tagsToCreate)

	if err != nil {
		return tagsMap, err
	}

	for _, createdTag := range createdTags {
		tagsMap.Set(createdTag.Name, createdTag)
	}

	return tagsMap, nil
}

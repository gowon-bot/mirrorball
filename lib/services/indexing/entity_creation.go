package indexing

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	"github.com/jivison/gowon-indexer/lib/meta"
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

func (i Indexing) generateAlbumsToCreate(albumNames []AlbumToConvert, albumsMap meta.AlbumConversionMap, existingArtistsMap *meta.ArtistConversionMap) ([]db.Album, error) {
	var albumsToCreate []db.Album

	var artistNames []string

	for _, album := range albumNames {
		artistNames = append(artistNames, album.ArtistName)
	}

	var artistsMap meta.ArtistConversionMap

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
		if _, _, ok := albumsMap.Get(album.ArtistName, album.AlbumName); !ok {
			artist, _, _ := artistsMap.Get(album.ArtistName)

			albumsToCreate = append(albumsToCreate, db.Album{
				ArtistID: artist.ID,
				Name:     album.AlbumName,
				Artist:   &artist,
			})
		}
	}

	return albumsToCreate, nil
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

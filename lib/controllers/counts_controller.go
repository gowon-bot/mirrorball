package controllers

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

// Returns a list of a users top scrobbled albums under a user
func ArtistTopAlbums(userInput model.UserInput, artistInput model.ArtistInput) (*model.ArtistTopAlbumsResponse, error) {
	usersService := users.CreateService()
	indexingService := indexing.CreateService()
	analysisService := analysis.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	artist, err := indexingService.GetArtist(artistInput, false)

	if err != nil {
		return nil, err
	}

	topAlbums, err := analysisService.ArtistTopAlbums(user.ID, artist.ID)

	return presenters.PresentArtistTopAlbums(artist, topAlbums), nil
}

func ArtistTopTracks(userInput model.UserInput, artistInput model.ArtistInput) (*model.ArtistTopTracksResponse, error) {
	usersService := users.CreateService()
	indexingService := indexing.CreateService()
	analysisService := analysis.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	artist, err := indexingService.GetArtist(artistInput, false)

	if err != nil {
		return nil, err
	}

	topTracks, err := analysisService.ArtistTopTracks(user.ID, artist.ID)

	return presenters.PresentArtistTopTracks(artist, topTracks), nil
}

func AlbumTopTracks(userInput model.UserInput, albumInput model.AlbumInput) (*model.AlbumTopTracksResponse, error) {
	usersService := users.CreateService()
	indexingService := indexing.CreateService()
	analysisService := analysis.CreateService()

	user := usersService.FindUserByInput(userInput)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	album, err := indexingService.GetAlbum(albumInput, false)

	if err != nil {
		return nil, err
	}

	topTracks, err := analysisService.AlbumTopTracks(user.ID, album.ID)

	return presenters.PresentAlbumTopTracks(album, topTracks), nil
}

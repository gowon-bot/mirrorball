package lastfm

import (
	"log"
	"strconv"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/api"
)

// AllTopArtists returns all of a user's top artists through each page
func (lfm API) AllTopArtists(username string) ([]TopArtist, error) {
	params := TopEntityParams{Username: username, Limit: 1000, Page: 1}

	var topArtists []TopArtist

	err, firstPage := lfm.TopArtists(params)

	if err != nil {
		return topArtists, customerrors.LastFMError(err.Message, int(err.Error))
	}

	topArtists = append(topArtists, firstPage.TopArtists.Artists...)

	totalPages, _ := strconv.Atoi(firstPage.TopArtists.Attributes.TotalPages)

	paginator := helpers.Paginator{
		PageSize:      params.Limit,
		TotalPages:    totalPages,
		SkipFirstPage: true,

		Function: func(pp helpers.PagedParams) {
			log.Printf("Fetching page %d", pp.Page)

			params.Page = pp.Page

			_, response := lfm.TopArtists(params)

			topArtists = append(topArtists, response.TopArtists.Artists...)
		},
	}

	paginator.GetAll()

	return topArtists, nil
}

// AllTopAlbums returns all of a user's top albums through each page
func (lfm API) AllTopAlbums(username string) ([]TopAlbum, error) {
	params := TopEntityParams{Username: username, Limit: 1000, Page: 1}

	var topAlbums []TopAlbum

	err, firstPage := lfm.TopAlbums(params)

	if err != nil {
		return topAlbums, customerrors.LastFMError(err.Message, int(err.Error))
	}

	topAlbums = append(topAlbums, firstPage.TopAlbums.Albums...)

	totalPages, _ := strconv.Atoi(firstPage.TopAlbums.Attributes.TotalPages)

	paginator := helpers.Paginator{
		PageSize:      params.Limit,
		TotalPages:    totalPages,
		SkipFirstPage: true,

		Function: func(pp helpers.PagedParams) {
			log.Printf("Fetching page %d", pp.Page)

			params.Page = pp.Page

			_, response := lfm.TopAlbums(params)

			topAlbums = append(topAlbums, response.TopAlbums.Albums...)
		},
	}

	paginator.GetAll()

	return topAlbums, nil
}

// AllTopTracks returns all of a user's top tracks through each page
func (lfm API) AllTopTracks(username string) ([]TopTrack, error) {
	params := TopEntityParams{Username: username, Limit: 1000, Page: 1}

	var topTracks []TopTrack

	err, firstPage := lfm.TopTracks(params)

	if err != nil {
		return topTracks, customerrors.LastFMError(err.Message, int(err.Error))
	}

	topTracks = append(topTracks, firstPage.TopTracks.Tracks...)

	totalPages, _ := strconv.Atoi(firstPage.TopTracks.Attributes.TotalPages)

	paginator := helpers.Paginator{
		PageSize:      params.Limit,
		TotalPages:    totalPages,
		SkipFirstPage: true,

		Function: func(pp helpers.PagedParams) {
			log.Printf("Fetching page %d", pp.Page)

			params.Page = pp.Page

			_, response := lfm.TopTracks(params)

			topTracks = append(topTracks, response.TopTracks.Tracks...)
		},
	}

	paginator.GetAll()

	return topTracks, nil
}

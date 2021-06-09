package lastfm

import (
	"strconv"
	"time"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/api"
)

// AllTopArtists returns all of a user's top artists through each page
func (lfm API) AllTopArtists(requestable Requestable) ([]TopArtist, error) {
	params := TopEntityParams{Username: requestable, Limit: 1000, Page: 1}

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
			params.Page = pp.Page

			_, response := lfm.TopArtists(params)

			topArtists = append(topArtists, response.TopArtists.Artists...)
		},
	}

	paginator.GetAll()

	return topArtists, nil
}

// AllTopAlbums returns all of a user's top albums through each page
func (lfm API) AllTopAlbums(requestable Requestable) ([]TopAlbum, error) {
	params := TopEntityParams{Username: requestable, Limit: 1000, Page: 1}

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
			params.Page = pp.Page

			_, response := lfm.TopAlbums(params)

			topAlbums = append(topAlbums, response.TopAlbums.Albums...)
		},
	}

	paginator.GetAll()

	return topAlbums, nil
}

// AllTopTracks returns all of a user's top tracks through each page
func (lfm API) AllTopTracks(requestable Requestable) ([]TopTrack, error) {
	params := TopEntityParams{Username: requestable, Limit: 1000, Page: 1}

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
			params.Page = pp.Page

			_, response := lfm.TopTracks(params)

			topTracks = append(topTracks, response.TopTracks.Tracks...)
		},
	}

	paginator.GetAll()

	return topTracks, nil
}

// AllScrobblesSince returns all of a users scrobbles since a certain date
func (lfm API) AllScrobblesSince(requestable Requestable, since *time.Time) ([]RecentTrack, error) {
	var tracks []RecentTrack

	params := RecentTracksParams{
		Username: requestable,
		Period:   "overall",
		Limit:    1000,
	}

	if since != nil {
		params.From = strconv.FormatInt(since.UTC().Unix()-1, 10)
	}

	err, recentTracks := lfm.RecentTracks(params)

	if err != nil {
		return tracks, customerrors.LastFMError(err.Message, int(err.Error))
	}

	tracks = append(tracks, recentTracks.RecentTracks.Tracks...)

	if recentTracks.RecentTracks.Attributes.Total == "0" {
		return tracks, nil
	}

	if totalPages, _ := strconv.Atoi(recentTracks.RecentTracks.Attributes.TotalPages); totalPages > 1 {
		paginator := helpers.Paginator{
			PageSize:      1000,
			TotalPages:    totalPages,
			SkipFirstPage: true,

			Function: func(pp helpers.PagedParams) {
				params.Page = pp.Page
				_, response := lfm.RecentTracks(params)

				tracks = append(tracks, response.RecentTracks.Tracks...)
			},
		}

		paginator.GetAllInParallel(7)
	}

	return tracks, nil
}

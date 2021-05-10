package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
)

func PresentArtistTopAlbums(dbArtist *db.Artist, albumCounts []db.AlbumCount) *model.ArtistTopAlbumsResponse {
	artist := PresentArtist(dbArtist)

	topAlbumsResponse := &model.ArtistTopAlbumsResponse{
		Artist:    artist,
		TopAlbums: []*model.AlbumCount{},
	}

	for _, albumCount := range albumCounts {
		topAlbumsResponse.TopAlbums = append(topAlbumsResponse.TopAlbums, PresentAlbumCount(&albumCount))
	}

	return topAlbumsResponse
}

func PresentArtistTopTracks(dbArtist *db.Artist, trackCounts []analysis.AmbiguousTrackCount) *model.ArtistTopTracksResponse {
	artist := PresentArtist(dbArtist)

	topTracksResponse := &model.ArtistTopTracksResponse{
		Artist:    artist,
		TopTracks: []*model.AmbiguousTrackCount{},
	}

	for _, trackCount := range trackCounts {
		topTracksResponse.TopTracks = append(topTracksResponse.TopTracks, PresentAmbiguousTrackCount(&trackCount))
	}

	return topTracksResponse
}

func PresentArtistCount(artistCount *db.ArtistCount) *model.ArtistCount {
	builtPlay := &model.ArtistCount{
		Playcount: int(artistCount.Playcount),
	}

	if artistCount.Artist != nil {
		builtPlay.Artist = PresentArtist(artistCount.Artist)
	}

	return builtPlay
}

func PresentArtistCounts(artistCounts []db.ArtistCount) []*model.ArtistCount {
	var builtPlays []*model.ArtistCount

	for _, play := range artistCounts {
		builtPlays = append(builtPlays, PresentArtistCount(&play))
	}

	return builtPlays
}

func PresentAlbumCount(albumCount *db.AlbumCount) *model.AlbumCount {
	presentedCount := &model.AlbumCount{
		Playcount: int(albumCount.Playcount),
	}

	if albumCount.Album != nil {
		presentedCount.Album = PresentAlbum(albumCount.Album)
	}

	return presentedCount
}

func PresentAlbumCounts(albumCounts []db.AlbumCount) []*model.AlbumCount {
	var builtPlays []*model.AlbumCount

	for _, play := range albumCounts {
		builtPlays = append(builtPlays, PresentAlbumCount(&play))
	}

	return builtPlays
}

func PresentAmbiguousTrackCount(trackCount *analysis.AmbiguousTrackCount) *model.AmbiguousTrackCount {
	presentedCount := &model.AmbiguousTrackCount{
		Playcount: int(trackCount.Playcount),
		Name:      trackCount.Name,
	}

	return presentedCount
}

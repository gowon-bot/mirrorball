package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// PresentArtist converts a database artist into a graphql artist
func PresentArtist(artist *db.Artist) *model.Artist {
	return &model.Artist{
		ID:   int(artist.ID),
		Name: artist.Name,
	}
}

// PresentAlbum converts a database album into a graphql album
func PresentAlbum(album *db.Album) *model.Album {
	builtAlbum := &model.Album{
		ID:   int(album.ID),
		Name: album.Name,
	}

	if album.Artist != nil {
		builtAlbum.Artist = PresentArtist(album.Artist)
	}

	return builtAlbum
}

// PresentTrack converts a database track into a graphql track
func PresentTrack(track *db.Track) *model.Track {
	builtTrack := &model.Track{
		ID:   int(track.ID),
		Name: track.Name,
	}

	if track.Artist != nil {
		builtTrack.Artist = PresentArtist(track.Artist)
	}
	if track.Album != nil {
		builtTrack.Album = PresentAlbum(track.Album)
	}

	return builtTrack
}

// PresentAmbiguousTrack converts a database track into a graphql ambiguous track
func PresentAmbiguousTrack(tracks []db.Track) *model.AmbiguousTrack {
	track := tracks[0]

	builtTrack := &model.AmbiguousTrack{
		Name: track.Name,
	}

	if track.Artist != nil {
		builtTrack.Artist = track.Artist.Name
	}

	for _, track := range tracks {
		if track.Album != nil {
			builtTrack.Albums = append(builtTrack.Albums, PresentAlbum(track.Album))
		}
	}

	return builtTrack
}

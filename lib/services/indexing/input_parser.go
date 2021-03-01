package indexing

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// ParseArtistInput parses an artist input into sql
func ParseArtistInput(query *orm.Query, artistInput model.ArtistInput) *orm.Query {
	if artistInput.Name != nil {
		query = query.Where("artist.name ILIKE ?", artistInput.Name)
	}

	return query
}

// ParseAlbumInput parses an album input into sql
func ParseAlbumInput(query *orm.Query, albumInput model.AlbumInput) *orm.Query {
	if albumInput.Name != nil && len(*albumInput.Name) > 0 {
		query = query.Where("album.name ILIKE ?", albumInput.Name)
	}

	if albumInput.Artist != nil {
		query = ParseArtistInput(query, *albumInput.Artist)
	}

	return query
}

// ParseTrackInput parses a track input into sql
func ParseTrackInput(query *orm.Query, trackInput model.TrackInput) *orm.Query {
	if trackInput.Name != nil {
		query = query.Where("track.name ILIKE ?", trackInput.Name)
	}

	if trackInput.Artist != nil {
		query = ParseArtistInput(query, *trackInput.Artist)
	}

	if trackInput.Album != nil {
		query = ParseAlbumInput(query, *trackInput.Album)
	}

	return query
}

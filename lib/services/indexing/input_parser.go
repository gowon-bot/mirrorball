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

func ParseAlbumInputForAlbumCount(query *orm.Query, albumInput model.AlbumInput) *orm.Query {
	if albumInput.Name != nil && len(*albumInput.Name) > 0 {
		query = query.Where("album.name ILIKE ?", albumInput.Name)
	}

	if albumInput.Artist != nil {
		query = ParseArtistInputForAlbumCount(query, *albumInput.Artist)
	}

	return query
}

func ParseArtistInputForAlbumCount(query *orm.Query, artistInput model.ArtistInput) *orm.Query {
	if artistInput.Name != nil {
		query = query.Where("\"album__artist\".name ILIKE ?", artistInput.Name)
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

func ParsePageInput(query *orm.Query, pageInput *model.PageInput) *orm.Query {
	if pageInput == nil {
		return query
	}

	if pageInput.Limit != nil {
		query = query.Limit(*pageInput.Limit)
	}

	return query
}

func ParseArtistPlaysSettings(query *orm.Query, artistPlaysSettings *model.ArtistPlaysSettings) *orm.Query {
	if artistPlaysSettings.Artist != nil {
		query = ParseArtistInput(query, *artistPlaysSettings.Artist)
	}

	if artistPlaysSettings.PageInput != nil {
		query = ParsePageInput(query, artistPlaysSettings.PageInput)
	}

	if artistPlaysSettings.Sort != nil {
		query = query.Order(*artistPlaysSettings.Sort)
	} else {
		query = query.Order("playcount desc")
	}

	return query
}

func ParseAlbumPlaysSettings(query *orm.Query, albumPlaysSettings *model.AlbumPlaysSettings) *orm.Query {
	if albumPlaysSettings.Album != nil {
		query = ParseAlbumInputForAlbumCount(query, *albumPlaysSettings.Album)
	}

	if albumPlaysSettings.PageInput != nil {
		query = ParsePageInput(query, albumPlaysSettings.PageInput)
	}

	if albumPlaysSettings.Sort != nil {
		query = query.Order(*albumPlaysSettings.Sort)
	} else {
		query = query.Order("playcount desc")
	}

	return query
}

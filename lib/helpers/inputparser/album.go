package inputparser

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

type AlbumInputSettings interface {
	ArtistInputSettings
	getAlbumPath() string
}

func (p InputParser) ParseAlbumInput(albumInput model.AlbumInput, settings AlbumInputSettings) *InputParser {
	if albumInput.Name != nil && len(*albumInput.Name) > 0 {
		p.query.Where(settings.getAlbumPath()+".name ILIKE ?", albumInput.Name)
	}

	if albumInput.Artist != nil {
		p.ParseArtistInput(*albumInput.Artist, settings)
	}

	return &p
}

type AlbumPlaysInputSettings interface {
	AlbumInputSettings
	SortSettings
}

func (p InputParser) ParseAlbumPlaysSettings(albumPlaysSettings *model.AlbumPlaysSettings, settings AlbumPlaysInputSettings) *InputParser {
	if albumPlaysSettings.Album != nil {
		p.ParseAlbumInput(*albumPlaysSettings.Album, settings)
	}

	if albumPlaysSettings.PageInput != nil {
		p.ParsePageInput(albumPlaysSettings.PageInput)
	}

	p.ParseSort(albumPlaysSettings.Sort, settings)

	return &p
}

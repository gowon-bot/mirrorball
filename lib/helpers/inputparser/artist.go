package inputparser

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

type ArtistInputSettings interface {
	getArtistPath() string
}

func (p InputParser) ParseArtistInput(artistInput model.ArtistInput, settings ArtistInputSettings) *InputParser {
	if artistInput.Name != nil {
		p.query.Where(settings.getArtistPath()+".name ILIKE ?", artistInput.Name)
	}

	return &p
}

type ArtistPlaysInputSettings interface {
	ArtistInputSettings
	SortSettings
}

func (p InputParser) ParseArtistPlaysSettings(artistPlaysSettings *model.ArtistPlaysSettings, settings ArtistPlaysInputSettings) *InputParser {
	if artistPlaysSettings.Artist != nil {
		p.ParseArtistInput(*artistPlaysSettings.Artist, settings)
	}

	if artistPlaysSettings.PageInput != nil {
		p.ParsePageInput(artistPlaysSettings.PageInput)
	}

	p.ParseSort(artistPlaysSettings.Sort, settings)

	return &p
}

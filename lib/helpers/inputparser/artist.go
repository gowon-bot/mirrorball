package inputparser

import (
	"github.com/go-pg/pg/v10"
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

func (p InputParser) ParseArtistInputs(artistInputs []*model.ArtistInput, settings ArtistInputSettings) *InputParser {
	if len(artistInputs) == 0 {
		return &p
	}

	var artistNames []string

	for _, artist := range artistInputs {
		if artist != nil && artist.Name != nil {
			artistNames = append(artistNames, *artist.Name)
		}
	}

	p.query.Where(settings.getArtistPath()+".name IN (?)", pg.In(artistNames))

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

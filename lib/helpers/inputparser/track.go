package inputparser

import "github.com/jivison/gowon-indexer/lib/graph/model"

type TrackInputSettings interface {
	AlbumInputSettings
	getTrackPath() string
}

func (p InputParser) ParseTrackInput(trackInput model.TrackInput, settings TrackInputSettings) *InputParser {
	if trackInput.Name != nil {
		p.query.Where(settings.getTrackPath()+".name ILIKE ?", trackInput.Name)
	}

	if trackInput.Artist != nil {
		p.ParseArtistInput(*trackInput.Artist, settings)
	}

	if trackInput.Album != nil {
		p.ParseAlbumInput(*trackInput.Album, settings)
	}

	return &p
}

package inputparser

import "github.com/jivison/gowon-indexer/lib/graph/model"

type PlaysInputSettings interface {
	TrackInputSettings
	UserInputSettings
	SortSettings
}

func (p InputParser) ParsePlaysInput(playsInput *model.PlaysInput, settings PlaysInputSettings) *InputParser {
	if playsInput == nil {
		return &p
	}

	if playsInput.Track != nil {
		p.ParseTrackInput(*playsInput.Track, settings)
	}

	if playsInput.User != nil {
		p.ParseUserInput(*playsInput.User, settings)
	}

	if playsInput.Sort != nil {
		p.ParseSort(playsInput.Sort, settings)
	}

	if playsInput.Timerange != nil {
		p.ParseTimerange(*playsInput.Timerange, "scrobbled_at")
	}

	return &p
}

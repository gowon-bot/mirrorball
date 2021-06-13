package inputparser

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

type UserInputSettings interface {
	getUserRelation() string
}

func (p InputParser) ParseUserInput(userInput model.UserInput, settings UserInputSettings) *InputParser {
	p.query.Relation(settings.getUserRelation())

	if userInput.DiscordID != nil {
		p.query.Where("discord_id = ?", userInput.DiscordID)
	}

	if userInput.LastFMUsername != nil {
		p.query.Where("user_type = 'Lastfm'").Where("username = ?", userInput.LastFMUsername)
	}

	if userInput.WavyUsername != nil {
		p.query.Where("user_type = 'Wavy'").Where("username = ?", userInput.WavyUsername)
	}

	return &p
}

package inputparser

import (
	"github.com/go-pg/pg/v10"
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

func (p InputParser) ParseUserInputs(userInputs []*model.UserInput, settings UserInputSettings) *InputParser {
	if len(userInputs) == 0 {
		return &p
	}

	if settings.getUserRelation() != "-" {
		p.query.Relation(settings.getUserRelation())
	}

	var discordIDs []string
	var lastFMUsernames []string

	for _, input := range userInputs {
		if input == nil {
			continue
		}

		if input.DiscordID != nil {
			discordIDs = append(discordIDs, *input.DiscordID)
		}

		if input.LastFMUsername != nil {
			lastFMUsernames = append(lastFMUsernames, *input.LastFMUsername)
		}
	}

	if len(discordIDs) > 0 {
		p.query.Where("discord_id IN (?)", pg.In(discordIDs))
	}

	if len(lastFMUsernames) > 0 {
		p.query.Where("username IN (?)", pg.In(lastFMUsernames)).Where("user_type = 'Lastfm'")
	}

	return &p
}

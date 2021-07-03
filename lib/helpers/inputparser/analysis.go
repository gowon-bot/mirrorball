package inputparser

import (
	"fmt"

	"github.com/jivison/gowon-indexer/lib/graph/model"
)

type WhoKnowsSettings interface {
	getUserIDPath() string
}

func (p InputParser) ParseWhoKnowsSettings(settingsInput *model.WhoKnowsSettings, settings WhoKnowsSettings) *InputParser {
	if settingsInput == nil {
		return &p
	}

	if settingsInput.GuildID != nil {
		p.query.Join(`JOIN "guild_members" ON "guild_members"."guild_id" = ? AND "guild_members"."user_id" = "`+settings.getUserIDPath()+`"`, settingsInput.GuildID)
	}

	if settingsInput.Limit != nil {
		p.query.Limit(*settingsInput.Limit)
	}

	return &p
}

func (p InputParser) ParseArtistSearchCriteria(criteria model.ArtistSearchCriteria, searchSettings *model.SearchSettings) *InputParser {
	safeSearchSettings := model.SearchSettings{}

	if searchSettings != nil {
		safeSearchSettings = *searchSettings
	}

	if criteria.Keywords != nil {
		p.parseKeywords(*criteria.Keywords, safeSearchSettings.Exact)
	}

	if safeSearchSettings.User != nil {
		p.query.Relation("User")
		p.ParseUserInput(*safeSearchSettings.User, InputSettings{})
	}

	return &p
}

func (p InputParser) parseKeywords(keywords string, exact *bool) *InputParser {
	if exact != nil && *exact {
		p.query.Where("name ~ ?", fmt.Sprintf("(^|\\s)%s($|\\s)", keywords))
	} else {
		p.query.Where("name ilike ?", fmt.Sprintf("%%%s%%", keywords))
	}

	return &p
}

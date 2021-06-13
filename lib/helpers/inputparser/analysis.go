package inputparser

import (
	"fmt"

	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func (p InputParser) ParseWhoKnowsSettings(settings *model.WhoKnowsSettings) *InputParser {
	if settings == nil {
		return &p
	}

	if settings.GuildID != nil {
		p.query.Relation("User").Join(`JOIN "guild_members" ON "guild_members"."guild_id" = ? AND "guild_members"."user_id" = "user"."id"`, settings.GuildID)
	}

	if settings.Limit != nil {
		p.query.Limit(*settings.Limit)
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

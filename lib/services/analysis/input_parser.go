package analysis

import (
	"fmt"

	"github.com/go-pg/pg/v10/orm"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

// ParseWhoKnowsSettings parses who knows settings input into sql
func ParseWhoKnowsSettings(query *orm.Query, settings *model.WhoKnowsSettings) *orm.Query {
	if settings == nil {
		return query
	}

	if settings.GuildID != nil {
		query = query.Relation("User").Join(`JOIN "guild_members" ON "guild_members"."guild_id" = ? AND "guild_members"."user_id" = "user"."id"`, settings.GuildID)
	}

	if settings.Limit != nil {
		query = query.Limit(*settings.Limit)
	}

	return query
}

func ParseArtistSearchCriteria(query *orm.Query, criteria model.ArtistSearchCriteria, searchSettings *model.SearchSettings) *orm.Query {
	safeSearchSettings := model.SearchSettings{}

	if searchSettings != nil {
		safeSearchSettings = *searchSettings
	}

	if criteria.Keywords != nil {
		query = parseKeywords(query, *criteria.Keywords, safeSearchSettings.Exact)
	}

	if safeSearchSettings.User != nil {
		usersService := users.CreateService()

		user := usersService.FindUserByInput(*safeSearchSettings.User)

		if user != nil {
			query = query.Where("user_id = ?", user.ID)
		}
	}

	return query
}

func parseKeywords(query *orm.Query, keywords string, exact *bool) *orm.Query {
	if exact != nil && *exact == true {
		return query.Where("name ~ ?", fmt.Sprintf("(^|\\s)%s($|\\s)", keywords))
	} else {
		return query.Where("name ilike ?", fmt.Sprintf("%%%s%%", keywords))
	}
}

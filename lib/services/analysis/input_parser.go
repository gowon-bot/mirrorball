package analysis

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/jivison/gowon-indexer/lib/graph/model"
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

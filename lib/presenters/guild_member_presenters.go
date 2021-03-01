package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// PresentGuildMember converts a database guild member into a graphql guild member
func PresentGuildMember(guildMember *db.GuildMember) *model.GuildMember {
	var user *model.User

	if guildMember.User != nil {
		user = PresentUser(guildMember.User)
	}

	return &model.GuildMember{
		UserID:  int(guildMember.UserID),
		GuildID: guildMember.GuildID,
		User:    user,
	}
}

// PresentGuildMembers converts a list of database guild members into a a list of graphql guild members
func PresentGuildMembers(guildMembers []db.GuildMember) []*model.GuildMember {
	var builtGuildMembers []*model.GuildMember

	for _, guildMember := range guildMembers {
		builtGuildMembers = append(builtGuildMembers, PresentGuildMember(&guildMember))
	}

	return builtGuildMembers
}

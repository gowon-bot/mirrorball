package guildmembers

import (
	"strings"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// GuildMembers holds methods for interacting with guild members
type GuildMembers struct{}

// CreateGuildMember creates a guild member in the database
func (gm GuildMembers) CreateGuildMember(user *db.User, guildID string) (*db.GuildMember, error) {
	guildMember := &db.GuildMember{
		GuildID: guildID,

		User:   user,
		UserID: user.ID,
	}

	_, err := db.Db.Model(guildMember).Insert()

	if err != nil {
		// "gm_uniqueness" is the name of the sql unique constraint for guild_members
		if strings.Contains(err.Error(), "gm_uniqueness") {
			return nil, customerrors.EntityAlreadyExistsError("guild member")
		}
		return nil, customerrors.DatabaseUnknownError()
	}

	return guildMember, nil
}

// DeleteGuildMember deletes a guild member in the database
func (gm GuildMembers) DeleteGuildMember(user *db.User, guildID string) error {
	_, err := db.Db.Model(&db.GuildMember{}).
		Where("user_id = ?", user.ID).
		Where("guild_id = ?", guildID).
		Delete()

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}

// ClearGuild removes all the guild members from a guild in the database
func (gm GuildMembers) ClearGuild(guildID string) error {
	_, err := db.Db.Model(&model.GuildMember{}).Where("guild_id = ?", guildID).Delete()

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}

// CreateService creates an instance of the guild members service object
func CreateService() *GuildMembers {
	service := &GuildMembers{}

	return service
}

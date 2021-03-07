package controllers

import (
	"fmt"
	"log"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/presenters"
	guildmembers "github.com/jivison/gowon-indexer/lib/services/guild_members"
	"github.com/jivison/gowon-indexer/lib/services/users"
)

// AddUserToGuild adds a member to a guild in the cache
func AddUserToGuild(discordID, guildID string) (*model.GuildMember, error) {
	log.Print(fmt.Sprintf("Adding %s to %s", discordID, guildID))

	usersService := users.CreateService()
	guildMembersService := guildmembers.CreateService()

	user := usersService.FindUserByDiscordID(discordID)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	guildMember, err := guildMembersService.CreateGuildMember(user, guildID)

	if err != nil {
		return nil, err
	}

	return presenters.PresentGuildMember(guildMember), nil
}

// RemoveUserFromGuild removes a member from a guild in the cache
func RemoveUserFromGuild(discordID, guildID string) (*string, error) {
	usersService := users.CreateService()
	guildMembersService := guildmembers.CreateService()

	user := usersService.FindUserByDiscordID(discordID)

	if user == nil {
		return nil, customerrors.EntityDoesntExistError("user")
	}

	err := guildMembersService.DeleteGuildMember(user, guildID)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

// SyncGuild syncs a list of members with a guild
func SyncGuild(discordIDs []string, guildID string) (*string, error) {
	guildMembersService := guildmembers.CreateService()
	usersService := users.CreateService()

	err := guildMembersService.ClearGuild(guildID)

	if err != nil {
		return nil, err
	}

	users, err := usersService.GetUsersByDiscordIDs(discordIDs)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		guildMembersService.CreateGuildMember(user, guildID)
	}

	return nil, nil
}

// GuildMembers lists the guild members in a given server
func GuildMembers(guildID string) ([]*model.GuildMember, error) {
	guildMembersService := guildmembers.CreateService()

	guildMembers, err := guildMembersService.ListGuildMembers(guildID)

	if err != nil {
		return nil, err
	}

	return presenters.PresentGuildMembers(guildMembers), nil
}

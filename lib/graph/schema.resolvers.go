package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jivison/gowon-indexer/lib/controllers"
	"github.com/jivison/gowon-indexer/lib/graph/generated"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, username string, discordID string, userType model.UserType) (*model.User, error) {
	return controllers.Login(username, discordID, userType.String())
}

func (r *mutationResolver) Logout(ctx context.Context, discordID string) (*string, error) {
	return controllers.Logout(discordID)
}

func (r *mutationResolver) AddUserToGuild(ctx context.Context, discordID string, guildID string) (*model.GuildMember, error) {
	return controllers.AddUserToGuild(discordID, guildID)
}

func (r *mutationResolver) RemoveUserFromGuild(ctx context.Context, discordID string, guildID string) (*string, error) {
	return controllers.RemoveUserFromGuild(discordID, guildID)
}

func (r *mutationResolver) SyncGuild(ctx context.Context, guildID string, discordIDs []string) (*string, error) {
	return controllers.SyncGuild(discordIDs, guildID)
}

func (r *mutationResolver) FullIndex(ctx context.Context, user model.UserInput) (*model.TaskStartResponse, error) {
	return controllers.FullIndex(user)
}

func (r *mutationResolver) Update(ctx context.Context, user model.UserInput) (*model.TaskStartResponse, error) {
	return controllers.Update(user)
}

func (r *queryResolver) WhoKnowsArtist(ctx context.Context, artist model.ArtistInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsArtistResponse, error) {
	return controllers.WhoKnowsArtist(artist, settings)
}

func (r *queryResolver) WhoKnowsAlbum(ctx context.Context, album model.AlbumInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsAlbumResponse, error) {
	return controllers.WhoKnowsAlbum(album, settings)
}

func (r *queryResolver) WhoKnowsTrack(ctx context.Context, track model.TrackInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsTrackResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GuildMembers(ctx context.Context, guildID string) ([]*model.GuildMember, error) {
	return controllers.GuildMembers(guildID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

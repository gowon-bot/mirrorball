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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveUserFromGuild(ctx context.Context, discordID string, guildID string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SyncGuild(ctx context.Context, guildID string, discordIDs []int) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FullIndex(ctx context.Context, user model.UserInput) (*model.TaskStartResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Update(ctx context.Context, user model.UserInput) (*model.TaskStartResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WhoKnowsArtist(ctx context.Context, artist model.ArtistInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsArtistResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WhoKnowsAlbum(ctx context.Context, album model.AlbumInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsAlbumResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WhoKnowsTrack(ctx context.Context, track model.TrackInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsTrackResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

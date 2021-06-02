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

func (r *mutationResolver) FullIndex(ctx context.Context, user model.UserInput, forceUserCreate *bool) (*model.TaskStartResponse, error) {
	return controllers.FullIndex(user, forceUserCreate)
}

func (r *mutationResolver) Update(ctx context.Context, user model.UserInput, forceUserCreate *bool) (*model.TaskStartResponse, error) {
	return controllers.Update(user, forceUserCreate)
}

func (r *mutationResolver) ImportRatings(ctx context.Context, csv string, user model.UserInput) (*string, error) {
	return controllers.ImportRatings(csv, user)
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return controllers.Ping()
}

func (r *queryResolver) WhoKnowsArtist(ctx context.Context, artist model.ArtistInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsArtistResponse, error) {
	return controllers.WhoKnowsArtist(artist, settings)
}

func (r *queryResolver) WhoKnowsAlbum(ctx context.Context, album model.AlbumInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsAlbumResponse, error) {
	return controllers.WhoKnowsAlbum(album, settings)
}

func (r *queryResolver) WhoKnowsTrack(ctx context.Context, track model.TrackInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsTrackResponse, error) {
	return controllers.WhoKnowsTrack(track, settings)
}

func (r *queryResolver) GuildMembers(ctx context.Context, guildID string) ([]*model.GuildMember, error) {
	return controllers.GuildMembers(guildID)
}

func (r *queryResolver) ArtistTopTracks(ctx context.Context, user model.UserInput, artist model.ArtistInput) (*model.ArtistTopTracksResponse, error) {
	return controllers.ArtistTopTracks(user, artist)
}

func (r *queryResolver) ArtistTopAlbums(ctx context.Context, user model.UserInput, artist model.ArtistInput) (*model.ArtistTopAlbumsResponse, error) {
	return controllers.ArtistTopAlbums(user, artist)
}

func (r *queryResolver) AlbumTopTracks(ctx context.Context, user model.UserInput, album model.AlbumInput) (*model.AlbumTopTracksResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchArtist(ctx context.Context, criteria model.ArtistSearchCriteria, settings *model.SearchSettings) (*model.ArtistSearchResults, error) {
	return controllers.SearchArtist(criteria, settings)
}

func (r *queryResolver) Plays(ctx context.Context, user model.UserInput, pageInput *model.PageInput) ([]*model.Play, error) {
	return controllers.Plays(user, pageInput)
}

func (r *queryResolver) ArtistPlays(ctx context.Context, user model.UserInput, settings *model.ArtistPlaysSettings) ([]*model.ArtistCount, error) {
	return controllers.ArtistPlays(user, settings)
}

func (r *queryResolver) AlbumPlays(ctx context.Context, user model.UserInput, settings *model.AlbumPlaysSettings) ([]*model.AlbumCount, error) {
	return controllers.AlbumPlays(user, settings)
}

func (r *queryResolver) TrackPlays(ctx context.Context, user model.UserInput, settings *model.TrackPlaysSettings) ([]*model.AmbiguousTrackCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ratings(ctx context.Context, settings *model.RatingsSettings) ([]*model.Rating, error) {
	return controllers.Ratings(settings)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

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

func (r *mutationResolver) Login(ctx context.Context, username string, session *string, discordID string, userType model.UserType) (*model.User, error) {
	return controllers.Login(username, session, discordID, userType.String())
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

func (r *mutationResolver) TagArtists(ctx context.Context, artists []*model.ArtistInput, tags []*model.TagInput, markAsChecked *bool) (*string, error) {
	return controllers.TagArtists(artists, tags, markAsChecked)
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return controllers.Ping()
}

func (r *queryResolver) Artists(ctx context.Context, inputs []*model.ArtistInput, tag *model.TagInput, pageInput *model.PageInput, requireTagsForMissing *bool) ([]*model.Artist, error) {
	return controllers.Artists(ctx, inputs, tag, requireTagsForMissing)
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

func (r *queryResolver) ArtistRank(ctx context.Context, artist model.ArtistInput, userInput model.UserInput, serverID *string) (*model.ArtistRankResponse, error) {
	return controllers.ArtistRank(artist, userInput, serverID)
}

func (r *queryResolver) WhoFirstArtist(ctx context.Context, artist model.ArtistInput, settings *model.WhoKnowsSettings, whoLast *bool) (*model.WhoFirstArtistResponse, error) {
	return controllers.WhoFirstArtist(artist, settings, whoLast)
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
	return controllers.AlbumTopTracks(user, album)
}

func (r *queryResolver) TrackTopAlbums(ctx context.Context, user model.UserInput, track model.TrackInput) (*model.TrackTopAlbumsResponse, error) {
	return controllers.TrackTopAlbums(user, track)
}

func (r *queryResolver) SearchArtist(ctx context.Context, criteria model.ArtistSearchCriteria, settings *model.SearchSettings) (*model.ArtistSearchResults, error) {
	return controllers.SearchArtist(criteria, settings)
}

func (r *queryResolver) Plays(ctx context.Context, playsInput model.PlaysInput, pageInput *model.PageInput) (*model.PlaysResponse, error) {
	return controllers.Plays(playsInput, pageInput)
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

func (r *queryResolver) Ratings(ctx context.Context, settings *model.RatingsSettings) (*model.RatingsResponse, error) {
	return controllers.Ratings(settings)
}

func (r *queryResolver) RateYourMusicArtist(ctx context.Context, keywords string) (*model.RateYourMusicArtist, error) {
	return controllers.RateYourMusicArtist(keywords)
}

func (r *queryResolver) Tags(ctx context.Context, settings *model.TagsSettings, requireTagsForMissing *bool) (*model.TagsResponse, error) {
	return controllers.Tags(settings, requireTagsForMissing)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jivison/gowon-indexer/lib/controllers"
	"github.com/jivison/gowon-indexer/lib/graph/generated"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/meta"
)

func (r *mutationResolver) Login(ctx context.Context, username string, session *string, discordID string) (*model.User, error) {
	err := meta.CheckUserMatches(ctx, discordID)

	if err != nil {
		return nil, err
	}

	return controllers.Login(username, session, discordID)
}

func (r *mutationResolver) Logout(ctx context.Context, discordID string) (*string, error) {
	err := meta.CheckUserMatches(ctx, discordID)

	if err != nil {
		return nil, err
	}

	return controllers.Logout(discordID)
}

func (r *mutationResolver) UpdatePrivacy(ctx context.Context, user model.UserInput, privacy *model.Privacy) (*string, error) {
	err := meta.CheckUserMatches(ctx, *user.DiscordID)

	if err != nil {
		return nil, err
	}

	return controllers.UpdatePrivacy(user, privacy)
}

func (r *mutationResolver) AddUserToGuild(ctx context.Context, discordID string, guildID string) (*model.GuildMember, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.AddUserToGuild(discordID, guildID)
}

func (r *mutationResolver) RemoveUserFromGuild(ctx context.Context, discordID string, guildID string) (*string, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.RemoveUserFromGuild(discordID, guildID)
}

func (r *mutationResolver) SyncGuild(ctx context.Context, guildID string, discordIDs []string) (*string, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.SyncGuild(discordIDs, guildID)
}

func (r *mutationResolver) DeleteGuild(ctx context.Context, guildID string) (*string, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.DeleteGuild(guildID)
}

func (r *mutationResolver) FullIndex(ctx context.Context, user model.UserInput, forceUserCreate *bool) (*model.TaskStartResponse, error) {
	err := meta.CheckUserMatches(ctx, *user.DiscordID)

	if err != nil {
		return nil, err
	}

	return controllers.FullIndex(user, forceUserCreate)
}

func (r *mutationResolver) Update(ctx context.Context, user model.UserInput, forceUserCreate *bool) (*model.TaskStartResponse, error) {
	err := meta.CheckUserMatches(ctx, *user.DiscordID)

	if err != nil {
		return nil, err
	}

	return controllers.Update(user, forceUserCreate)
}

func (r *mutationResolver) ImportRatings(ctx context.Context, csv string, user model.UserInput) (*string, error) {
	err := meta.CheckUserMatches(ctx, *user.DiscordID)

	if err != nil {
		return nil, err
	}
	return controllers.ImportRatings(csv, user)
}

func (r *mutationResolver) TagArtists(ctx context.Context, artists []*model.ArtistInput, tags []*model.TagInput, markAsChecked *bool) (*string, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.TagArtists(artists, tags, markAsChecked)
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return controllers.Ping()
}

func (r *queryResolver) Artists(ctx context.Context, inputs []*model.ArtistInput, tag *model.TagInput, pageInput *model.PageInput, requireTagsForMissing *bool) ([]*model.Artist, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.Artists(ctx, inputs, tag, requireTagsForMissing)
}

func (r *queryResolver) WhoKnowsArtist(ctx context.Context, artist model.ArtistInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsArtistResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.WhoKnowsArtist(artist, settings)
}

func (r *queryResolver) WhoKnowsAlbum(ctx context.Context, album model.AlbumInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsAlbumResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.WhoKnowsAlbum(album, settings)
}

func (r *queryResolver) WhoKnowsTrack(ctx context.Context, track model.TrackInput, settings *model.WhoKnowsSettings) (*model.WhoKnowsTrackResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.WhoKnowsTrack(track, settings)
}

func (r *queryResolver) ArtistRank(ctx context.Context, artist model.ArtistInput, userInput model.UserInput, serverID *string) (*model.ArtistRankResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.ArtistRank(artist, userInput, serverID)
}

func (r *queryResolver) WhoFirstArtist(ctx context.Context, artist model.ArtistInput, settings *model.WhoKnowsSettings, whoLast *bool) (*model.WhoFirstArtistResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.WhoFirstArtist(artist, settings, whoLast)
}

func (r *queryResolver) GuildMembers(ctx context.Context, guildID string) ([]*model.GuildMember, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.GuildMembers(guildID)
}

func (r *queryResolver) Users(ctx context.Context, inputs []*model.UserInput) ([]*model.User, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.Users(inputs)
}

func (r *queryResolver) ArtistTopTracks(ctx context.Context, user model.UserInput, artist model.ArtistInput) (*model.ArtistTopTracksResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.ArtistTopTracks(user, artist)
}

func (r *queryResolver) ArtistTopAlbums(ctx context.Context, user model.UserInput, artist model.ArtistInput) (*model.ArtistTopAlbumsResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.ArtistTopAlbums(user, artist)
}

func (r *queryResolver) AlbumTopTracks(ctx context.Context, user model.UserInput, album model.AlbumInput) (*model.AlbumTopTracksResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.AlbumTopTracks(user, album)
}

func (r *queryResolver) TrackTopAlbums(ctx context.Context, user model.UserInput, track model.TrackInput) (*model.TrackTopAlbumsResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.TrackTopAlbums(user, track)
}

func (r *queryResolver) SearchArtist(ctx context.Context, criteria model.ArtistSearchCriteria, settings *model.SearchSettings) (*model.ArtistSearchResults, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.SearchArtist(criteria, settings)
}

func (r *queryResolver) Plays(ctx context.Context, playsInput model.PlaysInput, pageInput *model.PageInput) (*model.PlaysResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.Plays(playsInput, pageInput)
}

func (r *queryResolver) ArtistPlays(ctx context.Context, user model.UserInput, settings *model.ArtistPlaysSettings) ([]*model.ArtistCount, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.ArtistPlays(user, settings)
}

func (r *queryResolver) AlbumPlays(ctx context.Context, user model.UserInput, settings *model.AlbumPlaysSettings) ([]*model.AlbumCount, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.AlbumPlays(user, settings)
}

func (r *queryResolver) TrackPlays(ctx context.Context, user model.UserInput, settings *model.TrackPlaysSettings) ([]*model.AmbiguousTrackCount, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ratings(ctx context.Context, settings *model.RatingsSettings) (*model.RatingsResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.Ratings(settings)
}

func (r *queryResolver) RateYourMusicArtist(ctx context.Context, keywords string) (*model.RateYourMusicArtist, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.RateYourMusicArtist(keywords)
}

func (r *queryResolver) Tags(ctx context.Context, settings *model.TagsSettings, requireTagsForMissing *bool) (*model.TagsResponse, error) {
	err := meta.CheckNoUser(ctx)

	if err != nil {
		return nil, err
	}

	return controllers.Tags(settings, requireTagsForMissing)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

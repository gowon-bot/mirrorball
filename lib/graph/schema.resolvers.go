package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	customerrors "github.com/jivison/gowon-indexer/lib/customErrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/generated"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/indexeddata"
	"github.com/jivison/gowon-indexer/lib/services/response"
	"github.com/jivison/gowon-indexer/lib/services/user"
	"github.com/jivison/gowon-indexer/lib/tasks"
)

func (r *mutationResolver) IndexUser(ctx context.Context, username string) (*model.TaskStartResponse, error) {
	responseService := response.CreateService()

	token := responseService.GenerateToken()

	tasks.TaskServer.SendIndexUserTask(username, token)

	return responseService.BuildTaskStartResponse(token), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, username string) (*model.TaskStartResponse, error) {
	responseService := response.CreateService()

	token := responseService.GenerateToken()

	tasks.TaskServer.SendUpdateUserTask(username, token)

	return responseService.BuildTaskStartResponse(token), nil
}

func (r *mutationResolver) SaveTrack(ctx context.Context, artist string, album *string, track string) (*model.Track, error) {
	indexedService := indexeddata.CreateIndexedMutationService()

	savedTrack := indexedService.SaveTrack(artist, track, album)

	gqlTrack := indexeddata.ConvertToGraphQLTrack(savedTrack)

	return gqlTrack, nil
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "Hello from Go!", nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var dbUsers []db.User
	var resultUsers []*model.User

	err := db.Db.Model(&dbUsers).Select()

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range dbUsers {
		resultUsers = append(resultUsers, &model.User{
			ID:             int(user.ID),
			LastFMUsername: user.LastFMUsername,
		})
	}

	return resultUsers, nil
}

func (r *queryResolver) GetUser(ctx context.Context, username string) (*model.User, error) {
	userService := user.CreateService()

	user, err := userService.GetUser(username)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:             int(user.ID),
		LastFMUsername: user.LastFMUsername,
	}, nil
}

func (r *queryResolver) UserTopArtists(ctx context.Context, username string) (int, error) {
	indexedQuery := indexeddata.CreateIndexedQueryService()

	count := indexedQuery.UserTopArtists(username)

	return count, nil
}

func (r *queryResolver) WhoKnows(ctx context.Context, artist string) (*model.WhoKnowsResponse, error) {
	indexedQuery := indexeddata.CreateIndexedQueryService()
	indexedMutation := indexeddata.CreateIndexedMutationService()

	dbArtist, err := indexedMutation.GetArtist(artist, false)
	gqlArtist := indexeddata.ConvertToGraphQLArtist(dbArtist)

	if err != nil {
		return nil, customerrors.NoOneKnows(artist)
	}

	whoKnows := indexedQuery.WhoKnowsArtist(dbArtist)

	var whoKnowsResponse []*model.WhoKnows

	for _, whoKnowsRow := range whoKnows {
		whoKnowsResponse = append(whoKnowsResponse, &model.WhoKnows{
			Artist: gqlArtist,
			User: &model.User{
				ID:             int(whoKnowsRow.User.ID),
				LastFMUsername: whoKnowsRow.User.LastFMUsername,
			},
			Playcount: int(whoKnowsRow.Playcount),
		})
	}

	return &model.WhoKnowsResponse{
		Artist: gqlArtist,
		Users:  whoKnowsResponse,
	}, nil
}

func (r *queryResolver) WhoKnowsAlbum(ctx context.Context, artist string, album string) (*model.WhoKnowsAlbumResponse, error) {
	indexedQuery := indexeddata.CreateIndexedQueryService()
	indexedMutation := indexeddata.CreateIndexedMutationService()

	dbAlbum, err := indexedMutation.GetAlbum(album, artist, false)
	gqlAlbum := indexeddata.ConvertToGraphQLAlbum(dbAlbum)

	if err != nil {
		return nil, err
	}

	whoKnows := indexedQuery.WhoKnowsAlbum(dbAlbum)

	var whoKnowsResponse []*model.WhoKnowsAlbum

	for _, whoKnowsRow := range whoKnows {
		whoKnowsResponse = append(whoKnowsResponse, &model.WhoKnowsAlbum{
			Album: gqlAlbum,
			User: &model.User{
				ID:             int(whoKnowsRow.User.ID),
				LastFMUsername: whoKnowsRow.User.LastFMUsername,
			},
			Playcount: int(whoKnowsRow.Playcount),
		})
	}

	return &model.WhoKnowsAlbumResponse{
		Album: gqlAlbum,
		Users: whoKnowsResponse,
	}, nil
}

func (r *queryResolver) WhoKnowsTrack(ctx context.Context, artist string, track string) (*model.WhoKnowsTrackResponse, error) {
	indexedQuery := indexeddata.CreateIndexedQueryService()
	indexedMutation := indexeddata.CreateIndexedMutationService()

	dbTracks, err := indexedMutation.GetTracks(track, artist)

	if err != nil {
		return nil, err
	}

	ambiguousTrack := &indexeddata.AmbiguousTrack{
		Name:   dbTracks[0].Name,
		Artist: dbTracks[0].Artist,
	}

	gqlTrack := indexeddata.ConvertToAmbiguousTrack(ambiguousTrack)

	whoKnowsTrack := indexedQuery.WhoKnowsTrack(dbTracks)

	var whoKnowsResponse []*model.WhoKnowsTrack

	for _, whoKnowsRow := range whoKnowsTrack {
		whoKnowsResponse = append(whoKnowsResponse, &model.WhoKnowsTrack{
			Track:     gqlTrack,
			Playcount: int(whoKnowsRow.Playcount),
			User: &model.User{
				ID:             int(whoKnowsRow.User.ID),
				LastFMUsername: whoKnowsRow.User.LastFMUsername,
			},
		})
	}

	return &model.WhoKnowsTrackResponse{
		Track: gqlTrack,
		Users: whoKnowsResponse,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

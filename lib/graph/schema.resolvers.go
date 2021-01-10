package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/generated"
	"github.com/jivison/gowon-indexer/lib/graph/model"
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

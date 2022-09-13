package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gaochuwuhan/gql/graph/generated"
	"github.com/gaochuwuhan/gql/graph/model"
	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	//panic(fmt.Errorf("not implemented: CreateUser - createUser"))
	coll := r.DB.Database(db).Collection("user")
	u := &model.User{
		ID:        uuid.New().String(),
		Name:      input.Name,
		CreatedOn: int(time.Now().Unix()),
		UpdatedOn: int(time.Now().Unix()),
		Deleted:   false,
		IsActive:  input.IsActive,
	}
	_, err := coll.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	log.Printf("get user ...%s", id)
	//panic(fmt.Errorf("not implemented: GetUser - getUser"))
	return &model.User{
			ID:        uuid.New().String(),
			Name:      "mock_user1",
			CreatedOn: int(time.Now().Unix()),
			UpdatedOn: int(time.Now().Unix()),
			Deleted:   true,
		},
		nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context, updateAt int, limit int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUsers - getUsers"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const (
	db = "test-module"
)

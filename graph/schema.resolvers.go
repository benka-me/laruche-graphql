package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/benka-me/laruche-graphql/graph/generated"
	"github.com/benka-me/laruche-graphql/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterReq) (*model.RegisterRes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, username string, password string) (*model.LoginRes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetHome(ctx context.Context, input model.HomeReq) (*model.Home, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetBee(ctx context.Context, input model.BeeReq) (*model.Bee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetBeeDetails(ctx context.Context, input model.BeeReq) (*model.BeeDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetFullBee(ctx context.Context, input model.BeeReq) (*model.FullBee, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

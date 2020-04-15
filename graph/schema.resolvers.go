package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/benka-me/laruche-graphql/graph/generated"
	"github.com/benka-me/laruche-graphql/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterReq) (*model.RegisterRes, error) {
	return &model.RegisterRes{
		Status:        true,
		StatusMessage: "test PitchValue",
	}, nil
}

func (r *queryResolver) Login(ctx context.Context, username string, password string) (*model.LoginRes, error) {
	return &model.LoginRes{
		Status:   true,
		TokenErr: "Test PitchValue",
	}, nil
}

func (r *queryResolver) GetHome(ctx context.Context, input model.HomeReq) (*model.Home, error) {
	ret := &model.Home{
		Bees: Bees,
		Hives: []*model.Hive{
			{
				Author: "queen",
				Name:   "basic-admin-ui",
				Repo:   "github.com/queen/basic-admin-ui",
			},
		},
	}
	return ret, nil
}

func (r *queryResolver) GetBee(ctx context.Context, input model.BeeReq) (*model.Bee, error) {
	return Bees[0], nil
}

func (r *queryResolver) GetBeeDetails(ctx context.Context, input model.BeeReq) (*model.BeeDetails, error) {
	return BeesDetails[0], nil
}

func (r *queryResolver) GetFullBee(ctx context.Context, input model.BeeReq) (*model.FullBee, error) {
	return &model.FullBee{
		Bee:     Bees[0],
		Details: BeesDetails[0],
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var BeesDetails = []*model.BeeDetails{
	{
		Readme: "Readme",
	},
}
var Bees = []*model.Bee{
	{
		Author:      "queen",
		Name:        "basic-log",
		Description: "Save and spread log from others micro-services.",
		PkgName:     "log",
		Repo:        "github.com/queen/basic-log",
		Port:        4556,
		License:     "MIT",
		Keywords:    []string{"golang", "go", "log", "micro-service"},
		DevLang:     0,
		Framework:   &model.Framework{Name: "basic"},
		Rpcs: []*model.RPC{
			{Line: "rpc Register (admin.RegisterReq) returns (admin.RegisterRes)"},
			{Line: "rpc Login (admin.LoginReq) returns (admin.RegisterRes)"},
		},
	},
	{
		Author:      "queen",
		Name:        "basic-user",
		Description: "Manage users login, register etc....",
		PkgName:     "user",
		Repo:        "github.com/queen/basic-user",
		Port:        4555,
		License:     "MIT",
		Keywords:    []string{"user", "micro-service"},
		DevLang:     1,
		Framework:   &model.Framework{Name: "basic"},
		Rpcs: []*model.RPC{
			{Line: "rpc Register (admin.RegisterReq) returns (admin.RegisterRes)"},
			{Line: "rpc Login (admin.LoginReq) returns (admin.RegisterRes)"},
		},
	},
	{
		Author:      "queen",
		Name:        "basic-admin",
		Description: "This micro-service register admin, login them, and provide auth",
		PkgName:     "admin",
		Repo:        "github.com/queen/basic-admin",
		Port:        4554,
		License:     "MIT",
		Keywords:    []string{"admin", "micro-service"},
		DevLang:     2,
		Framework:   &model.Framework{Name: "basic"},
		Rpcs: []*model.RPC{
			{Line: "rpc Register (admin.RegisterReq) returns (admin.RegisterRes)"},
			{Line: "rpc Login (admin.LoginReq) returns (admin.RegisterRes)"},
		},
	},
}

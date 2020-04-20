package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"google.golang.org/grpc/status"
	"log"

	"github.com/benka-me/laruche-graphql/graph/generated"
	"github.com/benka-me/laruche-graphql/graph/model"
	"github.com/benka-me/users/go-pkg/users"
)

func (r *queryResolver) Login(ctx context.Context, username string, password string) (*model.LoginRes, error) {
	fmt.Println("login: ", username, password)
	res, err := r.Clients.Users.Login(ctx, &users.LoginReq{
		Identifier: username,
		Password:   password,
	})
	if err != nil {
		return &model.LoginRes{
			Status:   false,
			TokenErr: "invalid password or username",
		}, nil
	}
	return &model.LoginRes{
		Status:   true,
		TokenErr: res.Auth,
	}, nil
}

func (r *queryResolver) GetHome(ctx context.Context, input model.HomeReq) (*model.Home, error) {
	ret := &model.Home{
		Bees:  []*model.Bee{},
		Hives: []*model.Hive{},
	}
	auth, _ := r.Clients.Users.Auth(ctx, &users.Token{Val: input.Token})
	if !auth.Val {
		return ret, errors.New("Your are not autenticated")
	}

	res, err := r.Clients.Larsrv.GetBees(ctx, nil)
	if err != nil {
		return ret, err
	}

	ret.Bees = model.ToBeez(res.Bees)
	return ret, nil
}

func (r *queryResolver) GetBee(ctx context.Context, input model.BeeReq) (*model.Bee, error) {
	auth, _ := r.Clients.Users.Auth(ctx, &users.Token{Val: input.Auth.Token})
	if !auth.Val {
		return &model.Bee{}, errors.New("Your are not autenticated")
	}
	res, err := r.Clients.Larsrv.GetBee(ctx, &laruche.BeeReq{
		BeeName: input.Author + "/" + input.Name,
	})
	if err != nil {
		return &model.Bee{}, err
	}
	return model.ToBee(res), nil
}

func (r *queryResolver) GetBeeDetails(ctx context.Context, input model.BeeReq) (*model.BeeDetails, error) {
	return &model.BeeDetails{}, nil
}

func (r *queryResolver) GetFullBee(ctx context.Context, input model.BeeReq) (*model.FullBee, error) {
	auth, _ := r.Clients.Users.Auth(ctx, &users.Token{Val: input.Auth.Token})
	if !auth.Val {
		return &model.FullBee{
			Bee:     &model.Bee{},
			Details: &model.BeeDetails{},
		}, errors.New("Your are not autenticated")
	}
	res, err := r.Clients.Larsrv.GetBee(ctx, &laruche.BeeReq{
		BeeName: input.Author + "/" + input.Name,
	})
	if err != nil {
		return &model.FullBee{
			Bee:     &model.Bee{},
			Details: &model.BeeDetails{},
		}, err
	}
	return &model.FullBee{
		Bee:     model.ToBee(res),
		Details: &model.BeeDetails{},
	}, err
}

func (r *queryResolver) Register(ctx context.Context, input model.RegisterReq) (*model.RegisterRes, error) {
	log.Println("register: ", input.Username)
	if input.Password != input.Password2 {
		return &model.RegisterRes{
			Status:        false,
			StatusMessage: "passwords don't matchs",
		}, nil
	}
	_, err := r.Clients.Users.Register(ctx, &users.RegisterReq{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return &model.RegisterRes{
			Status:        false,
			StatusMessage: status.Convert(err).Message(),
		}, nil
	}
	return &model.RegisterRes{
		Status:        true,
		StatusMessage: "successfully registred",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

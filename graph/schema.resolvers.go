package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmagoga/new-equimper-go-graphql/graph/generated"
	"github.com/jmagoga/new-equimper-go-graphql/graph/model"
)

func (r *mutationResolver) CreateRequest(ctx context.Context, input model.NewRequest) (*model.Request, error) {
	// request := &model.Request{
	// 	ID:       fmt.Sprintf("T%d", rand.Int()),
	// 	Username: input.Username,
	// 	Email:    input.Email,
	// 	Phone:    input.Phone,
	// 	FullText: input.FullText,
	// }
	// r.requests = append(r.requests, request)
	// return request, nil

	// new, with db
	if len(input.Username) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.FullText) < 3 {
		return nil, errors.New("text not long enough")
	}

	request := &model.Request{
		// ID:       fmt.Sprintf("T%d", rand.Int()),
		Username: input.Username,
		Email:    input.Email,
		Phone:    input.Phone,
		FullText: input.FullText,
	}

	return r.RequestsRepo.CreateRequest(request)
}

func (r *mutationResolver) CreateBee(ctx context.Context, input model.NewBee) (*model.Bee, error) {
	bee := &model.Bee{
		Username: input.Username,
		Info:     input.Info,
		ImageURL: input.ImageURL,
		WikiURL:  input.WikiURL,
	}

	return r.BeesRepo.CreateBee(bee)
}

func (r *mutationResolver) UpdateBee(ctx context.Context, id string, input model.UpdateBee) (*model.Bee, error) {
	bee, err := r.BeesRepo.GetBeeById(id)
	if err != nil || bee == nil {
		return nil, errors.New("bee doesn't exist")
	}

	didUpdate := false

	if input.Username != nil {
		if len(*input.Username) < 3 {
			return nil, errors.New("Name not long enough.")
		}
		bee.Username = *input.Username
		didUpdate = true
	}

	if input.Info != nil {
		if len(*input.Info) < 3 {
			return nil, errors.New("Name not long enough.")
		}
		bee.Info = *input.Info
		didUpdate = true
	}

	if input.ImageURL != nil {
		if len(*input.ImageURL) < 3 {
			return nil, errors.New("Name not long enough.")
		}
		bee.ImageURL = *input.ImageURL
		didUpdate = true
	}

	if input.WikiURL != nil {
		if len(*input.WikiURL) < 3 {
			return nil, errors.New("Name not long enough.")
		}
		bee.WikiURL = *input.WikiURL
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done.")
	}

	bee, err = r.BeesRepo.UpdateBee(bee)
	if err != nil {
		return nil, fmt.Errorf("error while updating bee: %v", err)
	}

	return bee, nil
}

func (r *mutationResolver) DeleteBee(ctx context.Context, id string) (bool, error) {
	bee, err := r.BeesRepo.GetBeeById(id)
	if err != nil || bee == nil {
		return false, errors.New("bee does not exist")
	}

	err = r.BeesRepo.Delete(bee)
	if err != nil {
		return false, fmt.Errorf("Error while deleting bee %v", err)
	}

	return true, nil
}

func (r *queryResolver) Bees(ctx context.Context) ([]*model.Bee, error) {
	// return r.bees, nil
	return r.BeesRepo.GetBees()
}

func (r *queryResolver) Requests(ctx context.Context, filter *model.RequestFilter, limit *int, offset *int) ([]*model.Request, error) {
	// return r.requests, nil
	return r.RequestsRepo.GetRequests(filter, limit, offset)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

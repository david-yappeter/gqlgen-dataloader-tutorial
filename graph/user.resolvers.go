package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/dataloader"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *userResolver) Transactions(ctx context.Context, obj *model.User) ([]*model.Transaction, error) {
	// return service.TransactionGetByUserID(ctx, obj.ID)
	return dataloader.For(ctx).TransactionByID.Load(obj.ID)
}

func (r *userOpsResolver) Create(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.User, error) {
	return service.UserCreate(ctx, input)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }

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

func (r *transactionResolver) Summary(ctx context.Context, obj *model.Transaction) (*model.TransactionSummary, error) {
	// return service.TransactionDetailGetSummaryByTransactionID(ctx, obj.ID)
	resp, err := dataloader.For(ctx).TransactionDetailByID.Load(obj.ID)
	if err != nil {
		return nil, err
	}

	totalPrice, totalDisc := 0.0, 0.0
	for _, val := range resp {
		totalPrice += val.Price
		if val.Discount != nil {
			totalDisc += *val.Discount
		}
	}

	return &model.TransactionSummary{
		TotalPrice:         totalPrice,
		TotalDiscount:      &totalDisc,
		TransactionDetails: resp,
	}, nil
}

func (r *transactionOpsResolver) Create(ctx context.Context, obj *model.TransactionOps, input model.NewTransaction) (*model.Transaction, error) {
	return service.TransactionCreate(ctx, input)
}

// Transaction returns generated.TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

// TransactionOps returns generated.TransactionOpsResolver implementation.
func (r *Resolver) TransactionOps() generated.TransactionOpsResolver {
	return &transactionOpsResolver{r}
}

type transactionResolver struct{ *Resolver }
type transactionOpsResolver struct{ *Resolver }

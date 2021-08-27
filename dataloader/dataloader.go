package dataloader

import (
	"context"
	"myapp/graph/model"
	"myapp/service"
	"net/http"
	"time"
)

type loadersString string

const loadersKey = loadersString("dataloaders")

type Loaders struct {
	TransactionByID       TransactionByIDLoader
	TransactionDetailByID TransactionDetailByIDLoader
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			TransactionByID: TransactionByIDLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([][]*model.Transaction, []error) {

					resp, err := service.TransactionGetByUserIds(context.Background(), ids)
					if err != nil {
						return nil, []error{err}
					}

					itemById := map[int][]*model.Transaction{}
					for _, val := range resp {
						itemById[val.UserID] = append(itemById[val.UserID], val)
					}

					items := make([][]*model.Transaction, len(ids))
					for i, id := range ids {
						items[i] = itemById[id]
					}

					return items, nil
				},
			},
			TransactionDetailByID: TransactionDetailByIDLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([][]*model.TransactionDetail, []error) {

					resp, err := service.TransactionDetailGetByTransactionIds(context.Background(), ids)
					if err != nil {
						return nil, []error{err}
					}

					itemById := map[int][]*model.TransactionDetail{}
					for _, val := range resp {
						itemById[val.TransactionID] = append(itemById[val.TransactionID], val)
					}

					items := make([][]*model.TransactionDetail, len(ids))
					for i, id := range ids {
						items[i] = itemById[id]
					}

					return items, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

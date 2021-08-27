package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
)

func TransactionDetailGetAll(ctx context.Context) ([]*model.TransactionDetail, error) {
	db := config.GetDB()

	var transactionDetails []*model.TransactionDetail
	if err := db.Model(&model.TransactionDetail{}).Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return transactionDetails, nil
}

func TransactionDetailGetByTransactionID(ctx context.Context, transactionID int) ([]*model.TransactionDetail, error) {
	db := config.GetDB()

	var transactionDetails []*model.TransactionDetail
	if err := db.Model(&model.TransactionDetail{}).Where("transaction_id = ?", transactionID).Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return transactionDetails, nil
}

func TransactionDetailGetByTransactionIds(ctx context.Context, transactionIds []int) ([]*model.TransactionDetail, error) {
	db := config.GetDB()

	var transactionDetails []*model.TransactionDetail
	if err := db.Model(&model.TransactionDetail{}).Where("transaction_id IN (?)", transactionIds).Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return transactionDetails, nil
}

func TransactionDetailGetSummaryByTransactionID(ctx context.Context, transactionID int) (*model.TransactionSummary, error) {
	getTransactionDetails, err := TransactionDetailGetByTransactionID(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	totalPrice, totalDisc := 0.0, 0.0
	for _, val := range getTransactionDetails {
		totalPrice += val.Price
		if val.Discount != nil {
			totalDisc += *val.Discount
		}
	}

	return &model.TransactionSummary{
		TotalPrice:         totalPrice,
		TotalDiscount:      &totalDisc,
		TransactionDetails: getTransactionDetails,
	}, nil
}

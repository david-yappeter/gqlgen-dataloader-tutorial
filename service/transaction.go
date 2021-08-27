package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
	"time"
)

func TransactionCreate(ctx context.Context, input model.NewTransaction) (*model.Transaction, error) {
	db := config.GetDB()

	transaction := model.Transaction{
		CreatedAt: time.Now().UTC(),
		UserID:    input.UserID,
	}

	for _, val := range input.TransactionDetails {
		transaction.TransactionDetails = append(transaction.TransactionDetails, model.TransactionDetail{
			Name:        val.Name,
			Description: val.Description,
			Price:       val.Price,
			Discount:    val.Discount,
		})
	}

	if err := db.Model(transaction).Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func TransactionGetAll(ctx context.Context) ([]*model.Transaction, error) {
	db := config.GetDB()

	var transactions []*model.Transaction
	if err := db.Model(model.Transaction{}).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func TransactionGetByUserID(ctx context.Context, userID int) ([]*model.Transaction, error) {
	db := config.GetDB()

	var transactions []*model.Transaction
	if err := db.Model(model.Transaction{}).Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func TransactionGetByUserIds(ctx context.Context, userIds []int) ([]*model.Transaction, error) {
	db := config.GetDB()

	var transactions []*model.Transaction
	if err := db.Model(model.Transaction{}).Where("user_id IN (?)", userIds).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

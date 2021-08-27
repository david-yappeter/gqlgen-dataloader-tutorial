package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := config.GetDB()

	user := model.User{
		Name: input.Name,
		Age:  input.Age,
	}

	if err := db.Model(user).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserGetAll(ctx context.Context) ([]*model.User, error) {
	db := config.GetDB()

	var users []*model.User
	if err := db.Model(model.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

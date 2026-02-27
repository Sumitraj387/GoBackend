package repo

import (
	"context"
	"gobackend/internal/model"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user model.UserEntity) (model.UserEntity, error)
	GetUserById(ctx context.Context, id int64) (model.UserEntity, error)
}

func (repo *Repository) CreateUser(ctx context.Context, user model.UserEntity) (model.UserEntity, error) {
	err := repo.Db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.UserEntity{}, err
	}
	return user, nil
}
func (repo *Repository) GetUserById(ctx context.Context, id int64) (model.UserEntity, error) {
	var user model.UserEntity
	err := repo.Db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return model.UserEntity{}, err
	}
	return user, nil
}

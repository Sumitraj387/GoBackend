package core

import (
	"context"
	"gobackend/internal/model"
	"gobackend/internal/repo"

	"github.com/sirupsen/logrus"
)

type ICore interface {
	CreateUser(ctx context.Context, user model.User) (model.UserResponse, error)
	GetUserById(ctx context.Context, id int64) (model.UserResponse, error)
}
type Core struct {
	Logger *logrus.Entry
	RepoV1 repo.IUserRepo
}

func (core *Core) CreateUser(ctx context.Context, user model.User) (model.UserResponse, error) {
	UserResponse, err := core.RepoV1.CreateUser(ctx, model.UserEntity{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	})
	if err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		Id:       UserResponse.ID,
		Name:     UserResponse.Name,
		Phone:    UserResponse.Phone,
		Email:    UserResponse.Email,
		IsActive: UserResponse.IsActive,
	}, nil
}
func (core *Core) GetUserById(ctx context.Context, userId int64) (model.UserResponse, error) {
	user, err := core.RepoV1.GetUserById(ctx, userId)
	if err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		Id:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		IsActive: user.IsActive,
	}, nil
}

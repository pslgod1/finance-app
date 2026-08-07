package service

import (
	"FinanceApp/internal/model"
	"FinanceApp/internal/repository"
	"context"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(ctx context.Context, username, password string) (*model.User, error) {
	return u.repo.CreateUser(ctx, username, password)
}

func (u *UserService) GetUser(ctx context.Context, id int) (*model.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

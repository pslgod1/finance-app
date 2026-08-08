package service

import (
	"FinanceApp/internal/model"
	"FinanceApp/internal/repository"
	"context"
	"fmt"
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

func (u *UserService) DeleteUser(ctx context.Context, id int) error {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	if user.Role == "admin" {
		return fmt.Errorf("нельзя удалить администратора")
	}
	return u.repo.DeleteUser(ctx, id)
}

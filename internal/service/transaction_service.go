package service

import (
	"FinanceApp/internal/model"
	"FinanceApp/internal/repository"
	"context"
)

type TransactionService struct {
	repo *repository.Repository
}

func NewTransactionService(repo *repository.Repository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (t *TransactionService) CreateTransaction(
	ctx context.Context,
	userID int,
	transactionType string,
	amount float64,
	category string,
	description string,
) (*model.Transaction, error) {

	return t.repo.CreateTranscation(ctx, userID, transactionType, amount, category, description)
}

func (t *TransactionService) GetTransactions(ctx context.Context, userID int) ([]model.Transaction, error) {
	return t.repo.GetTransactionsByUserID(ctx, userID)
}

func (t *TransactionService) GetTransactionByID(ctx context.Context, id int) (*model.Transaction, error) {
	return t.repo.GetTransactionByID(ctx, id)
}

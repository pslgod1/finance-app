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

func (t *TransactionService) UpdateTransaction(
	ctx context.Context,
	id int,
	transactionType string,
	amount float64,
	category string,
	description string,
) (*model.Transaction, error) {

	_, err := t.GetTransactionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return t.repo.UpdateTransaction(ctx, id, transactionType, amount, category, description)
}

func (t *TransactionService) DeleteTransaction(ctx context.Context, id int) error {
	_, err := t.GetTransactionByID(ctx, id)
	if err != nil {
		return err
	}
	return t.repo.DeleteTransaction(ctx, id)
}

// STATISTICS!!!
func (t *TransactionService) GetStatistics(ctx context.Context, userID int) (*model.Statistics, error) {
	return t.repo.GetStatistics(ctx, userID)
}

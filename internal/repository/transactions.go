package repository

import (
	"FinanceApp/internal/model"
	"context"
)

func (r *Repository) CreateTranscation(
	ctx context.Context,
	userID int,
	transactionType string,
	amount float64,
	category string,
	description string,
) (*model.Transaction, error) {
	sqlQuery := `
	INSERT INTO transactions (user_id, type, amount, category, description)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, type, amount, category, description, created_at
`
	transaction := &model.Transaction{}
	err := r.Pool.QueryRow(ctx, sqlQuery, userID, transactionType, amount, category, description).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Category,
		&transaction.Description,
		&transaction.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *Repository) GetTransactionsByUserID(ctx context.Context, userID int) ([]model.Transaction, error) {
	sqlQuery := `
    SELECT id, user_id, type, amount, category, description, created_at
    FROM transactions
    WHERE user_id = $1
    ORDER BY created_at DESC
    `

	rows, err := r.Pool.Query(ctx, sqlQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		err := rows.Scan(&t.ID, &t.UserID, &t.Type, &t.Amount, &t.Category, &t.Description, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (r *Repository) GetTransactionByID(ctx context.Context, id int) (*model.Transaction, error) {
	sqlQuery := `
    SELECT id, user_id, type, amount, category, description, created_at
    FROM transactions
    WHERE id = $1
    `
	transaction := &model.Transaction{}
	err := r.Pool.QueryRow(ctx, sqlQuery, id).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Category,
		&transaction.Description,
		&transaction.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *Repository) UpdateTransaction(
	ctx context.Context,
	id int,
	transactionType string,
	amount float64,
	category string,
	description string,
) (*model.Transaction, error) {

	sqlQuery := `
	UPDATE transactions
	SET type = $1, amount = $2, category = $3, description = $4
	WHERE id = $5
	RETURNING id, user_id, type, amount, category, description, created_at
`
	transaction := &model.Transaction{}
	err := r.Pool.QueryRow(ctx, sqlQuery, transactionType, amount, category, description, id).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Category,
		&transaction.Description,
		&transaction.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return transaction, nil
}

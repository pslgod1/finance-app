package repository

import (
	"FinanceApp/internal/model"
	"context"
)

func (r *Repository) GetStatistics(ctx context.Context, userID int) (*model.Statistics, error) {
	sqlQuery := `
	SELECT 
    COALESCE(SUM(CASE WHEN type='income' THEN amount ELSE 0 END), 0),
    COALESCE(SUM(CASE WHEN type='expense' THEN amount ELSE 0 END), 0)
	FROM transactions WHERE user_id = $1
`
	statistics := &model.Statistics{}
	err := r.Pool.QueryRow(ctx, sqlQuery, userID).Scan(
		&statistics.TotalIncome,
		&statistics.TotalExpense,
	)
	if err != nil {
		return nil, err
	}

	statistics.Balance = statistics.TotalIncome - statistics.TotalExpense
	return statistics, nil

}

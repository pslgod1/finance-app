package repository

import (
	"FinanceApp/internal/model"
	"context"
)

func (r *Repository) CreateUser(
	ctx context.Context,
	username,
	password string,
) (*model.User, error) {
	sqlQuery := `
	INSERT INTO users (username, password, role)
	VALUES ($1, $2, 'user') 
	RETURNING id, username, password, role, created_at
`
	user := &model.User{}
	err := r.Pool.QueryRow(ctx, sqlQuery, username, password).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	sqlQuery := `
	SELECT id, username, password, role, created_at
	FROM users
	WHERE id = $1
`
	user := &model.User{}
	err := r.Pool.QueryRow(ctx, sqlQuery, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// dlya Admina
func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.Pool.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}

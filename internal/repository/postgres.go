package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool}
}

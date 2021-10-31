package repository

import (
	"context"
	"database/sql"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

// Repository interface
type Repository interface {
	GetAll(ctx context.Context) ([]Employee, error)
	AddEmployee(ctx context.Context, employees []*Employee) (int64, error)
	GetByID(ctx context.Context, id int64) (*Employee, error)
	Update(ctx context.Context, id int64, employee *Employee) (int64, error)
}

type repository struct {
	*sql.DB
}

// NewRepository returns repository
func NewRepository(db *sql.DB) Repository {
	return &repository{DB: db}
}

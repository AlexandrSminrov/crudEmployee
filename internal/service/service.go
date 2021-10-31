package service

import (
	"context"
	"crudEmployee/internal/repository"
	"regexp"
)

// regexp patterns
var (
	onlyRu    = regexp.MustCompile(`[^А-Яа-я]`)
	onlyRuEng = regexp.MustCompile(`[^A-Za-zА-Яа-я]`)
	onlyNum   = regexp.MustCompile(`[^0-9]`)
	address   = regexp.MustCompile(`[^a-zA-zа-яА-Я0-9,.\s№]`)
	email     = regexp.MustCompile(`[^а-яА-Я0-9,.\s№]`)
)

// Service interface
type Service interface {
	GetAll(ctx context.Context) ([]byte, error)
	AddEmployee(ctx context.Context, employees []*repository.Employee) (string, error)
	GetByID(ctx context.Context, id int64) ([]byte, error)
	Update(ctx context.Context, id int64, employee *repository.Employee) (string, error)
}

type service struct {
	repository.Repository
}

// NewService return service
func NewService(repository repository.Repository) Service {
	return &service{repository}
}

package service

import (
	"context"

	"github.com/cobnalt/Go/internal/database"
)

// Product struct
type Product struct {
	ID    int
	Name  string
	Slug  string
	Price float64
}

type Database interface {
	GetProducts(ctx context.Context, limit int, offset int) (result []database.Product, err error)
	GetProductByID(ctx context.Context, id int) (result database.Product, err error)
	CreateProduct(ctx context.Context, pr database.Product) (id int, err error)
}

type Service struct {
	db Database
}

func New(db Database) (*Service, error) {

	return &Service{
		db: db,
	}, nil
}

func (s *Service) GetProducts(ctx context.Context, limit, offset int) (result []Product, err error) {
	s.db.GetProducts(ctx, limit, offset)
}

package service

import (
	"context"

	"github.com/cobnalt/Go/internal/database"
)

// Product struct
type Product struct {
	ID             int
	ManufacturerID int
	CategoryID     int
	Name           string
	Slug           string
	Price          float64
	Description    string
}

//Database interface
type Database interface {
	GetProducts(ctx context.Context, limit int, offset int) (result []database.Product, err error)
	GetProductByID(ctx context.Context, id int) (result database.Product, err error)
	CreateProduct(ctx context.Context, pr database.Product) (id int, err error)
}

//Service struct
type Service struct {
	db Database
}

//New func
func New(db Database) (*Service, error) {

	return &Service{
		db: db,
	}, nil
}

//GetProducts func
func (s *Service) GetProducts(ctx context.Context, limit, offset int) (result []Product, err error) {
	res, err := s.db.GetProducts(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	result = make([]Product, len(res))
	for i, val := range res {
		result[i] = Product{
			ID:    val.ID,
			Name:  val.Name,
			Slug:  val.Slug,
			Price: val.Price,
		}
	}
	return result, nil
}

// GetProductByID func
func (s *Service) GetProductByID(ctx context.Context, id int) (result Product, err error) {
	dbProduct, err := s.db.GetProductByID(ctx, id)
	result = Product{
		ID:             dbProduct.ID,
		ManufacturerID: dbProduct.ManufacturerID,
		CategoryID:     dbProduct.CategoryID,
		Name:           dbProduct.Name,
		Slug:           dbProduct.Slug,
		Price:          dbProduct.Price,
		Description:    dbProduct.Description,
	}
	return result, nil
}

// CreateProduct func
func (s *Service) CreateProduct(ctx context.Context, pr database.Product) (id int, err error) {
	id, err = s.db.CreateProduct(ctx, pr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

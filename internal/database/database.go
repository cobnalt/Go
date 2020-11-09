package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// Product struct
type Product struct {
	ID    int
	Name  string
	Slug  string
	Price float64
}
type Database interface {
	GetProducts(ctx context.Context, limit int, offset int) (result []Product, err error)
	GetProductByID(ctx context.Context, id int) (result Product, err error)
	CreateProduct(ctx context.Context, pr Product) (id int, err error)
}

// DB connect
type DB struct {
	conn *sqlx.DB
}

// New DB
func New() (*DB, error) {
	conn, err := sqlx.Connect("postgres", "postgresql://postgres:postgres@localhost:5432/simple_catalog?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &DB{
		conn: conn,
	}, nil
}

// GetProducts func
func (d *DB) GetProducts(ctx context.Context, limit int, offset int) (result []Product, err error) {
	q := "SELECT id, name, slug, price FROM Product LIMIT ;"
	if err = d.conn.SelectContext(ctx, &result, q); err != nil {
		return nil, err
	}
	return result, err
}

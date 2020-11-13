package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
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

// Database interface
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
	q := fmt.Sprintf("SELECT id, manufacturer_id, category_id, name, slug, price, description FROM Product LIMIT %d OFFSET %d;", limit, offset)
	if err = d.conn.SelectContext(ctx, &result, q); err != nil {
		return nil, err
	}
	return result, err
}

// GetProductByID func
func (d *DB) GetProductByID(ctx context.Context, id int) (result Product, err error) {
	if err := d.conn.Get(&result, "SELECT id, manufacturer_id, category_id, name, slug, price, description FROM Product WHERE id=$1", id); err != nil {
		// проверка на sql ErrNoRows
		if err == sql.ErrNoRows {
			fmt.Errorf("No Rows Error")
		} else {
			return result, err
		}

	}
	return result, nil
}

// CreateProduct func
func (d *DB) CreateProduct(ctx context.Context, pr Product) (int, error) {
	q := "INSERT INTO Product (name, slug, price) VALUES ($1, $2, $3) RETURNING id;"
	var id int
	err := d.conn.GetContext(ctx, &id, q, pr.Name, pr.Slug, pr.Price)

	if err != nil {
		return 0, err
	}
	return id, nil
}

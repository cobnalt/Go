package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cobnalt/Go/internal/config"
	"github.com/jmoiron/sqlx"
)

// Product struct
type Product struct {
	ID             int     `db:"id"`
	ManufacturerID int     `db:"manufacturer_id"`
	CategoryID     int     `db:"category_id"`
	Name           string  `db:"name"`
	Slug           string  `db:"slug"`
	Price          float64 `db:"price"`
	Description    string  `db:"description"`
}

// Database interface
type Database interface {
	GetProducts(ctx context.Context, limit int, offset int) (result []Product, err error)
	GetProductByID(ctx context.Context, id int) (result Product, err error)
	CreateProduct(ctx context.Context, pr Product) (id int, err error)
}

var ErrNotFound = errors.New("Not found")

// DB connect
type DB struct {
	conn *sqlx.DB
}

// New DB
func New(cfg config.DatabaseConfig) (*DB, error) {
	conn, err := sqlx.Connect("postgres", cfg.ConnectionString)

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
	err = d.conn.GetContext(ctx, &result, "SELECT id, manufacturer_id, category_id, name, slug, price, description FROM Product WHERE id=$1", id)
	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Println("this err no rows ", err)
			return Product{}, ErrNotFound
		}
		return Product{}, err

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

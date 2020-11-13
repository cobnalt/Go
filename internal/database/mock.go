package database

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
)

//DatabaseMock struct
type DatabaseMock struct {
	mock.Mock
}

// GetProductByID func
func (m *DatabaseMock) GetProductByID(ctx context.Context, id int) (Product, error) {
	fmt.Println("Mocked func running")
	fmt.Printf("Value passed in: %d\n", id)

	args := m.Called(ctx, id)

	return args.Get(0).(Product), args.Error(1)
}

// GetProducts func
func (m *DatabaseMock) GetProducts(ctx context.Context, limit int, offset int) (result []Product, err error) {

	return nil, nil
}

// CreateProduct func
func (m *DatabaseMock) CreateProduct(ctx context.Context, pr Product) (int, error) {

	return 1, nil
}

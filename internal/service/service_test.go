package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/cobnalt/Go/internal/database"
)

func TestService_GetProductByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	myDatabase := new(database.DatabaseMock)
	myService, err := New(myDatabase)
	if err != nil {
		fmt.Errorf("Error")
		return
	}
	tests := []struct {
		name       string
		s          *Service
		args       args
		wantResult Product
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "Service_GetProductByID",
			s:    myService,
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantResult: Product{
				ID:             1,
				ManufacturerID: 1,
				CategoryID:     1,
				Name:           "dbProduct.Name",
				Slug:           "dbProduct.Slug",
				Price:          2.0,
				Description:    "dbProduct.Description",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myDatabase.On("GetProductByID", tt.args.ctx, tt.args.id).Return(database.Product{
				ID:             1,
				ManufacturerID: 1,
				CategoryID:     1,
				Name:           "dbProduct.Name",
				Slug:           "dbProduct.Slug",
				Price:          2.0,
				Description:    "dbProduct.Description",
			}, nil)
			gotResult, err := tt.s.GetProductByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Service.GetProductByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestService_CreateProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		pr  database.Product
	}
	dataBaseTest := new(database.DatabaseMock)
	myService, err := New(dataBaseTest)
	if err != nil {
		fmt.Errorf("Error")
		return
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantId  int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestService_CreateProduct",
			s:    myService,
			args: args{
				ctx: context.Background(),
				pr: database.Product{
					ID:             2,
					ManufacturerID: 1,
					CategoryID:     1,
					Name:           "dbProduct.Name",
					Slug:           "dbProduct.Slug",
					Price:          4.0,
					Description:    "dbProduct.Description",
				},
			},
			wantId:  2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataBaseTest.On("CreateProduct", tt.args.ctx, tt.args.pr).Return(2, nil)
			gotId, err := tt.s.CreateProduct(tt.args.ctx, tt.args.pr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Service.CreateProduct() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

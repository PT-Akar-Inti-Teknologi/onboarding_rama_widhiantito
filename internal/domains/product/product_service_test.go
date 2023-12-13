package product

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProductByID(productID uint) (*Product, error) {
	// This is where we capture the arguments passed to the GetProductByID method
	args := m.Called(productID)

	// args.Get(0) will fetch the first argument passed to the method
	// Here, we define what to return based on the provided productID
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Product), args.Error(1)
}

func (m *MockProductRepository) CreateProduct(*Product) error {
	return nil
}

func (m *MockProductRepository) UpdateProduct(dataProduct *Product) error {
	return nil
}

func Test_productService_GetProductByID(t *testing.T) {
	type args struct {
		productID uint
	}
	tests := []struct {
		name    string
		ps      *productService
		args    args
		want    *Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetProductByID(tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("productService.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productService.GetProductByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productService_UpdateProduct(t *testing.T) {
	type args struct {
		productID   uint
		dataProduct *Product
	}
	tests := []struct {
		name    string
		ps      *productService
		args    args
		want    *Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.UpdateProduct(tt.args.productID, tt.args.dataProduct)
			if (err != nil) != tt.wantErr {
				t.Errorf("productService.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productService.UpdateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

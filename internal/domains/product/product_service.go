package product

import (
	"errors"
	"fmt"
)

type ProductService interface {
	CreateProduct(newProduct *Product) (*Product, error)
	GetProductByID(productID uint) (*Product, error)
	UpdateProduct(productID uint, dataProduct *Product) (*Product, error)
	ValidateStock(producList []ProductList) ([]string, bool)
	UpdateStock(productID uint, quantity int) error
	// Add other methods for product management
}

type productService struct {
	productRepository ProductRepository
}

func NewProductService(productRepo ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (ps *productService) CreateProduct(newProduct *Product) (*Product, error) {
	if err := ps.productRepository.CreateProduct(newProduct); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return newProduct, nil
}

func (ps *productService) GetProductByID(productID uint) (*Product, error) {
	product, err := ps.productRepository.GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (ps *productService) UpdateProduct(productID uint, dataProduct *Product) (*Product, error) {

	existingProduct, err := ps.productRepository.GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	dataProduct.ID = existingProduct.ID
	// Update other fields as needed

	if err := ps.productRepository.UpdateProduct(dataProduct); err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	updatedProduct, err := ps.productRepository.GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	return updatedProduct, nil
}

func (ps *productService) ValidateStock(productList []ProductList) ([]string, bool) {
	var errorList []string

	for _, v := range productList {
		dataProduct, err := ps.productRepository.GetProductByID(v.ProductID)
		if err != nil {
			return nil, false
		}

		if dataProduct.Stock < v.Quantity {
			errMsg := fmt.Sprintf("Insufficient stock for product %d", v.ProductID)
			errorList = append(errorList, errMsg)
		}
	}

	// If there are errors, return the error list
	if len(errorList) > 0 {
		return errorList, false
	}

	// Return true indicating no errors if all products have sufficient stock
	return nil, true
}

func (ps *productService) UpdateStock(productID uint, quantity int) error {
	dataProduct, err := ps.productRepository.GetProductByID(productID)
	if err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	dataProduct.Stock -= quantity

	if err := ps.productRepository.UpdateProduct(dataProduct); err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil

}

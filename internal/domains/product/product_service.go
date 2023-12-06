package product

import (
	"errors"
	"fmt"
)

type ProductService interface {
	CreateProduct(name string, price int) (*Product, error)
	GetProductByID(productID uint) (*Product, error)
	UpdateProduct(productID uint, dataProduct *Product) (*Product, error)
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

func (ps *productService) CreateProduct(name string, price int) (*Product, error) {
	newProduct := &Product{
		Name:  name,
		Price: price,
		// Assign other fields as needed
	}

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

// Implement other ProductService methods for product management

package product

import "gorm.io/gorm"

type ProductRepository interface {
	CreateProduct(product *Product) error
	GetProductByID(productID uint) (*Product, error)
	UpdateProduct(product *Product) error
	// Add other methods for product management
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}

func (pr *productRepo) CreateProduct(product *Product) error {
	return pr.db.Create(product).Error
}

func (pr *productRepo) GetProductByID(productID uint) (*Product, error) {
	var product Product
	if err := pr.db.First(&product, productID).Error; err != nil {
		return nil, err // Handle not found error separately if needed
	}
	return &product, nil
}

func (pr *productRepo) UpdateProduct(product *Product) error {
	return pr.db.Save(product).Error
}

// Implement other ProductRepository methods

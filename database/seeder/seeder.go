package seeder

import (
	"order_transaction/internal/domains/customer"
	"order_transaction/internal/domains/product"

	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	productRepo := product.NewProductRepository(db)

	products := []product.Product{
		{Name: "Laptop A", Price: 9000000, Stock: 5},
		{Name: "Handphone A", Price: 7000000, Stock: 3},
	}

	for _, p := range products {
		if err := productRepo.CreateProduct(&p); err != nil {
			return err
		}
	}

	return nil
}

func SeedCustomers(db *gorm.DB) error {
	customerRepo := customer.NewCustomerRepository(db)

	customers := []customer.Customer{
		{Name: "John Doe"},
		{Name: "Jane Smith"},
		// Add more customers as needed
	}

	for _, c := range customers {
		if err := customerRepo.CreateCustomer(&c); err != nil {
			return err
		}
	}

	return nil
}

package customer

import (
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *Customer) error
	GetCustomerByEmail(email string) (*Customer, error)
	// Other methods for customer management
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db}
}

func (cr *customerRepo) CreateCustomer(customer *Customer) error {
	return cr.db.Create(customer).Error
}

func (cr *customerRepo) GetCustomerByEmail(email string) (*Customer, error) {
	var customer Customer
	if err := cr.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

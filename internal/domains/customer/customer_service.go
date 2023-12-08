package customer

import (
	"errors"
	"fmt"
)

type CustomerService interface {
	GetCustomerByEmail(email string) (*Customer, error)
}

type customerService struct {
	customerRepository CustomerRepository
}

func NewCustomerService(customerRepo CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepo,
	}
}

func (cs *customerService) GetCustomerByEmail(email string) (*Customer, error) {
	customer, err := cs.customerRepository.GetCustomerByEmail(email)

	if err != nil {
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

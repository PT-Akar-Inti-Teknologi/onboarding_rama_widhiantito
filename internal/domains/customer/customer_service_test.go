package customer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) CreateCustomer(customer *Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) GetCustomerByEmail(email string) (*Customer, error) {
	args := m.Called(email)
	return args.Get(0).(*Customer), args.Error(1)
}

func Test_customerService_GetCustomerByEmail(t *testing.T) {
	mockRepo := new(MockCustomerRepository)

	cs := &customerService{customerRepository: mockRepo}

	tests := []struct {
		name    string
		args    string
		mockFn  func()
		want    *Customer
		wantErr bool
	}{
		{
			name: "Customer found",
			args: "test@example.com",
			mockFn: func() {
				mockRepo.On("GetCustomerByEmail", "test@example.com").Return(&Customer{Name: "Test Customer", Email: "test@example.com"}, nil)
			},
			want:    &Customer{Name: "Test Customer", Email: "test@example.com"},
			wantErr: false,
		},
		{
			name: "Customer not found",
			args: "nothing@example.com",
			mockFn: func() {
				mockRepo.On("GetCustomerByEmail", "nothing@example.com").Return((*Customer)(nil), errors.New("customer not found"))
			},
			want:    nil,
			wantErr: true,
		},
		// Add more test cases if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock behavior for the current test case
			tt.mockFn()

			// Call the method being tested
			got, err := cs.GetCustomerByEmail(tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("customerService.GetCustomerByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customerService.GetCustomerByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestNewCustomerService(t *testing.T) {
	mockRepo := new(MockCustomerRepository)

	tests := []struct {
		name string
		args struct {
			customerRepo CustomerRepository
		}
		want CustomerService
	}{
		{
			name: "New CustomerService creation",
			args: struct{ customerRepo CustomerRepository }{customerRepo: mockRepo},
			want: &customerService{customerRepository: mockRepo},
		},
		// Add more test cases if needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerService(tt.args.customerRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

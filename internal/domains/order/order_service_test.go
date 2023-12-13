package order

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockOrderRepository struct {
	mock.Mock
}

type MockOrderItemRepository struct {
	mock.Mock
}

// CreateOrder mocks the creation of an order
func (m *MockOrderRepository) CreateOrder(newOrder *Order) error {
	args := m.Called(newOrder)
	return args.Error(0)
}

func (m *MockOrderRepository) DeleteOrder(orderID uint) error {
	args := m.Called(orderID)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrderByID(orderID uint) (*Order, error) {
	args := m.Called(orderID)
	return args.Get(0).(*Order), args.Error(1)
}

func (m *MockOrderRepository) UpdateOrder(dataOrder *Order) error {
	args := m.Called(dataOrder)
	return args.Error(0)
}

func (m *MockOrderItemRepository) CreateOrderItem(newOrderItem *OrderItem) error {
	args := m.Called(newOrderItem)
	return args.Error(0)
}

func (m *MockOrderItemRepository) DeleteOrderItem(orderItemID uint) error {
	args := m.Called(orderItemID)
	return args.Error(0)
}

func (m *MockOrderItemRepository) GetOrderItemByID(orderItemID uint) (*OrderItem, error) {
	args := m.Called(orderItemID)
	return args.Get(0).(*OrderItem), args.Error(1)
}

func (m *MockOrderItemRepository) UpdateOrderItem(dataOrderItem *OrderItem) error {
	args := m.Called(dataOrderItem)
	return args.Error(0)
}

func Test_orderService_CreateOrder(t *testing.T) {
	// Mock Order Repository
	mockRepo := new(MockOrderRepository)

	// Creating Order Service using Mock Repository
	os := &orderService{orderRepository: mockRepo}

	// Define test cases
	tests := []struct {
		name    string
		args    *Order
		want    *Order
		wantErr bool
	}{
		{
			name:    "Successful order creation",
			args:    &Order{CustomerID: 1, Status: "pending"}, // Provide the necessary order details for the test case
			want:    &Order{CustomerID: 1, Status: "pending"}, // Define the expected result
			wantErr: false,
		},
		{
			name:    "Unsuccessful order creation",
			args:    &Order{CustomerID: 2, Status: "pending"}, // Provide the necessary order details for the test case
			want:    nil,                                      // Define the expected result
			wantErr: true,
		},
	}

	// Mock the behavior of CreateOrder method in the mock repository
	mockRepo.On("CreateOrder", mock.Anything).Return(nil).Once()                              // Adjust based on your specific case
	mockRepo.On("CreateOrder", mock.Anything).Return(errors.New("error create order")).Once() // Adjust based on your specific case

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := os.CreateOrder(tt.args)

			// Check for error condition
			if (err != nil) != tt.wantErr {
				t.Errorf("orderService.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check the returned order
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderService.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}

	// Assert that the expected methods were called
	mockRepo.AssertExpectations(t)
}

func Test_orderService_GetOrderByID(t *testing.T) {
	// Mock Order Repository
	mockRepo := new(MockOrderRepository)

	// Creating Order Service using Mock Repository
	os := &orderService{orderRepository: mockRepo}

	// Define test cases
	tests := []struct {
		name    string
		args    uint
		mockFn  func()
		want    *Order
		wantErr bool
	}{
		{
			name: "Valid order ID",
			args: 1, // Provide the order ID
			mockFn: func() {
				mockRepo.On("GetOrderByID", uint(1)).Return(&Order{}, nil)
			},
			want:    &Order{}, // Define the expected order for this case
			wantErr: false,
		},
		{
			name: "Order not found",
			args: 2, // Provide a non-existing order ID
			mockFn: func() {
				mockRepo.On("GetOrderByID", uint(2)).Return((*Order)(nil), errors.New("order not found"))
			},
			want:    nil,
			wantErr: true,
		},
		// Add more test cases for different scenarios if needed
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock behavior for the current test case
			tt.mockFn()

			// Call the method being tested
			got, err := os.GetOrderByID(tt.args)

			// Check for error condition
			if (err != nil) != tt.wantErr {
				t.Errorf("orderService.GetOrderByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check the returned order
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderService.GetOrderByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderService_UpdateOrder(t *testing.T) {
	// Mock Order Repository
	mockRepo := new(MockOrderRepository)

	// Creating Order Service using Mock Repository
	os := &orderService{orderRepository: mockRepo}

	type args struct {
		orderID uint
		order   *Order
	}

	tests := []struct {
		name    string
		os      *orderService
		args    args
		mockFn  func()
		want    *Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Successful Update order",
			args: args{orderID: 1, order: &Order{Model: gorm.Model{ID: 1}, Status: "completed"}},
			mockFn: func() {
				// Mock GetOrderByID method to return an Order for the specified order ID
				mockRepo.On("GetOrderByID", uint(1)).Return(&Order{Model: gorm.Model{ID: 1}, Status: "pending"}, nil)

				// Mock UpdateOrder method to return success after updating the order status
				mockRepo.On("UpdateOrder", &Order{Model: gorm.Model{ID: 1}, Status: "completed"}).Return(nil)
			},
			want:    &Order{Model: gorm.Model{ID: 1}, Status: "completed"},
			wantErr: false,
		},
		{
			name: "Unsuccessful Update order - order not found",
			args: args{orderID: 2, order: &Order{Model: gorm.Model{ID: 2}, Status: "completed"}},
			mockFn: func() {
				// Mock GetOrderByID method to return an Order for the specified order ID
				mockRepo.On("GetOrderByID", uint(2)).Return(&Order{}, errors.New("Order not found"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed to update order",
			args: args{orderID: 3, order: &Order{Model: gorm.Model{ID: 3}, Status: "completed"}},
			mockFn: func() {
				// Mock GetOrderByID method to return an Order for the specified order ID
				mockRepo.On("GetOrderByID", uint(3)).Return(&Order{Model: gorm.Model{ID: 3}, Status: "pending"}, nil)

				// Mock UpdateOrder method to return success after updating the order status
				mockRepo.On("UpdateOrder", &Order{Model: gorm.Model{ID: 3}, Status: "completed"}).Return(errors.New("Failed to update order"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			// Call the method being tested using the 'os' instance
			got, err := os.UpdateOrder(tt.args.orderID, tt.args.order)

			if (err != nil) != tt.wantErr {
				t.Errorf("orderService.UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderService.UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderService_DeleteOrder(t *testing.T) {
	// Mock Order Repository
	mockRepo := new(MockOrderRepository)

	// Creating Order Service using Mock Repository
	os := &orderService{orderRepository: mockRepo}

	type args struct {
		orderID uint
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func()
		wantErr bool
	}{
		{
			name: "Successful deletion",
			args: args{orderID: 1}, // Provide the order ID you want to delete
			mockFn: func() { // Mock the DeleteOrder method to return success
				mockRepo.On("DeleteOrder", uint(1)).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Failed deletion",
			args: args{orderID: 2}, // Provide a different order ID for failed deletion
			mockFn: func() { // Mock the DeleteOrder method to return an error
				mockRepo.On("DeleteOrder", uint(2)).Return(errors.New("deletion failed"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock behavior for the current test case
			tt.mockFn()

			// Call the method being tested
			err := os.DeleteOrder(tt.args.orderID)

			// Check for error condition
			if (err != nil) != tt.wantErr {
				t.Errorf("orderService.DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewOrderService(t *testing.T) {
	// Mock Order Repository
	mockRepo := new(MockOrderRepository)

	type args struct {
		orderRepo OrderRepository
	}
	tests := []struct {
		name string
		args args
		want OrderService
	}{
		{
			name: "New OrderService creation",
			args: args{orderRepo: mockRepo},
			want: &orderService{orderRepository: mockRepo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderService(tt.args.orderRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrderItemService(t *testing.T) {
	mockRepo := new(MockOrderItemRepository)

	type args struct {
		orderItemRepo OrderItemRepository
	}
	tests := []struct {
		name string
		args args
		want OrderItemService
	}{
		{
			name: "New OrderItemService creation",
			args: args{orderItemRepo: mockRepo},
			want: &orderItemService{orderItemRepository: mockRepo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderItemService(tt.args.orderItemRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

package order

import (
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *MockDB {
	m.Called(value)
	return m
}

func (m *MockDB) Error() error {
	args := m.Called()
	return args.Error(0)
}

// func Test_orderRepo_CreateOrder(t *testing.T) {
// 	mockRepo := &MockOrderRepository{} // Initialize the mock repository

// 	orderRepo := &orderRepo{db: &gorm.DB{}} // Use the mock repository in the OrderRepo

// 	type args struct {
// 		order *Order
// 	}
// 	tests := []struct {
// 		name    string
// 		or      *orderRepo
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "Successful order creation",
// 			or:   orderRepo,
// 			args: args{
// 				order: &Order{
// 					// Define your order details here
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		// Add more test cases for different scenarios if needed
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Mock the behavior of CreateOrder method in the mock repository
// 			mockRepo.On("CreateOrder", mock.Anything).Return(nil).Once()

// 			if err := tt.or.CreateOrder(tt.args.order); (err != nil) != tt.wantErr {
// 				t.Errorf("orderRepo.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

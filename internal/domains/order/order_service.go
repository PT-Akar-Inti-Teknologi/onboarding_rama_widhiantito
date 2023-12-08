package order

import (
	"errors"
	"fmt"
)

type OrderService interface {
	CreateOrder(newOrder *Order) (*Order, error)
	GetOrderByID(orderID uint) (*Order, error)
	UpdateOrder(orderID uint, status string) (*Order, error)
	DeleteOrder(orderID uint) error
	// Other methods as required
}

type OrderItemService interface {
	CreateOrderItem(orderID, productID uint, quantity int) (*OrderItem, error)
	GetOrderItemByID(orderItemID uint) (*OrderItem, error)
	UpdateOrderItem(orderItemID, quantity int) (*OrderItem, error)
	DeleteOrderItem(orderItemID uint) error
	// Other methods as required
}

type orderService struct {
	orderRepository     OrderRepository
	orderItemRepository OrderItemRepository
}

type orderItemService struct {
	orderItemRepository OrderItemRepository
}

func NewOrderService(orderRepo OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepo,
	}
}

func (os *orderService) CreateOrder(newOrder *Order) (*Order, error) {
	if err := os.orderRepository.CreateOrder(newOrder); err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return newOrder, nil
}

func (os *orderService) GetOrderByID(orderID uint) (*Order, error) {
	order, err := os.orderRepository.GetOrderByID(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (os *orderService) UpdateOrder(orderID uint, status string) (*Order, error) {
	order, err := os.orderRepository.GetOrderByID(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	if order == nil {
		return nil, errors.New("order not found")
	}

	order.Status = status

	if err := os.orderRepository.UpdateOrder(order); err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	return order, nil
}

func (os *orderService) DeleteOrder(orderID uint) error {
	return os.orderRepository.DeleteOrder(orderID)
}

func NewOrderItemService(orderItemRepo OrderItemRepository) OrderItemService {
	return &orderItemService{
		orderItemRepository: orderItemRepo,
	}
}

func (ois *orderItemService) CreateOrderItem(orderID, productID uint, quantity int) (*OrderItem, error) {
	newOrderItem := &OrderItem{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		// Add other fields as needed
	}

	if err := ois.orderItemRepository.CreateOrderItem(newOrderItem); err != nil {
		return nil, fmt.Errorf("failed to create order item: %w", err)
	}

	return newOrderItem, nil
}

func (ois *orderItemService) GetOrderItemByID(orderItemID uint) (*OrderItem, error) {
	orderItem, err := ois.orderItemRepository.GetOrderItemByID(orderItemID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}
	if orderItem == nil {
		return nil, errors.New("order item not found")
	}
	return orderItem, nil
}

func (ois *orderItemService) UpdateOrderItem(orderItemID, quantity int) (*OrderItem, error) {
	orderItem, err := ois.orderItemRepository.GetOrderItemByID(uint(orderItemID))
	if err != nil {
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}
	if orderItem == nil {
		return nil, errors.New("order item not found")
	}

	orderItem.Quantity = quantity
	// Update other fields if needed

	if err := ois.orderItemRepository.UpdateOrderItem(orderItem); err != nil {
		return nil, fmt.Errorf("failed to update order item: %w", err)
	}

	return orderItem, nil
}

func (ois *orderItemService) DeleteOrderItem(orderItemID uint) error {
	return ois.orderItemRepository.DeleteOrderItem(orderItemID)
}

// Implement other methods as required

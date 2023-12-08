package order

import "gorm.io/gorm"

type OrderRepository interface {
	CreateOrder(order *Order) error
	GetOrderByID(orderID uint) (*Order, error)
	UpdateOrder(order *Order) error
	DeleteOrder(orderID uint) error
	// Other methods as required
}

type OrderItemRepository interface {
	CreateOrderItem(orderItem *OrderItem) error
	GetOrderItemByID(orderItemID uint) (*OrderItem, error)
	UpdateOrderItem(orderItem *OrderItem) error
	DeleteOrderItem(orderItemID uint) error
	// Other methods as required
}

type orderRepo struct {
	db *gorm.DB
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db}
}

func (or *orderRepo) CreateOrder(order *Order) error {
	return or.db.Create(order).Error
}

func (or *orderRepo) GetOrderByID(orderID uint) (*Order, error) {
	var order Order
	if err := or.db.Preload("OrderItems").First(&order, orderID).Error; err != nil {
		return nil, err // Handle not found error separately if needed
	}
	return &order, nil
}

func (or *orderRepo) UpdateOrder(order *Order) error {
	return or.db.Save(order).Error
}

func (or *orderRepo) DeleteOrder(orderID uint) error {
	return or.db.Delete(&Order{}, orderID).Error
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db}
}

func (oir *orderItemRepository) CreateOrderItem(orderItem *OrderItem) error {
	return oir.db.Create(orderItem).Error
}

func (oir *orderItemRepository) GetOrderItemByID(orderItemID uint) (*OrderItem, error) {
	var orderItem OrderItem
	if err := oir.db.First(&orderItem, orderItemID).Error; err != nil {
		return nil, err // Handle not found error separately if needed
	}
	return &orderItem, nil
}

func (oir *orderItemRepository) UpdateOrderItem(orderItem *OrderItem) error {
	return oir.db.Save(orderItem).Error
}

func (oir *orderItemRepository) DeleteOrderItem(orderItemID uint) error {
	return oir.db.Delete(&OrderItem{}, orderItemID).Error
}

// Implement other methods as required

package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID      uint        `json:"id_customer"`
	OrderDate       time.Time   `json:"order_date"`
	Status          string      `json:"status"`
	TransferProofID uint        `json:"id_file"`
	OrderItems      []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint `json:"id_order"`
	ProductID uint `json:"id_product"`
	Quantity  int  `json:"quantity"`
}

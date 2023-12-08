package handlers

import (
	"net/http"
	"order_transaction/internal/domains/customer"
	"order_transaction/internal/domains/order"
	"order_transaction/internal/domains/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService     order.OrderService
	OrderItemService order.OrderItemService
	ProductService   product.ProductService
	CustomerService  customer.CustomerService
}

func (oh *OrderHandler) CreateOrder(c *gin.Context) {

	var bodyRequest struct {
		Customer    customer.Customer     `json:"customer"`
		ProductList []product.ProductList `json:"product_list"`
	}

	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataCustomer, err := oh.CustomerService.GetCustomerByEmail(bodyRequest.Customer.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	errList, isValid := oh.ProductService.ValidateStock(bodyRequest.ProductList)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errList})
		return
	}

	dataOrder := new(order.Order)
	dataOrder.CustomerID = dataCustomer.ID
	dataOrder.Status = "pending"

	newOrder, err := oh.OrderService.CreateOrder(dataOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, v := range bodyRequest.ProductList {
		if _, err := oh.OrderItemService.CreateOrderItem(newOrder.ID, v.ProductID, v.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		oh.ProductService.UpdateStock(v.ProductID, v.Quantity)
	}

	c.JSON(http.StatusCreated, newOrder)
}

func (oh *OrderHandler) GetOrderByID(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := oh.OrderService.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Implement handlers for UpdateOrder, DeleteOrder, and other functionalities

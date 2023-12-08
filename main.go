package main

import (
	"order_transaction/database"
	"order_transaction/internal/domains/customer"
	"order_transaction/internal/domains/order"
	"order_transaction/internal/domains/product"
	"order_transaction/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		// Handle the error
		panic(err)
	}

	//Run AutoMigration
	db.AutoMigrate(product.Product{})
	db.AutoMigrate(customer.Customer{})
	db.AutoMigrate(order.Order{})
	db.AutoMigrate(order.OrderItem{})

	// if err := seeder.SeedProducts(db); err != nil {
	// 	// Handle error
	// 	return
	// }

	// if err := seeder.SeedCustomers(db); err != nil {
	// 	//handle error
	// 	return
	// }

	// Initialize repositories
	productRepo := product.NewProductRepository(db)
	orderRepo := order.NewOrderRepository(db)
	orderItemRepo := order.NewOrderItemRepository(db)
	customerRepo := customer.NewCustomerRepository(db)

	// Initialize services
	productService := product.NewProductService(productRepo)
	orderService := order.NewOrderService(orderRepo)
	orderItemService := order.NewOrderItemService(orderItemRepo)
	customerService := customer.NewCustomerService(customerRepo)

	// Create Gin router
	router := gin.Default()

	// Initialize handlers
	productHandler := &handlers.ProductHandler{
		ProductService: productService,
	}

	orderHandler := &handlers.OrderHandler{
		OrderService:     orderService,
		OrderItemService: orderItemService,
		ProductService:   productService,
		CustomerService:  customerService,
	}

	// Define routes
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.GetProductByID)
		productRoutes.PUT("update/:id", productHandler.UpdateProduct)
		// Define other product routes
	}

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", orderHandler.CreateOrder)
	}

	// Start the server
	router.Run(":8080")
}

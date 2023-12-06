package main

import (
	"order_transaction/database"
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

	// Initialize repositories
	//userRepo := user.NewUserRepository(db)
	productRepo := product.NewProductRepository(db)

	// Initialize services
	//userService := user.NewUserService(userRepo)
	productService := product.NewProductService(productRepo)

	// Create Gin router
	router := gin.Default()

	// Initialize handlers
	// userHandler := &handlers.UserHandler{
	// 	UserService: userService,
	// }

	productHandler := &handlers.ProductHandler{
		ProductService: productService,
	}

	// Define routes
	// userRoutes := router.Group("/users")
	// {
	// 	userRoutes.POST("/", userHandler.CreateUser)
	// 	userRoutes.GET("/:id", userHandler.GetUserByID)
	// 	// Define other user routes
	// }

	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.GetProductByID)
		productRoutes.PUT("update/:id", productHandler.UpdateProduct)
		// Define other product routes
	}

	// Start the server
	router.Run(":8080")
}

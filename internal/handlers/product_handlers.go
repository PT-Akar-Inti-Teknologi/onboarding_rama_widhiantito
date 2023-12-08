package handlers

import (
	"fmt"
	"net/http"
	"order_transaction/internal/domains/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService product.ProductService
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	newProduct := new(product.Product)

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct, err := ph.ProductService.CreateProduct(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func (ph *ProductHandler) GetProductByID(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := ph.ProductService.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	updateProduct := new(product.Product)
	if err := c.ShouldBindJSON(updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	updateProduct.ID = uint(productID)
	fmt.Println(updateProduct.Name)
	updatedProduct, err := ph.ProductService.UpdateProduct(uint(productID), updateProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// Implement other product-related handlers

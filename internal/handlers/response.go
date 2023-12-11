package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	respCodeOrder200 = "ONBOARDING-ORDER-200"

	respCodeProduct200 = "ONBOARDING-PRODUCT-200"
)

// ResponseSchema represents the schema of the response
type ResponseSchema struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

// CreateSuccessResponse generates a success response with detail
func CreateDetailResponse(c *gin.Context, responseCode string, responseData interface{}) {
	response := gin.H{
		"response_schema": ResponseSchema{
			ResponseCode:    responseCode,
			ResponseMessage: "Sukses",
		},
		"response_output": gin.H{
			"detail": responseData,
		},
	}
	c.JSON(http.StatusOK, response)
}

// CreateListResponse generates a list response with pagination and content
func CreateListResponse(c *gin.Context, responseCode string, responseData interface{}, pagination interface{}) {
	response := gin.H{
		"response_schema": ResponseSchema{
			ResponseCode:    responseCode,
			ResponseMessage: "Sukses",
		},
		"response_output": gin.H{
			"list": gin.H{
				"pagination": pagination,
				"content":    responseData,
			},
		},
	}
	c.JSON(http.StatusOK, response)
}

// CreateErrorResponse generates an error response with errors array
func CreateErrorResponse(c *gin.Context, responseCode string, responseData interface{}) {
	response := gin.H{
		"response_schema": ResponseSchema{
			ResponseCode:    responseCode,
			ResponseMessage: "Parameter tidak valid",
		},
		"response_output": gin.H{
			"errors": responseData,
		},
	}
	c.JSON(http.StatusBadRequest, response)
}

func CreateInvalidResponse(c *gin.Context, responseCode, responseMessage string, responseData interface{}) {
	response := gin.H{
		"response_schema": ResponseSchema{
			ResponseCode:    responseCode,
			ResponseMessage: responseMessage,
		},
		"response_output": gin.H{},
	}
	c.JSON(http.StatusBadRequest, response)
}

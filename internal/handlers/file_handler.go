package handlers

import (
	"fmt"
	"net/http"
	files "order_transaction/internal/domains/file"
	"order_transaction/internal/domains/order"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	FileService  files.FileService
	OrderService order.OrderService
}

func (fh *FileHandler) HandleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID, err := strconv.Atoi(c.PostForm("order_id"))
	fmt.Println(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataOrder, err := fh.OrderService.GetOrderByID(uint(orderID))
	if err != nil {
		CreateInvalidResponse(c, respCodeOrder200, err.Error(), nil)
	}

	// Save the file to a specific folder on the server
	filePath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	newFile := files.File{
		Name:       file.Filename,
		MimeType:   file.Header.Get("Content-Type"),
		FilePath:   filePath,
		IsUploaded: true,
	}

	data, err := fh.FileService.Create(&newFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file to db"})
		return
	}

	dataOrder.TransferProofID = data.ID
	dataOrder, _ = fh.OrderService.UpdateOrder(dataOrder.ID, "in_review")

	dataResp := gin.H{
		"file":  data,
		"order": dataOrder,
	}

	CreateDetailResponse(c, respCodeOrder200, dataResp)
}

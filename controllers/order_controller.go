package controllers

import (
	"food-delivery/config"
	"food-delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderQueue - Channel for processing orders asynchronously
var OrderQueue = make(chan uint, 10) // Buffered channel with capacity 10

// PLACE ORDER (Customer) - Send to worker queue
func PlaceOrder(c *gin.Context) {
	var order models.Order

	// Bind JSON request to order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.Status = "pending"

	// Save order to database
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error placing order"})
		return
	}

	// Send order to worker queue for async processing
	go func(orderID uint) {
		OrderQueue <- orderID
	}(order.ID)

	c.JSON(http.StatusCreated, gin.H{"message": "Order placed successfully. Processing started.", "order": order})
}

// GET ALL ORDERS
func GetOrders(c *gin.Context) {
	var orders []models.Order

	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GET ORDER BY ID
func GetOrderByID(c *gin.Context) {
	var order models.Order
	orderID := c.Param("id")

	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UPDATE ORDER STATUS
func UpdateOrder(c *gin.Context) {
	var order models.Order
	orderID := c.Param("id")

	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": order})
}

// DELETE ORDER
func DeleteOrder(c *gin.Context) {
	var order models.Order
	orderID := c.Param("id")

	if err := config.DB.Delete(&order, orderID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

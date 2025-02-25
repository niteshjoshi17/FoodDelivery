package controllers

import (
	"food-delivery/config"
	"food-delivery/models"
	"log"
)

// StartOrderWorkerPool initializes worker goroutines to process orders asynchronously.
func StartOrderWorkerPool(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go orderWorker(i)
	}
}

// orderWorker processes incoming orders from the queue
func orderWorker(workerID int) {
	for orderID := range OrderQueue {
		log.Printf("Worker %d: Processing order ID %d\n", workerID, orderID)
		processOrder(orderID)
	}
}

// Process order
func processOrder(orderID uint) {
	// Directly update order status in the database
	result := config.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", "completed")

	if result.Error != nil {
		log.Printf("Error processing order ID %d: %v\n", orderID, result.Error)
		return
	}

	log.Printf("Order ID %d has been processed and marked as completed.\n", orderID)
}

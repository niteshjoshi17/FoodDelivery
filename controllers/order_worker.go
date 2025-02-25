package controllers

import (
	"food-delivery/config"
	"food-delivery/models"
	"log"
	"time"
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

// processOrder simulates order processing
func processOrder(orderID uint) {
	time.Sleep(5 * time.Second) // Simulate processing delay

	// Update order status in the database
	config.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", "completed")

	log.Printf("Order ID %d has been processed and marked as completed.\n", orderID)
}

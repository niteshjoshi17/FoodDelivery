package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Order structure
type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Status    string `json:"status"`
	Timestamp time.Time
}

var orderQueue = make(chan Order, 10) // Buffered Channel for concurrent processing
var orderID int = 1
var mu sync.Mutex // Mutex to avoid race conditions

// Worker Pool to process orders concurrently
func StartOrderWorkerPool(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			for order := range orderQueue {
				fmt.Printf("Worker %d processing Order ID: %d\n", workerID, order.ID)
				time.Sleep(2 * time.Second) // Simulate processing time
				fmt.Printf("Worker %d completed Order ID: %d\n", workerID, order.ID)
			}
		}(i + 1)
	}
}

// PlaceOrder - API to place an order (handled concurrently)
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	newOrder := Order{
		ID:        orderID,
		UserID:    orderID, // Simulated User ID
		Status:    "Pending",
		Timestamp: time.Now(),
	}
	orderID++
	mu.Unlock()

	orderQueue <- newOrder // Send order to the queue for processing

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Order %d placed successfully!", newOrder.ID),
	})
}

// AcceptOrder - API to accept an order
func AcceptOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order accepted!"))
}

// MarkDelivered - API to mark an order as delivered
func MarkDelivered(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order marked as delivered!"))
}

// GetOrderHistory - API to fetch order history
func GetOrderHistory(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Fetched order details")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Fetched payment details")
	}()

	wg.Wait() // Wait for both goroutines to complete

	w.Write([]byte("User order history!"))
}

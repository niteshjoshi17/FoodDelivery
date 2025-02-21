package main

import (
	"log"
	"net/http"

	"food-delivery/config"
	"food-delivery/controllers"
	"food-delivery/routes"
)

func main() {
	// Initialize database
	config.InitDB()

	// Start Order Processing Worker Pool
	go controllers.StartOrderWorkerPool(3) // 3 concurrent workers

	r := routes.RegisterRoutes()

	log.Println(" Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

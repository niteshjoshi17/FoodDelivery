package main

import (
	"log"
	"net/http"

	"food-delivery/config"
	"food-delivery/controllers"
	"food-delivery/models"
	"food-delivery/routes"
)

func main() {
	// Initialize database
	config.InitDB()

	config.DB.AutoMigrate(&models.Restaurant{}, &models.Order{})

	// ✅ Start Order Processing Worker Pool
	go controllers.StartOrderWorkerPool(3) // 3 concurrent workers

	r := routes.RegisterRoutes()

	log.Println("Server running on port 9090")
	http.ListenAndServe(":9090", r)
}

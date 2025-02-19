package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Models
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type"` // customer / delivery_agent
}

type Restaurant struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float32 `json:"rating"`
}

type Order struct {
	ID           uint    `gorm:"primaryKey"`
	UserID       uint    `json:"user_id"`
	RestaurantID uint    `json:"restaurant_id"`
	AgentID      *uint   `json:"agent_id"`
	Status       string  `json:"status"` // pending, accepted, in_transit, delivered
	TotalPrice   float64 `json:"total_price"`
}

func initDB() {

	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connected!")

	db.AutoMigrate(&User{}, &Restaurant{}, &Order{})
	fmt.Println("Database migrated!")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User registered!"))
}

func placeOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order placed!"))
}

func acceptOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order accepted!"))
}

func markDelivered(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order marked as delivered!"))
}

func getOrderHistory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User order history!"))
}

func getRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Restaurant details!"))
}

func main() {
	initDB()
	r := mux.NewRouter()

	r.HandleFunc("/register", registerUser).Methods("POST")
	r.HandleFunc("/place-order", placeOrder).Methods("POST")
	r.HandleFunc("/accept-order", acceptOrder).Methods("POST")
	r.HandleFunc("/mark-delivered", markDelivered).Methods("POST")
	r.HandleFunc("/order-history/{user_id}", getOrderHistory).Methods("GET")
	r.HandleFunc("/restaurant/{id}", getRestaurant).Methods("GET")

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

package routes

import (
	"fmt"
	"net/http"

	"food-delivery/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Default home route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Food Delivery API!")
	}).Methods("GET")

	// API routes
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/place-order", controllers.PlaceOrder).Methods("POST")
	r.HandleFunc("/accept-order", controllers.AcceptOrder).Methods("POST")
	r.HandleFunc("/mark-delivered", controllers.MarkDelivered).Methods("POST")
	r.HandleFunc("/order-history/{user_id}", controllers.GetOrderHistory).Methods("GET")
	r.HandleFunc("/restaurant/{id}", controllers.GetRestaurant).Methods("GET")

	return r
}

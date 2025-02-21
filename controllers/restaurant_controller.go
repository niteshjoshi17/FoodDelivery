package controllers

import (
	"net/http"
)

func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Restaurant details!"))
}

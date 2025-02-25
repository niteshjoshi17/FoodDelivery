package routes

import (
	"food-delivery/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	// 🔹 Default Route (Homepage)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Food Delivery API!"})
	})

	// 🔹 Order Routes
	r.POST("/orders", controllers.PlaceOrder)
	r.GET("/orders", controllers.GetOrders)
	r.GET("/orders/:id", controllers.GetOrderByID)
	r.PUT("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	// 🔹 Restaurant Routes
	r.POST("/restaurants", controllers.CreateRestaurant)
	r.GET("/restaurants", controllers.GetRestaurants)
	r.GET("/restaurants/:id", controllers.GetRestaurantByID)
	r.PUT("/restaurants/:id", controllers.UpdateRestaurant)
	r.DELETE("/restaurants/:id", controllers.DeleteRestaurant)

	return r
}

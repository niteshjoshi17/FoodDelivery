package controllers

import (
	"food-delivery/config"
	"food-delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 📌 1️⃣ CREATE RESTAURANT
func CreateRestaurant(c *gin.Context) {
	var restaurant models.Restaurant

	// JSON binding check
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into DB
	if err := config.DB.Create(&restaurant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create restaurant"})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{"message": "Restaurant added successfully", "restaurant": restaurant})
}

// 📌 2️⃣ GET ALL RESTAURANTS
func GetRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant

	// Fetch all restaurants
	if err := config.DB.Find(&restaurants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurants"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, restaurants)
}

// 📌 3️⃣ GET SINGLE RESTAURANT BY ID
func GetRestaurantByID(c *gin.Context) {
	var restaurant models.Restaurant
	restaurantID := c.Param("id")

	// Find restaurant
	if err := config.DB.First(&restaurant, restaurantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, restaurant)
}

// 📌 4️⃣ UPDATE RESTAURANT DETAILS
func UpdateRestaurant(c *gin.Context) {
	var restaurant models.Restaurant
	restaurantID := c.Param("id")

	// Check if restaurant exists
	if err := config.DB.First(&restaurant, restaurantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// Bind new data
	var updateData models.Restaurant
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	config.DB.Model(&restaurant).Updates(updateData)

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Restaurant updated successfully", "restaurant": restaurant})
}

// 📌 5️⃣ DELETE RESTAURANT
func DeleteRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")

	// Delete from DB
	if err := config.DB.Delete(&models.Restaurant{}, restaurantID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete restaurant"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Restaurant deleted successfully"})
}

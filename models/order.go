package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `json:"user_id"`
	RestaurantID uint      `json:"restaurant_id"`
	AgentID      *uint     `json:"agent_id"`
	Status       string    `json:"status"` // ["pending", "accepted", "in_transit", "delivered"]
	TotalPrice   float64   `json:"total_price"`
	CreatedAt    time.Time `json:"created_at"`
}

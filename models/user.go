package models

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type"` // customer / delivery_agent
}

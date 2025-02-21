package models

type Restaurant struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float32 `json:"rating"`
}

package domain

import "time"

// order //
type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderAt      time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}

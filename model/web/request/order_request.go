package request

import "time"

type OrderCreate struct {
	CustomerName string       `json:"customer_name"`
	OrderAt      time.Time    `json:"ordered_at"`
	Items        []ItemCreate `json:"items"`
}

type OrderUpdate struct {
	ID           uint         `json:"id"`
	CustomerName string       `json:"customer_name"`
	OrderAt      time.Time    `json:"ordered_at"`
	Items        []ItemUpdate `json:"items"`
}

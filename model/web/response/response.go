package response

import (
	"github.com/xvbnm48/go-learn-api-scalable/model/domain"
	"time"
)

type OrderResponse struct {
	ID           uint          `json:"id"`
	CustomerName string        `json:"customer_name"`
	OrderedAt    time.Time     `json:"ordered_at"`
	Items        []domain.Item `json:"items"`
}

package models

import (
    "time"

    "github.com/google/uuid"
)

type Order struct {
    ID         uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    UserID     uuid.UUID   `gorm:"type:uuid" json:"user_id"`
    TotalPrice float64     `json:"total_price"`
    Status     string      `json:"status"`
    CreatedAt  time.Time   `json:"created_at"`

    Items []OrderItem `json:"items"`
}
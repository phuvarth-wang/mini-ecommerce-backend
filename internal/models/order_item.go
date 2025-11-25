package models

import (
    "time"

    "github.com/google/uuid"
)

type OrderItem struct {
    ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    OrderID         uuid.UUID `gorm:"type:uuid" json:"order_id"`
    ProductID       uuid.UUID `gorm:"type:uuid" json:"product_id"`
    Quantity        int       `json:"quantity"`
    PriceAtPurchase float64   `json:"price_at_purchase"`

    Product Product `gorm:"foreignKey:ProductID" json:"product"`

    CreatedAt time.Time `json:"created_at"`
}
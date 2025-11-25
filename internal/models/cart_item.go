package models

import (
    "time"

    "github.com/google/uuid"
)

type CartItem struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

    CartID    uuid.UUID `gorm:"type:uuid" json:"cart_id"`
    ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
    
    Quantity int `json:"quantity"`

    Product Product `gorm:"foreignKey:ProductID" json:"product"`

    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
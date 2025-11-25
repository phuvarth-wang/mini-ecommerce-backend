package models

import (
    "time"

    "github.com/google/uuid"
)

type Product struct {
    ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    ImageURL    string    `json:"image_url"`
    Stock       int       `json:"stock"`

    CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
    Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`

    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
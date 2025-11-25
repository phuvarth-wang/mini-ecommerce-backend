package models

import (
    "time"

    "github.com/google/uuid"
)

type Cart struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"` // Supabase Auth user ID
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

    Items []CartItem `json:"items"`
}

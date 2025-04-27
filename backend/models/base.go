package models

import "time"

// Base 包含公共字段
type Base struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

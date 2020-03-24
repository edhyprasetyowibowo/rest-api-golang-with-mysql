package models

import (
	"time"
)

type (
	// Users
	Users struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Address   string    `json:"address"`
		Phone     string    `json:"phone"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

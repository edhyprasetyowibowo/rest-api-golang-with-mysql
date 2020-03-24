package models

import (
	"time"
)

type (
	// Users
	Users struct {
		ID        int       `json:"id"`
		Name      string    `name:"name"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

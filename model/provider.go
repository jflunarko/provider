package model

import "time"

type Provider struct {
	ID        int        `json:"id"`
	Provider  string     `json:"provider"`
	Location  string     `json:"location"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type NewProvider struct {
	Provider string `json:"provider"`
	Location string `json:"location"`
}

type UpdateProvider struct {
	Provider string `json:"provider"`
	Location string `json:"location"`
}

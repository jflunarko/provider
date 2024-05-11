package model

import "time"

type Packet struct {
	ID         int        `json:"id"`
	ProviderID int        `json:"provider_id"`
	Name       string     `json:"name"`
	DataLimit  string     `json:"data_limit"`
	Validity   string     `json:"validity"`
	Price      int        `json:"price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type NewPacket struct {
	ProviderID int    `json:"provider_id"`
	Name       string `json:"name"`
	DataLimit  string `json:"data_limit"`
	Validity   string `json:"validity"`
	Price      int    `json:"price"`
}

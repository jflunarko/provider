package model

type UserProvider struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	ProviderID int `json:"provider_id"`
}

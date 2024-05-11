package model

import "time"

type User struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	Provider  *[]Provider `json:"provider" gorm:"many2many:user_provider"`
	Packet    *[]Packet   `json:"packet" gorm:"many2many:user_packet"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUser struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

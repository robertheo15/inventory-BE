package models

import "time"

var (
	Owner   = 0
	Admin   = 1
	Courier = 2
)

type User struct {
	ID        string    `json:"user_id"`
	Name      string    `json:"full_name"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone_number"`
	Email     string    `json:"email"`
	Role      int       `json:"role"`
	Active    bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

type Customer struct {
	ID        string    `json:"c_id"`
	Name      string    `json:"full_name"`
	Phone     string    `json:"phone_number"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

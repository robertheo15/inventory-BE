package models

import (
	"time"
)

var (
	Owner   = 0
	Admin   = 1
	Courier = 2
)

type User struct {
	ID          string    `json:"id"`
	FullName    string    `json:"full_name"`
	Password    string    `validate:"required" json:"password"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `validate:"required,email" json:"email" `
	Role        int32     `json:"role"`
	Active      int32     `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}

type Customer struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone_number"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

type Supplier struct {
	ID        string    `json:"id"`
	BrandName string    `json:"brand_name"`
	Phone     string    `json:"phone_number"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

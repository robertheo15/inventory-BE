package models

import (
	"time"
)

type Product struct {
	ID          string     `json:"id"`
	ProductID   string     `json:"product_id"`
	Name        string     `json:"name" valid:"required~Name is required"`
	Brand       string     `json:"brand"`
	Description string     `json:"description"`
	Stock       int32      `json:"stock"`
	BasePrice   float64    `json:"base_price"`
	EceranPrice float64    `json:"eceran_price"`
	GrosirPrice float64    `json:"grosir_price"`
	Image       string     `json:"image"`
	Type        string     `json:"type"`
	Children    []*Product `json:"children"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedBy   string     `json:"updated_by"`
}

type ProductVariant struct {
	ID        string    `json:"pv_id"`
	Name      string    `json:"name"`
	Colour    string    `json:"colour"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

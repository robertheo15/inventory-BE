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
	SupplierID  string     `json:"supplier_id"`
	BasePrice   float64    `json:"base_price"`
	EceranPrice float64    `json:"eceran_price"`
	GrosirPrice float64    `json:"grosir_price"`
	Children    []*Product `json:"children"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedBy   string     `json:"updated_by"`
}

type ProductVariant struct {
	ID               string            `json:"id"`
	ProductVariantID string            `json:"pv_id"`
	ProductID        string            `json:"p_id"`
	ProductName      string            `json:"product_name"`
	Children         []*ProductVariant `json:"children"`
	Type             string            `json:"type"`
	Name             string            `json:"name"`
	Stock            int32             `json:"stock"`
	Location         string            `json:"location"`
	Colour           string            `json:"colour"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	CreatedBy        string            `json:"created_by"`
	UpdatedBy        string            `json:"updated_by"`
}

type Price struct {
	ID        string    `json:"id"`
	Eceran    float64   `json:"eceran"`
	Grosir    float64   `json:"grosir"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

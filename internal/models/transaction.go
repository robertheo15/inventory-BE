package models

import "time"

type Transaction struct {
	ID         string         `json:"transaction_id"`
	Invoice    string         `json:"invoice"`
	Status     string         `json:"status"`
	Type       string         `json:"type"`
	TotalPrice float64        `json:"total_price"`
	Childs     *[]Transaction `json:"child"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	CreatedBy  string         `json:"created_by"`
	UpdatedBy  string         `json:"updated_by"`
}

type TransactionDetail struct {
	ID            string    `json:"td_id"`
	TransactionID string    `json:"t_id"`
	ProductID     string    `json:"p_id"`
	Price         float64   `json:"price"`
	Qty           int32     `json:"qty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
}

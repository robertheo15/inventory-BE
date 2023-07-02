package models

import "time"

type Transaction struct {
	ID                 string               `json:"id"`
	TransactionID      string               `json:"transaction_id,omitempty"`
	CID                string               `json:"c_id"`
	Invoice            string               `json:"invoice"`
	Status             string               `json:"status"`
	Type               string               `json:"type"`
	Methode            string               `json:"methode"`
	Customer           *Customer            `json:"customer"`
	Supplier           *Supplier            `json:"supplier"`
	Children           []*Transaction       `json:"children,omitempty"`
	TransactionDetails []*TransactionDetail `json:"transaction_details,omitempty"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	CreatedBy          string               `json:"created_by"`
	UpdatedBy          string               `json:"updated_by"`
}

type TransactionDetail struct {
	ID               string    `json:"td_id"`
	ProductVariantID string    `json:"pv_id"`
	TransactionID    string    `json:"t_id"`
	ProductID        string    `json:"p_id"`
	Price            float64   `json:"price"`
	Qty              int32     `json:"qty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedBy        string    `json:"updated_by"`
}

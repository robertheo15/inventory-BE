package response

import "time"

type TransactionCustomerRequest struct {
	ID                 string               `json:"id"`
	TransactionID      string               `json:"transaction_id,omitempty"`
	CID                string               `json:"c_id"`
	Invoice            string               `json:"invoice"`
	Status             string               `json:"status"`
	Type               string               `json:"type"`
	Methode            string               `json:"methode"`
	Children           []*Transaction       `json:"children,omitempty"`
	TransactionDetails []*TransactionDetail `json:"transaction_details,omitempty"`
}

type TransactionSupplierRequest struct {
	ID                 string               `json:"id"`
	Invoice            string               `json:"invoice"`
	Status             string               `json:"status"`
	Type               string               `json:"type"`
	Methode            string               `json:"methode"`
	TransactionDetails []*TransactionDetail `json:"transaction_details,omitempty"`
}

type Transaction struct {
	ID                 string               `json:"id"`
	TransactionID      string               `json:"transaction_id,omitempty"`
	CID                string               `json:"c_id"`
	Invoice            string               `json:"invoice"`
	Status             string               `json:"status"`
	Type               string               `json:"type"`
	Methode            string               `json:"methode"`
	Children           []*Transaction       `json:"children,omitempty"`
	TransactionDetails []*TransactionDetail `json:"transaction_details,omitempty"`
	CreatedAt          time.Time            `json:"created_at,omitempty"`
	UpdatedAt          time.Time            `json:"updated_at,omitempty"`
	CreatedBy          string               `json:"created_by,omitempty"`
	UpdatedBy          string               `json:"updated_by,omitempty"`
}

type TransactionDetail struct {
	ID               string    `json:"td_id"`
	TransactionID    string    `json:"t_id"`
	ProductID        string    `json:"product_id"`
	ProductVariantID string    `json:"product_variant_id"`
	Price            float64   `json:"price"`
	Qty              int32     `json:"qty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	CreatedBy        string    `json:"created_by,omitempty"`
	UpdatedBy        string    `json:"updated_by,omitempty"`
}

type TransactionCustomerResponse struct {
	ID                 string               `json:"id"`
	TransactionID      string               `json:"transaction_id,omitempty"`
	CID                string               `json:"c_id,omitempty"`
	Invoice            string               `json:"invoice"`
	Status             string               `json:"status"`
	Type               string               `json:"type"`
	Children           []*Transaction       `json:"children,omitempty"`
	TransactionDetails []*TransactionDetail `json:"transaction_details,omitempty"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	CreatedBy          string               `json:"created_by"`
	UpdatedBy          string               `json:"updated_by"`
}

type TransactionSupplierResponse struct {
	ID                 string               `json:"id"`
	Invoice            string               `json:"invoice"`
	CID                string               `json:"c_id,omitempty"`
	Status             string               `json:"status"`
	Type               string               `json:"type"`
	TransactionDetails []*TransactionDetail `json:"transaction_details,omitempty"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	CreatedBy          string               `json:"created_by"`
	UpdatedBy          string               `json:"updated_by"`
}

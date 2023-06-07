// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: transaction.sql

package sqlc

import (
	"context"
	"time"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (id,
                          c_id,
                          transaction_id,
                          invoice,
                          status,
                          type,
                          created_at,
                          updated_at,
                          created_by,
                          updated_by)
VALUES (gen_random_uuid(),
        $1::char(36),
        $2::char(36),
        $3::varchar,
        $4::varchar,
        $5::varchar,
        now() at time zone 'Asia/Jakarta',
        now() at time zone 'Asia/Jakarta',
        $6::varchar,
        $7::varchar)
RETURNING id::char(36)
`

type CreateTransactionParams struct {
	CID           string `json:"c_id"`
	TransactionID string `json:"transaction_id"`
	Invoice       string `json:"invoice"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	CreatedBy     string `json:"created_by"`
	UpdatedByy    string `json:"updated_byy"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.CID,
		arg.TransactionID,
		arg.Invoice,
		arg.Status,
		arg.Type,
		arg.CreatedBy,
		arg.UpdatedByy,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteTransactionByID = `-- name: DeleteTransactionByID :one
DELETE
FROM transactions
WHERE id = $1::char(36)
    returning id
`

func (q *Queries) DeleteTransactionByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteTransactionByID, id)
	err := row.Scan(&id)
	return id, err
}

const getTransactionByID = `-- name: GetTransactionByID :one
SELECT
    id::char(36),
    c_id::char(36),
    transaction_id::char(36),
    invoice::varchar,
    status::varchar,
    type::varchar,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar
FROM transactions WHERE id = $1::char(36)
`

type GetTransactionByIDRow struct {
	ID            string    `json:"id"`
	CID           string    `json:"c_id"`
	TransactionID string    `json:"transaction_id"`
	Invoice       string    `json:"invoice"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
}

func (q *Queries) GetTransactionByID(ctx context.Context, id string) (GetTransactionByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getTransactionByID, id)
	var i GetTransactionByIDRow
	err := row.Scan(
		&i.ID,
		&i.CID,
		&i.TransactionID,
		&i.Invoice,
		&i.Status,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const getTransactions = `-- name: GetTransactions :many
SELECT id::char(36),
       c_id::char(36),
       transaction_id::char(36),
       invoice::varchar,
       status::varchar,
       type::varchar,
       created_at::timestamp,
       updated_at::timestamp,
       created_by::varchar,
       updated_by::varchar
FROM transactions
`

type GetTransactionsRow struct {
	ID            string    `json:"id"`
	CID           string    `json:"c_id"`
	TransactionID string    `json:"transaction_id"`
	Invoice       string    `json:"invoice"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
}

func (q *Queries) GetTransactions(ctx context.Context) ([]GetTransactionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTransactionsRow
	for rows.Next() {
		var i GetTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.CID,
			&i.TransactionID,
			&i.Invoice,
			&i.Status,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransactionByID = `-- name: UpdateTransactionByID :one
UPDATE transactions
SET c_id = $1::char(36),
    transaction_id = $2::char(36),
    invoice = $3::varchar,
    status = $4::varchar,
    type = $5::varchar,
    created_at = $6::timestamp,
    updated_at = (now() at time zone 'Asia/Jakarta')::timestamp,
    created_by = $7::varchar,
    updated_by = $8::varchar
WHERE id = $9::char(36) RETURNING
    id::char(36),
    c_id::char(36),
    transaction_id::char(36),
    invoice::varchar,
    status::varchar,
    type::varchar,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar
`

type UpdateTransactionByIDParams struct {
	CID           string    `json:"c_id"`
	TransactionID string    `json:"transaction_id"`
	Invoice       string    `json:"invoice"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	ID            string    `json:"id"`
}

type UpdateTransactionByIDRow struct {
	ID            string    `json:"id"`
	CID           string    `json:"c_id"`
	TransactionID string    `json:"transaction_id"`
	Invoice       string    `json:"invoice"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
}

func (q *Queries) UpdateTransactionByID(ctx context.Context, arg UpdateTransactionByIDParams) (UpdateTransactionByIDRow, error) {
	row := q.db.QueryRowContext(ctx, updateTransactionByID,
		arg.CID,
		arg.TransactionID,
		arg.Invoice,
		arg.Status,
		arg.Type,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedBy,
		arg.ID,
	)
	var i UpdateTransactionByIDRow
	err := row.Scan(
		&i.ID,
		&i.CID,
		&i.TransactionID,
		&i.Invoice,
		&i.Status,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}
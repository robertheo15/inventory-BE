-- name: GetTransactions :many
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
FROM transactions;

-- name: GetTransactionByID :one
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
FROM transactions WHERE id = @id::char(36);

-- name: GetTransactionByChildID :many
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
FROM transactions WHERE transaction_id = @transaction_id::char(36);

-- name: CreateTransaction :one
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
        @c_id::char(36),
        @transaction_id::char(36),
        @invoice::varchar,
        @status::varchar,
        @type::varchar,
        now() at time zone 'Asia/Jakarta',
        now() at time zone 'Asia/Jakarta',
        @created_by::varchar,
        @updated_byy::varchar)
RETURNING id::char(36), invoice::varchar;

-- name: UpdateTransactionByID :one
UPDATE transactions
SET c_id = @c_id::char(36),
    transaction_id = @transaction_id::char(36),
    status = @status::varchar,
    type = @type::varchar,
    updated_at = (now() at time zone 'Asia/Jakarta')::timestamp,
    created_by = @created_by::varchar,
    updated_by = @updated_by::varchar
WHERE id = @id::char(36) RETURNING
    id::char(36),
    c_id::char(36),
    transaction_id::char(36),
    invoice::varchar,
    status::varchar,
    type::varchar,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar;

-- name: DeleteTransactionByID :one
DELETE
FROM transactions
WHERE id = @id::char(36)
    returning id;
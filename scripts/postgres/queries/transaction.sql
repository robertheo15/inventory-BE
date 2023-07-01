-- name: GetTransactions :many
SELECT id::char(36),
       c_id::char(36),
       transaction_id::char(36),
       invoice::varchar,
       status::varchar,
       type::varchar,
       methode::varchar,
       created_at::timestamp,
       updated_at::timestamp,
       created_by::varchar,
       updated_by::varchar
FROM transactions WHERE type = ANY($1::varchar[]);

-- name: GetTransactionsByStatus :many
SELECT  id::char(36),
        c_id::char(36),
        transaction_id::char(36),
        invoice::varchar,
        status::varchar,
        type::varchar,
        methode::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM transactions WHERE type = ANY($1::varchar[]) and status = @status::varchar
ORDER BY created_at ASC;

-- name: GetTransactionsSupplierByStatus :many
SELECT  t.id::char(36),
        t.c_id::char(36),
        t.transaction_id::char(36),
        t.invoice::varchar,
        t.status::varchar,
        t.type::varchar,
        t.methode::varchar,
        s.brand_name::varchar,
        t.created_at::timestamp,
        t.updated_at::timestamp,
        t.created_by::varchar,
        t.updated_by::varchar
FROM transactions t
     INNER JOIN suppliers s on t.s_id = s.id
WHERE t.type = ANY($1::varchar[]) and t.status = @status::varchar
ORDER BY created_at ASC;


-- name: GetTransactionByID :one
SELECT
    id::char(36),
    c_id::char(36),
    transaction_id::char(36),
    invoice::varchar,
    status::varchar,
    type::varchar,
    methode::varchar,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar
FROM transactions WHERE id = @id::char(36) and type = ANY($1::varchar[]);

-- name: GetTransactionByChildID :many
SELECT
    id::char(36),
        c_id::char(36),
        transaction_id::char(36),
        invoice::varchar,
        status::varchar,
        type::varchar,
        methode::varchar,
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
                          methode,
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
        @methode::varchar,
        now() at time zone 'Asia/Jakarta',
        now() at time zone 'Asia/Jakarta',
        @created_by::varchar,
        @updated_by::varchar)
RETURNING id::char(36), invoice::varchar;

-- name: UpdateTransactionByID :one
UPDATE transactions
SET c_id = @c_id::char(36),
    transaction_id = @transaction_id::char(36),
    status = @status::varchar,
    type = @type::varchar,
    methode = @methode::varchar,
    updated_at = (now() at time zone 'Asia/Jakarta')::timestamp,
    updated_by = @updated_by::varchar
WHERE id = @id::char(36) RETURNING
    id::char(36),
    c_id::char(36),
    transaction_id::char(36),
    invoice::varchar,
    status::varchar,
    type::varchar,
    methode::varchar,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar;

-- name: UpdateStatusTransactionByID :one
UPDATE transactions
SET status = @status::varchar,
    updated_at = (now() at time zone 'Asia/Jakarta')::timestamp,
    updated_by = @updated_by::varchar
WHERE id = @id::char(36) RETURNING
    id::char(36);

-- name: DeleteTransactionByID :one
DELETE
FROM transactions
WHERE id = @id::char(36)
    returning id;
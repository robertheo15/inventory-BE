-- name: CreateTransactionDetail :one
INSERT INTO transaction_details (id, t_id, p_id, price, qty, created_at, updated_at, created_by, updated_by)
VALUES (gen_random_uuid(),
        @t_id::char(36),
        @p_id::char(36),
        @price::float,
        @qty::integer,
        now() at time zone 'Asia/Jakarta',
        now() at time zone 'Asia/Jakarta',
        @created_by::varchar,
        @updated_by::varchar) RETURNING id::char(36);

-- name: GetTransactionDetails :many
SELECT  id::char(36),
        t_id::char(36),
        p_id::char(36),
        price::float,
        qty::integer,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM transaction_details;

-- name: GetTransactionDetailByID :one
SELECT  id::char(36),
        t_id::char(36),
        p_id::char(36),
        price::float,
        qty::integer,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM transaction_details
WHERE id = @id::char(36);

-- name: GetTransactionDetailByTID :many
SELECT  id::char(36),
        t_id::char(36),
        p_id::char(36),
        price::float,
        qty::integer,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM transaction_details
WHERE t_id = @t_id::char(36);

-- name: UpdateTransactionDetailByID :one
UPDATE transaction_details
SET t_id = @t_id::char(36),
    p_id = @t_id::char(36),
    price = @price::float,
    qty = @qty::integer,
    updated_at = now() at time zone 'Asia/Jakarta',
    updated_by = @updated_by::varchar
WHERE id = @id::char(36)
    RETURNING id::char(36),
    t_id::char(36),
    p_id::char(36),
    price::float,
    qty::integer,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar;

-- name: DeleteTransactionDetailByID :one
DELETE
FROM transaction_details
WHERE id = @id::char(36)
RETURNING id::char(36);

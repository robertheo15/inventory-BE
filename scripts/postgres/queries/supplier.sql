-- name: GetSuppliers :many
SELECT id::char(36),
       brand_name::varchar,
       phone_number::varchar,
       address::varchar,
       email::varchar,
       created_at::timestamp,
       updated_at::timestamp,
       created_by::varchar,
       updated_by::varchar
FROM suppliers;

-- name: GetSupplierByID :one
SELECT id::char(36),
        brand_name::varchar,
        phone_number::varchar,
        address::varchar,
        email::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM suppliers
WHERE id = @id::char(36);

-- name: CreateSupplier :one
INSERT INTO suppliers (id,
                       brand_name,
                       phone_number,
                       address,
                       email,
                       created_at,
                       updated_at,
                       created_by,
                       updated_by)
VALUES (gen_random_uuid(),
        @brand_name::varchar,
        @phone_number::varchar,
        @address::varchar,
        @email::varchar,
        (now() at time zone 'Asia/Jakarta')::timestamp,
        (now() at time zone 'Asia/Jakarta')::timestamp,
        @created_by::varchar,
        @updated_by::varchar) RETURNING id::char(36);


-- name: UpdateSupplierByID :one
UPDATE suppliers
SET brand_name = @brand_name::varchar,
    phone_number = @phone_number::varchar,
    address = @address::varchar,
    email = @email::varchar,
    created_at = @created_at::timestamp,
    updated_at= now()::timestamp,
    created_by= @created_by::varchar,
    updated_by= @updated_by::varchar
WHERE id = @id::char(36) RETURNING
    id::char(36),
    brand_name::varchar,
    phone_number::varchar,
    address::varchar,
    email::varchar,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar;

-- name: DeleteSupplierByID :one
DELETE
FROM suppliers
WHERE id = @id::char(36) returning id::char(36);


-- name: CreateCustomer :one
INSERT INTO customers (id, full_name, phone_number, address, created_at, updated_at, created_by,
                       updated_by)
VALUES (gen_random_uuid(),
        @full_name::varchar,
        @phone_number::varchar,
        @address::varchar,
        now() at time zone 'Asia/Jakarta',
        now() at time zone 'Asia/Jakarta',
        @created_by::varchar,
        @updated_by::varchar) RETURNING id::char(36);

-- name: GetCustomers :many
SELECT  id::char(36),
        full_name::varchar,
        phone_number::varchar,
        address::varchar,
        email::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM customers;


-- name: GetCustomersByID :one
SELECT  id::char(36),
        full_name::varchar,
        phone_number::varchar,
        address::varchar,
        email::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM customers WHERE id= @id::char(36) ;

-- name: UpdateCustomerByID :one
UPDATE customers
SET full_name = @full_name::varchar,
    phone_number = @phone_number::varchar,
    address = @address::varchar,
    email = @email::varchar,
    created_at = @created_at::timestamp,
    updated_at = (now() at time zone 'Asia/Jakarta')::timestamp,
    updated_by = @updated_by::varchar
WHERE   id = @id::char(36) RETURNING
    id::char(36),
    full_name::varchar,
    phone_number::varchar,
    address::varchar,
    email::varchar,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar;

-- name: DeleteCustomerByID :one
DELETE
FROM customers
WHERE id = @id::varchar(36) returning id::varchar(36);


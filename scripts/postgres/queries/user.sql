-- name: CreateUser :one
INSERT INTO users (id, full_name, password, phone_number, email, role, active,
                          created_at, updated_at, created_by,
                          updated_by)
VALUES ((gen_random_uuid()):: char (36), @full_name::varchar, @password::varchar, @phone_number::varchar,
        @email::varchar,
        @role::integer,
        @active::integer, (now() at time zone 'Asia/Jakarta'):: timestamp,
        (now() at time zone 'Asia/Jakarta'):: timestamp, @created_by::varchar,
        @updated_by::varchar) RETURNING id::char(36);

-- name: GetUsers :many
SELECT full_name::varchar,
        phone_number::varchar,
        email::varchar,
        role::integer,
        active::integer
 FROM users;

-- name: GetUserByID :one
SELECT id::char(36), full_name::varchar, password::varchar, phone_number::varchar, email::varchar, role::integer, active::integer, created_at::timestamp, updated_at::timestamp, created_by::varchar, updated_by::varchar
FROM users
WHERE id = @id::char(36);

-- name: GetUserByEmail :one
SELECT id::char(36), full_name::varchar, password::varchar,phone_number::varchar, email::varchar, role::integer, active::integer, created_at::timestamp, updated_at::timestamp, created_by::varchar, updated_by::varchar
FROM users
WHERE email = @email::varchar;

-- name: UpdateUserByID :one
UPDATE users
SET full_name = @full_name::varchar,
        phone_number = @phone_number::varchar,
        email = @email::varchar,
        role = @role::integer,
        active = @active::integer,
        updated_at = (now() at time zone 'Asia/Jakarta'):: timestamp,
        created_by = @created_by::varchar,
        updated_by = @updated_by::varchar
WHERE id = @id:: char (36) returning id;

-- name: UpdateUserPasswordByID :one
UPDATE users
SET     password = @password::varchar,
        updated_at = (now() at time zone 'Asia/Jakarta'):: timestamp,
        created_by = @created_by::varchar,
        updated_by = @updated_by::varchar
WHERE id = @id:: char (36) returning id::char(36), created_at::timestamp;

-- name: DeleteUserByID :one
DELETE
FROM users
WHERE id = @id::char(36) returning id;

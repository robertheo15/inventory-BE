-- name: CreateUser :one
INSERT INTO public.users (id, full_name, password, phone_number, email, role, active,
                          created_at, updated_at, created_by,
                          updated_by)
VALUES ((gen_random_uuid()):: char (36), @full_name::varchar, @password::varchar, @phone_number::varchar,
        @email::varchar,
        @role::integer,
        @active::boolean, (now() at time zone 'Asia/Jakarta'):: timestamp,
        (now() at time zone 'Asia/Jakarta'):: timestamp, @created_by::varchar,
        @updated_by::varchar) RETURNING *;

-- name: GetUserByID :one
select *
from users
where id = @id::varchar;

-- name: UpdateUserByID :one
UPDATE users
SET full_name = @full_name::varchar,
        password = @password::varchar,
        phone_number = @phone_number::varchar,
        email = @email::varchar,
        role = @role::integer,
        active = @active::boolean,
        updated_at = (now() at time zone 'Asia/Jakarta'):: timestamp,
        updated_by = @updated_by::varchar
WHERE id = @id::char(36) returning id;

-- name: DeleteUserByID :one
DELETE
FROM users
WHERE id = @id::char(36) returning id;
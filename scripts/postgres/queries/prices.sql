-- name: GetPrice :one
SELECT id::char(36),
        eceran::float(36),
        grosir::float(36),
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM prices limit 1;

-- name: UpdatePriceByID :one
UPDATE product_variants
SET eceran = @eceran::float,
    grosir = @grosir::float,
    updated_at = (now() at time zone 'Asia/Jakarta')::timestamp,
    updated_by = @updated_by::varchar
WHERE id = @id::char(36) returning id::char(36);

-- name: CreatePrice :one
INSERT INTO prices (id, eceran, grosir, created_at, updated_at, created_by, updated_by)
VALUES (gen_random_uuid(), @eceran::float, @grosir::float, now() at time zone 'Asia/Jakarta',
        now() at time zone 'Asia/Jakarta', @created_by::varchar, @updated_by::varchar) RETURNING id::char(36);

-- name: DeletePriceByID :one
DELETE
FROM prices
WHERE id = @id::char(36)
    RETURNING id;
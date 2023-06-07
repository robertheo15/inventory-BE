-- name: CreateProductVariant :one
INSERT INTO product_variants (id, p_id, name, colour,
                             created_at, updated_at, created_by, updated_by)
VALUES (gen_random_uuid(), @p_id::char(36), @name::varchar, @colour::varchar,
        now() at time zone 'Asia/Jakarta', now() at time zone 'Asia/Jakarta', @created_by::varchar, @updated_by::varchar)
RETURNING id::char(36);

-- name: GetProductVariants :many
SELECT id:: char(36),
        p_id::char(36),
        name::varchar,
        colour::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM product_variants;

-- name: GetProductVariantsByProductID :many
SELECT id:: char(36),
        p_id::char(36),
        name::varchar,
        colour::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM product_variants WHERE p_id = @product_id::char(36);

-- name: GetProductVariantByID :one
SELECT id:: char(36),
        p_id::char(36),
        name::varchar,
        colour::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM product_variants
WHERE id = @id::char(36);

-- name: UpdateProductVariantByID :one
UPDATE product_variants SET
        p_id= @p_id::char(36),
        name= @name::varchar,
        colour= @colour::varchar,
        updated_at= @updated_at::timestamp,
        created_by= @created_by::varchar,
        updated_by= @updated_by::varchar
WHERE id = @id::char(36) RETURNING id::char(36), p_id::char(36), name::varchar, colour::varchar,
    created_at::timestamp, updated_at::timestamp, created_by::varchar, updated_by::varchar;

-- name: DeleteProductVariantByID :one
DELETE
FROM product_variants
WHERE id = @id::char(36)
RETURNING id;
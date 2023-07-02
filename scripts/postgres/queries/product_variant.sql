-- name: CreateProductVariant :one
INSERT INTO product_variants (id, p_id, pv_id, name, colour, stock, location, type,
                             created_at, updated_at, created_by, updated_by)
VALUES (gen_random_uuid(), @p_id::char(36), @pv_id::char(36), @name::varchar, @colour::varchar, @stock::integer, @location::varchar,
        @type::varchar,
        now() at time zone 'Asia/Jakarta', now() at time zone 'Asia/Jakarta', @created_by::varchar, @updated_by::varchar)
RETURNING id::char(36);

-- name: GetProductVariants :many
SELECT pv.id::char(36),
       pv.pv_id::char(36),
       p.id::char(36),
       p.name::varchar,
       pv.name::varchar,
       p.brand::varchar,
       pv.colour::varchar,
       pv.stock::integer,
       pv.location::varchar,
       pv.type::varchar,
       pv.created_at::timestamp,
       pv.updated_at::timestamp,
       pv.created_by::varchar,
       pv.updated_by::varchar
FROM  product_variants as pv
        INNER JOIN products p ON pv.p_id = p.id
ORDER BY pv.name ASC;


-- name: GetProductVariantsByProductID :many
SELECT id:: char(36),
        p_id::char(36),
        pv_id::char(36),
        name::varchar,
        colour::varchar,
        stock::integer,
        location::varchar,
        type::varchar,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM product_variants WHERE p_id = @product_id::char(36) and location = 'gudang';

-- name: GetProductVariantByID :one
SELECT id:: char(36),
        p_id::char(36),
        pv_id::char(36),
        name::varchar,
        colour::varchar,
        stock::integer,
        location::varchar,
        type::varchar,
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
        stock = @stock::integer,
        updated_at= @updated_at::timestamp,
        created_by= @created_by::varchar,
        updated_by= @updated_by::varchar
WHERE id = @id::char(36) RETURNING id::char(36), p_id::char(36), name::varchar, colour::varchar, stock::integer, location::varchar,
    created_at::timestamp, updated_at::timestamp, created_by::varchar, updated_by::varchar;

-- name: UpdateProductVariantStockByID :one
UPDATE product_variants
SET stock= @stock::integer,
    updated_at = (now() at time zone 'Asia/Jakarta')::timestamp,
    updated_by = @updatedBy:: varchar
WHERE id = @id::char(36) returning id::char(36);

-- name: DeleteProductVariantByID :one
DELETE
FROM product_variants
WHERE id = @id::char(36)
RETURNING id;
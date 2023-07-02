-- name: GetProducts :many
SELECT id::char(36),
    product_id::char(36),
    name::varchar,
    brand::varchar,
    description::varchar,
    base_price::float,
    price_eceran::float,
    price_grosir::float,
    created_at::timestamp,
    updated_at::timestamp,
    created_by::varchar,
    updated_by::varchar
FROM products;

-- name: GetProductByID :one
SELECT id::char(36),
        product_id::char(36),
        name::varchar,
        brand::varchar,
        description::varchar,
        base_price::float,
        price_eceran::float,
        price_grosir::float,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM products
WHERE id = @productID::char(36);

-- name: GetProductsBySupplierID :many
SELECT id::char(36),
        product_id::char(36),
        s_id::char(36),
        name::varchar,
        brand::varchar,
        description::varchar,
        base_price::float,
        price_eceran::float,
        price_grosir::float,
        created_at::timestamp,
        updated_at::timestamp,
        created_by::varchar,
        updated_by::varchar
FROM products where s_id = @s_id::char(36);

-- name: CreateProduct :one
INSERT INTO products (id, product_id, name, brand, description, base_price,
                      price_eceran, price_grosir, created_at, updated_at, created_by, updated_by)
VALUES ((gen_random_uuid()::char(36)), @product_id::char(36), @name::varchar, @brand::varchar,
        @description::varchar, @base_price::float,
        @price_eceran::float, @price_grosir::float,
        now() at time zone 'Asia/Jakarta', now() at time zone 'Asia/Jakarta', @created_by::varchar, @updated_by::varchar) returning id;

-- name: UpdateProductByID :one
UPDATE products
SET name= @name:: varchar,
    product_id= @product_id:: varchar,
    brand= @brand:: varchar,
    description = @description:: varchar,
    base_price = @basePrice:: float,
    price_eceran = @priceEceran:: float,
    price_grosir = @priceGrosir:: float,
    updated_at = (now() at time zone 'Asia/Jakarta'):: timestamp,
    updated_by = @updatedBy:: varchar
WHERE id = @id:: char (36) returning id::char(36), product_id::char(36),
    name:: varchar, brand::varchar, description::varchar, base_price::float, price_eceran::float, price_grosir::float
    , created_at::timestamp, updated_at::timestamp, created_by::varchar, updated_by::varchar;

-- name: UpdateProductStockByID :exec
UPDATE products
SET stock= @stock:: integer,
    updated_at = (now() at time zone 'Asia/Jakarta'):: timestamp,
    updated_by = @updatedBy:: varchar
WHERE id = @id:: char (36);

-- name: UpdateProductStockChildByID :exec
UPDATE products
SET stock= @stock:: integer,
    updated_at = (now() at time zone 'Asia/Jakarta'):: timestamp,
    updated_by = @updatedBy:: varchar
WHERE product_id = @product_id:: char (36);

-- name: DeleteProductByID :one
DELETE
FROM products
WHERE id = @id::char(36)
    returning id;
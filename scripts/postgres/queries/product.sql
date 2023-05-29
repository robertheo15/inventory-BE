-- name: GetProducts :many
SELECT *
FROM products;

-- name GetProductByID :one
SELECT *
FROM products
WHERE id = @productID::char(36);

-- name CreateProduct :exec
INSERT INTO products (id, product_id, name, brand, description, stock, base_price,
                      price_eceran, price_grosir, image, type, created_at, updated_at, created_by, updated_by)
VALUES ((gen_random_uuid()::char(36)), @productID::char(36), @name::varchar, @brand::varchar,
        @description::varchar, @stock::integer, @basePrice::numeric,
        @priceEceran::numeric, @priceGrosir::numeric, @image::varchar, @type::varchar,
        now() at time zone 'Asia/Jakarta', now() at time zone 'Asia/Jakarta', @createdBy::varchar, @updatedBy::varchar);

-- name UpdateProductByID :one
UPDATE SET name=@name:: varchar,
    brand=@brand:: varchar,
    description=@description:: varchar,
    stock=@stock:: integer,
    base_price=@basePrice:: numeric,
    price_eceran=@priceEceran:: numeric,
    price_grosir=@priceGrosir:: numeric,
    image=@image:: varchar,
    type =@type:: varchar,
    updated_at=now(),
    updated_by=@updatedBy:: varchar
WHERE id = @id:: char (36)
    returning *;

DELETE
FROM products
WHERE id = @id::char(36)
    returning id;
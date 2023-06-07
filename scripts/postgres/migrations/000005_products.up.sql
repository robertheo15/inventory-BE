CREATE TABLE products
(
    id           char(36) PRIMARY KEY,
    s_id         char(36)
        references suppliers,
    product_id   char(36),
    name         varchar,
    brand        varchar,
    description  varchar,
    stock        integer,
    base_price   numeric,
    price_eceran numeric,
    price_grosir numeric,
    image        varchar,
    type         varchar,
    created_at   timestamp,
    updated_at   timestamp,
    created_by   varchar,
    updated_by   varchar
);
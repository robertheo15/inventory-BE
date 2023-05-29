CREATE TABLE product_variants
(
    id         char(36) PRIMARY KEY,
    p_id       char(36),
    name       varchar,
    colour     varchar,
    created_at timestamp,
    updated_at timestamp,
    created_by varchar,
    updated_by varchar
);

ALTER TABLE product_variant ADD FOREIGN KEY (id) REFERENCES products (id);
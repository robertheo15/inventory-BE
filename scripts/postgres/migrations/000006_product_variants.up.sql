CREATE TABLE product_variants
(
    id         char(36) PRIMARY KEY,
    p_id       char(36),
    name       varchar,
    colour     varchar,
    qty     integer,
    location     varchar,
    created_at timestamp,
    updated_at timestamp,
    created_by varchar,
    updated_by varchar
);

ALTER TABLE product_variant ADD FOREIGN KEY (id) REFERENCES products (id);

create table prices
(
    id         char(36) not null
        primary key,
    eceran numeric,
    grosir numeric,
    created_at timestamp,
    updated_at timestamp,
    created_by varchar,
    updated_by varchar
);
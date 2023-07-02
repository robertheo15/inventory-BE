create table suppliers
(
    id           char(36) not null
        constraint supplier_pkey
            primary key,
    brand_name   varchar
        constraint supplier_brand_name_key
            unique,
    phone_number varchar,
    address     text,
    email        varchar
        constraint supplier_email_key
            unique,
    created_at   timestamp,
    updated_at   timestamp,
    created_by   varchar,
    updated_by   varchar
);

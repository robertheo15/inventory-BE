CREATE TABLE customers
(
    id           char(36) PRIMARY KEY,
    full_name    varchar,
    phone_number varchar,
    address      varchar,
    created_at   timestamp,
    updated_at   timestamp,
    created_by   varchar,
    updated_by   varchar
);
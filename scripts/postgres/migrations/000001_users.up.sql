CREATE TABLE users
(
    id           char(36) PRIMARY KEY,
    full_name    varchar,
    password     varchar,
    phone_number varchar,
    email        varchar UNIQUE,
    role         integer,
    active       boolean,
    created_at   timestamp,
    updated_at   timestamp,
    created_by   varchar,
    updated_by   varchar
);
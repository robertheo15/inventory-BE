CREATE TABLE transactions
(
    id             char(36) PRIMARY KEY,
    transaction_id char(36),
    c_id           char(36),
    invoice        varchar,
    status         varchar,
    type           varchar,
    created_at     timestamp,
    updated_at     timestamp,
    created_by     varchar,
    updated_by     varchar
);
ALTER TABLE transactions
    ADD FOREIGN KEY (c_id) REFERENCES customers (id);

CREATE TABLE transactions
(
    id             char(36) PRIMARY KEY,
    transaction_id char(36),
    invoice        varchar,
    status         varchar,
    type           varchar,
    total_price    numeric,
    created_at     timestamp,
    updated_at     timestamp,
    created_by     varchar,
    updated_by     varchar
);
ALTER TABLE transactions ADD FOREIGN KEY (id) REFERENCES customers (id);

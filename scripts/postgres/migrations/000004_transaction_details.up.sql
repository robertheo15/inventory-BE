CREATE TABLE transaction_details
(
    id         char(36) PRIMARY KEY,
    t_id       char(36),
    p_id       char(36),
    pv_id      char(36),
    price      numeric,
    qty        integer,
    created_at timestamp,
    updated_at timestamp,
    created_by varchar,
    updated_by varchar
);
ALTER TABLE transaction_details
    ADD FOREIGN KEY (p_id) REFERENCES products (id);
ALTER TABLE transaction_details
    ADD FOREIGN KEY (pv_id) REFERENCES product_variants (id);
ALTER TABLE transaction_details
    ADD FOREIGN KEY (id) REFERENCES transactions (id);

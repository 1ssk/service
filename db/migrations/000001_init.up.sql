CREATE SCHEMA IF NOT EXISTS schema_name;

CREATE TABLE IF NOT EXISTS schema_name.orders
(
    id    BIGSERIAL NOT NULL
        CONSTRAINT orders_pk PRIMARY KEY,
    name  TEXT      NOT NULL,
    price INTEGER   DEFAULT 0
);
DROP TABLE IF EXISTS orders;


CREATE TABLE orders (
    order_uid TEXT UNIQUE,
    data jsonb
);

CREATE TABLE IF NOT EXISTS customer
(
    id         BIGSERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name  TEXT NOT NULL,
    email      TEXT NOT NULL UNIQUE,
    address    TEXT NOT NULL,
    create_at  TIMESTAMP WITH TIME ZONE
);

INSERT INTO customer (id, first_name, last_name, email, address, create_at)
VALUES (1, 'James', 'Bond', 'james.bond@gmail.com', 'London,UK', now()),
       (2, 'Ali', 'Ahmed', 'ali.ahmed@gmail.com', 'Florida,USA', now()),
       (3, 'Jamila', 'Bah', 'jamila.bah@yahoo.com', 'Dubai,UAE', now()),
       (4, 'Ana', 'Maria', 'ana.maria@hotmail.com', 'Madrid,SPAIN', now());

CREATE TABLE IF NOT EXISTS customer_order
(
    id           BIGSERIAL PRIMARY KEY,
    customer_id  BIGINT         NOT NULL REFERENCES customer (id),
    total_amount NUMERIC(10, 2) NOT NULL,
    create_at    TIMESTAMP WITH TIME ZONE
);

INSERT INTO customer_order(customer_id, total_amount, create_at)
VALUES (1, 1.80, now()),
       (2, 0.28, now()),
       (3, 1.89, now()),
       (4, 1.90, now());

CREATE TABLE IF NOT EXISTS product
(
    id           BIGSERIAL PRIMARY KEY,
    product_name TEXT NOT NULL,
    price        NUMERIC(10, 2) NOT NULL,
    discontinued BOOLEAN NOT NULL
);

INSERT INTO product (id, product_name, price, discontinued)
VALUES (1, 'bread', 1.40, false),
       (2, 'apple', 0.40, false),
       (3, 'banana', 0.19, false),
       (4, 'milk', 0.89, false),
       (5, 'nuts', 1.90, false),
       (6, 'pepsi', 1.00, false);

CREATE TABLE IF NOT EXISTS order_item
(
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT NOT NULL REFERENCES customer_order (id),
    product_id BIGINT NOT NULL REFERENCES product (id),
    quantity   INT NOT NULL CHECK ( quantity > 0 ),
    price      NUMERIC(10, 2) NOT NULL
);

INSERT INTO order_item (order_id, product_id, quantity, price)
VALUES (1, 1, 1, 1.40),
       (1, 2, 1, 0.40),
       (2, 3, 2, .38),
       (4, 4, 1, 1.00),
       (4, 5, 1, 1.90);

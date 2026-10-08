-- S1 database practice: customer/order relational model.
-- TODO 1: add a useful constraint for nonblank customer names.
CREATE TABLE customers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE orders (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers(id),
    total_cents INTEGER NOT NULL CHECK (total_cents >= 0)
);
-- TODO 2: insert at least three customers and five orders.
-- TODO 3: write a query that reports each customer's order count and total spend.

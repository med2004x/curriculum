-- CHALLENGE STARTER: model a small shop relationally.
-- TODO 1: create customers and orders with primary/foreign keys.
-- TODO 2: insert at least three customers and five orders.
-- TODO 3: write one SELECT that proves the relationship.

CREATE TABLE customers (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers(id),
    total_cents INTEGER NOT NULL CHECK (total_cents >= 0)
);

-- CHALLENGE TODO: INSERT test data here.
-- CHALLENGE TODO: SELECT the joined result here.

-- EXERCISE STARTER: enforce invariants in a normalized schema.
-- TODO: split repeating data into tables, then add constraints that make invalid state impossible.

CREATE TABLE customer (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL
);

-- TODO: add the dependent table(s).
-- TODO: add UNIQUE, CHECK, NOT NULL, and foreign-key constraints where justified.

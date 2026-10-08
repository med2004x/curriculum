-- TODO: add a constraint that rejects negative values, then query the rows in descending order.
DROP TABLE IF EXISTS exercise_items;
CREATE TABLE exercise_items (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL, value INTEGER NOT NULL);
INSERT INTO exercise_items (name, value) VALUES ('alpha', 10), ('beta', 20), ('gamma', 30);
-- TODO: write the validation query.

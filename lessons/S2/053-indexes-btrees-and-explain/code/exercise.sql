-- EXERCISE STARTER: index + EXPLAIN.
-- TODO: identify the selective predicate, create the smallest useful index, and compare plans.

EXPLAIN (ANALYZE, BUFFERS)
SELECT id, customer_id, total_cents
FROM orders
WHERE customer_id = 42
ORDER BY id DESC
LIMIT 20;

-- TODO: CREATE INDEX ...
-- TODO: run EXPLAIN again and explain the change.

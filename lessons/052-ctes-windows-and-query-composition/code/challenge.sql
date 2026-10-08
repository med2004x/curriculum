-- CHALLENGE STARTER: CTE + window function.
-- CHALLENGE TODO: rank each customer's orders by amount and compute a running total.

WITH ranked AS (
    SELECT
        o.*,
        ROW_NUMBER() OVER (PARTITION BY customer_id ORDER BY total_cents DESC) AS rn
    FROM orders AS o
)
SELECT
    customer_id,
    id,
    total_cents,
    rn,
    SUM(total_cents) OVER (PARTITION BY customer_id ORDER BY id) AS running_total
FROM ranked
WHERE rn <= 3
ORDER BY customer_id, rn;

-- EXERCISE STARTER: joins + aggregation.
-- TODO: report each customer, order count, and total spend.
-- Keep customers with zero orders.

SELECT
    c.id,
    c.name,
    COUNT(o.id) AS order_count,
    COALESCE(SUM(o.total_cents), 0) AS total_cents
FROM customer AS c
LEFT JOIN orders AS o ON o.customer_id = c.id
GROUP BY c.id, c.name
ORDER BY c.id;

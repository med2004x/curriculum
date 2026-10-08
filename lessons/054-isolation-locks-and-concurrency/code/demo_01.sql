BEGIN;
SELECT id, balance
FROM app.accounts
WHERE id IN (1, 2)
ORDER BY id
FOR UPDATE;
-- Validate funds and update both rows in this transaction.
COMMIT;

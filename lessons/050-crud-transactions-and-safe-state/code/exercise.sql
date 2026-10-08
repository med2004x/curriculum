-- EXERCISE STARTER: implement a safe transfer transaction.
-- TODO: make the debit and credit atomic.
-- TODO: reject insufficient funds.
-- TODO: prove the total balance is conserved.

BEGIN;

-- TODO: SELECT ... FOR UPDATE the two accounts in a deterministic order.
-- TODO: perform the debit and credit.
-- TODO: verify row counts and constraints before COMMIT.

COMMIT;

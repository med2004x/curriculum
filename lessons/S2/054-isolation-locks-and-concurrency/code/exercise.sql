-- EXERCISE STARTER: concurrent sessions.
-- Session A and Session B should be run in separate psql terminals.
-- TODO: reproduce the blocking behavior, then change isolation/locking deliberately.

BEGIN;
SELECT * FROM accounts WHERE id = 1 FOR UPDATE;
-- TODO: pause here and coordinate with Session B.
COMMIT;

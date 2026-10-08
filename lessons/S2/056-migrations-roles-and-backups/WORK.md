# Lesson 056 — Migrations, Roles, Privileges, and Backups Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/001_add_description.up.sql`, `code/001_add_description.down.sql`, and `code/exercise.sql`. `code/setup.sql` is the disposable schema fixture.

## Problems

1. Write an ordered migration pair adding one backward-compatible column and state its compatibility window.

2. Define a migration role and a restricted application role with explicit grants.

3. Select a logical backup procedure and describe an isolated restore plus verification checklist.

## Run and check

Create a disposable database, then load `code/setup.sql`. Apply `code/001_add_description.up.sql` and inspect the schema; apply `code/001_add_description.down.sql` and confirm the original schema is restored. Run `code/exercise.sql` for the role and backup tasks. Use `psql -v ON_ERROR_STOP=1 -f <task-file>` for each file and record the exact commands and output.

## Optional challenge

### Migration + restore drill

Create forward and backward schema changes, roles with least privilege, and a backup/restore procedure. Prove a restored database satisfies a checksum or row-count invariant.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

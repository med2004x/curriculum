# Lesson 056 — Problem Set

**Topic:** Migrations, Roles, Privileges, and Backups.
**Your files:** `code/001_add_description.up.sql`, `code/001_add_description.down.sql`, and `code/exercise.sql`. `code/setup.sql` is the disposable schema fixture. `demo_*` files are teaching examples; implement the assigned task files.

## Problems

1. Write an ordered migration pair adding one backward-compatible column and state its compatibility window.

2. Define a migration role and a restricted application role with explicit grants.

3. Select a logical backup procedure and describe an isolated restore plus verification checklist.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Create a disposable database, then load `code/setup.sql`. Apply `code/001_add_description.up.sql` and inspect the schema; apply `code/001_add_description.down.sql` and confirm the original schema is restored. Run `code/exercise.sql` for the role and backup tasks. Use `psql -v ON_ERROR_STOP=1 -f <task-file>` for each file and record the exact commands and output.

## Stretch problem

### Migration + restore drill

Create forward and backward schema changes, roles with least privilege, and a backup/restore procedure. Prove a restored database satisfies a checksum or row-count invariant.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

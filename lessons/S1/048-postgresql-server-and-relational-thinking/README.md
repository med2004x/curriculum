# Lesson 048 — PostgreSQL, psql, and Relational Thinking

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Learn where database state lives and how a client addresses it, then interpret tables and SQL queries as operations over related data rather than as unstructured spreadsheets.

**The contract.** PostgreSQL is a database server; `psql` is one client. A server can host multiple databases, a database contains schemas, schemas contain objects such as tables, and a connection operates against one database under a role. A table has named columns and rows; keys and constraints define which states are valid.

**Worked example.** Use the assigned `psql` commands to identify the current connection, database, schema, tables, and roles. Then query a small `app.tasks` table with `SELECT id, title, status FROM app.tasks;` and add a `WHERE` predicate to choose only one status.

## 2. Mechanism and vocabulary

**Mechanism.** `SELECT` chooses output expressions, `FROM` names input relations, and `WHERE` filters rows. SQL commonly uses bag (multiset) semantics, so duplicate rows can remain unless `DISTINCT` is requested. A NULL comparison is unknown rather than true or false; use `IS NULL` to test for nullness.

**Second point.** A relational model describes relations by attributes and tuples; SQL tables are a practical implementation with constraints and additional behavior. Do not assume row order without `ORDER BY`.

`\conninfo` describes the current client connection; `\l`, `\dn`, `\dt`, and `\du` list databases, schemas, tables, and roles respectively. These psql meta-commands are not SQL statements.

**Worked example.** Compare `SELECT status FROM app.tasks;` with `SELECT DISTINCT status FROM app.tasks;`. Explain why the first can return repeated status values and why neither query promises row order without an `ORDER BY` clause.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
\conninfo
\l
\dn
\dt
\du
SELECT current_database(), current_user, version();
\q
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Record the current database and role, list the target schema/table, then run a query with and without `DISTINCT`. Add an `ORDER BY` and explain the guarantee it supplies.

**Failure mode.** Connect to a nonexistent database or query a nonexistent table. Distinguish connection failure from a valid query returning zero rows.

**Completion criterion.** You can locate the current database objects, write and interpret basic SELECT queries, explain duplicate and NULL behavior, and distinguish SQL from psql commands.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

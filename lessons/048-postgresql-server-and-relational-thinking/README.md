# Lesson 048 — PostgreSQL, psql, and Relational Thinking

## Threads
**Primary:** Databases/Go

**Supporting:** server/client; tables; relational model

**Math companion:** sets, relations, predicates

**Time:** 3–4 hours. **Prerequisites:** Lesson 047.

**Time:** 3–4 hours. **Prerequisites:** Lesson 047.

## What you will be able to do

1. Distinguish PostgreSQL server, database, schema, table, and client connection.
2. Use `psql` to inspect connection identity, databases, schemas, tables, and roles.
3. Explain rows and columns as a relation's current tabular representation and write basic selection/projection queries.
4. Explain why SQL query results can contain duplicates and why NULL requires three-valued logic.

## Retrieve first (no notes)

1. What does a packet capture show that application logs may not?
2. What is one caveat when inferring wire packet boundaries from a capture?
3. What is the difference between a network-level retransmission and an application retry?
4. Why should a latency measurement identify its start and end events?

## 1. Problem and model

**What problem this solves.** Learn where database state lives and how a client addresses it, then interpret tables and SQL queries as operations over related data rather than as unstructured spreadsheets.

**The contract.** PostgreSQL is a database server; `psql` is one client. A server can host multiple databases, a database contains schemas, schemas contain objects such as tables, and a connection operates against one database under a role. A table has named columns and rows; keys and constraints define which states are valid.

**Worked example.** Use the assigned `psql` commands to identify the current connection, database, schema, tables, and roles. Then query a small `app.tasks` table with `SELECT id, title, status FROM app.tasks;` and add a `WHERE` predicate to choose only one status.

Self-check: which values come from the connection context and which are rows returned by SQL?

## 2. Mechanism and vocabulary

**Mechanism.** `SELECT` chooses output expressions, `FROM` names input relations, and `WHERE` filters rows. SQL commonly uses bag (multiset) semantics, so duplicate rows can remain unless `DISTINCT` is requested. A NULL comparison is unknown rather than true or false; use `IS NULL` to test for nullness.

**Second point.** A relational model describes relations by attributes and tuples; SQL tables are a practical implementation with constraints and additional behavior. Do not assume row order without `ORDER BY`.

`\conninfo` describes the current client connection; `\l`, `\dn`, `\dt`, and `\du` list databases, schemas, tables, and roles respectively. These psql meta-commands are not SQL statements.

**Worked example.** Compare `SELECT status FROM app.tasks;` with `SELECT DISTINCT status FROM app.tasks;`. Explain why the first can return repeated status values and why neither query promises row order without an `ORDER BY` clause.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Record the current database and role, list the target schema/table, then run a query with and without `DISTINCT`. Add an `ORDER BY` and explain the guarantee it supplies.

**Failure mode.** Connect to a nonexistent database or query a nonexistent table. Distinguish connection failure from a valid query returning zero rows.

**Completion criterion.** You can locate the current database objects, write and interpret basic SELECT queries, explain duplicate and NULL behavior, and distinguish SQL from psql commands.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict which database and role the current connection uses.
2. Predict how many rows `DISTINCT` can remove from repeated status data.
3. Predict whether a query without `ORDER BY` has a guaranteed row order.
4. Predict whether `status = NULL` selects rows whose status is null.
5. Predict which psql command lists schemas and which lists tables.
6. Distinguish an empty result set from a missing-table error.
7. Predict the effect of adding `WHERE status = 'open'`.
8. State whether each command is SQL or a psql meta-command.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Connect with `psql` and inspect connection identity, database, schema, tables, and roles.

2. Query `app.tasks` for selected columns and filter rows with a `WHERE` predicate.

3. Compare duplicates, NULL predicates, and result ordering with and without the relevant SQL clauses.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Construct a relation that is reflexive and symmetric but *not* transitive — for example, "person A and person B have met in person" on a set of people, where A met B, B met C, but A never met C.

2. Confirm explicitly that this relation fails the transitivity check with your constructed example, and explain why treating this kind of relation as if it defined clean equivalence classes (e.g. grouping "people who've met" into isolated clusters) would silently produce wrong groupings — someone in cluster {A,B} and someone in cluster {B,C} might get incorrec.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 048 — PostgreSQL, psql, and Relational Thinking: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Distinguish a PostgreSQL server, database, schema, and table.
2. What does `\conninfo` report?
3. How do `WHERE` and `DISTINCT` affect query results?
4. Why is row order unspecified without `ORDER BY`?
5. Why does `status = NULL` not test for null values?
6. How does a missing-table error differ from a query returning no rows?

## Evidence to record

Create `lesson048/EVIDENCE.md` with exactly this structure:

```text
# Lesson 048 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating psql meta-commands as SQL statements.
- Assuming a table has a stable row order without `ORDER BY`.
- Assuming SQL always removes duplicate rows.
- Comparing NULL with `=` rather than using `IS NULL`.
- Confusing an empty result set with a connection or relation error.

## Precision checks

1. Identify the current database, schema, and role.
2. Explain the columns and rows selected by a query.
3. State when duplicates remain and how `DISTINCT` changes that.
4. Explain NULL predicate behavior and ordering guarantees.
5. Distinguish psql meta-commands from SQL.

Relational thinking starts with identifying the relation and predicate; the client and server context determine where the query runs.
## University revision bridge

For the LIG track, connect relational tables and predicates to set/logic foundations, while accounting for SQL's duplicate and NULL semantics.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 048 — PostgreSQL, psql, and Relational Thinking without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can prove a set identity from definitions, correctly verify whether a given relation is an equivalence relation or partial order by checking each required property explicitly, and you can identify a real (or plausible) case in your own systems where a "similarity" relation might silently fail to be transitive. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

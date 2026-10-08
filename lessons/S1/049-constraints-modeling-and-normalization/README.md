# Lesson 049 — Constraints, Modeling, and Normalization

## Threads
**Primary:** Databases/Design

**Supporting:** keys; constraints; relationships; normalization

**Math companion:** relations and functions

**Time:** 3–4 hours. **Prerequisites:** Lesson 048.

**Time:** 3–4 hours. **Prerequisites:** Lesson 048.

## What you will be able to do

1. Translate entity and relationship requirements into tables, primary keys, and foreign keys.
2. Enforce domain rules with `NOT NULL`, `UNIQUE`, `CHECK`, and referential constraints.
3. Model many-to-many relationships with a junction table such as `task_tags(task_id, tag_id)`.
4. Identify an update anomaly and normalize the schema without losing required facts.

## Retrieve first (no notes)

1. Why does SQL not guarantee row order without `ORDER BY`?
2. How do `WHERE` and `DISTINCT` affect query results?
3. How should SQL test a NULL value?
4. What is the difference between a schema and a database?

## 1. Problem and model

**What problem this solves.** A relational schema should represent entities and relationships without allowing invalid states or forcing repeated facts to be updated in many rows.

**The contract.** Give each task and tag a stable key. A task can have many tags and a tag can belong to many tasks, so represent the association with a junction table whose foreign keys reference both parent tables. Use a composite primary key `(task_id, tag_id)` to prevent duplicate associations.

**Worked example.** Model `tasks(id, title, status)`, `tags(id, name UNIQUE)`, and `task_tags(task_id REFERENCES tasks, tag_id REFERENCES tags, PRIMARY KEY(task_id, tag_id))`. This stores each tag name once and prevents links to nonexistent tasks or tags.

Self-check: where should a rule live if every client must obey it? Put core invariants in database constraints as well as application validation where appropriate.

## 2. Mechanism and vocabulary

**Mechanism.** A primary key identifies a row; a foreign key enforces a reference to a candidate key; `UNIQUE` prevents duplicate values; `NOT NULL` excludes missing values; `CHECK` enforces a row predicate. Functional dependencies express which attributes determine others and help identify repeated facts and update anomalies.

**Second point.** Normalization decomposes relations to reduce redundancy while preserving information and dependencies where possible. A practical 3NF design avoids storing a non-key fact that depends transitively on a key through another non-key attribute; normalization is not a substitute for understanding query and workload requirements.

Do not store a comma-separated tag list in `tasks` if tags must be queried, constrained, or shared. A junction table gives each association a row that can be indexed and constrained.

**Worked example.** Add one boundary case to the # 049 — Constraints, Modeling, and Normalization exercise.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
tags
task_tags(task_id, tag_id)
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Attempt a duplicate tag name, duplicate task/tag association, nonexistent foreign key, and invalid status. Confirm each is rejected by the intended constraint. Then query a task's tags through the junction table.

**Failure mode.** Store repeated tag names in every task row, rename one tag in only some rows, and identify the resulting update anomaly. Replace the repeated text with a tag key and association table.

**Completion criterion.** You can explain the key and relationship choices, show constraints reject invalid data, and justify a normalization step by the anomaly it prevents.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether a duplicate tag name satisfies `UNIQUE`.
2. Predict what happens when a task/tag link references a missing task.
3. Predict whether the same task/tag pair can be inserted twice.
4. Predict how a many-to-many association is represented.
5. Identify the update anomaly caused by repeating tag names.
6. State what a primary key guarantees.
7. Distinguish a foreign-key constraint from application-side validation.
8. Identify which functional dependency motivates the decomposition.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Define task, tag, and task-tag association tables with primary and foreign keys.

2. Add `UNIQUE`, `NOT NULL`, and `CHECK` constraints that encode the domain rules.

3. Demonstrate invalid inserts and query the many-to-many relationship.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one rule: permit a task with no status, introduce tag categories, or allow duplicate tag names in different scopes. Explain which key/constraint changes and which data anomalies become possible.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does a primary key identify?
2. What does a foreign key enforce?
3. How does a junction table model many-to-many relationships?
4. What does `UNIQUE` prevent?
5. What is an update anomaly?
6. How does normalization reduce repeated facts?

## Evidence to record

Create `lesson049/EVIDENCE.md` with exactly this structure:

```text
# Lesson 049 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Storing a many-to-many relationship as comma-separated text.
- Relying only on application validation for invariants that the database can enforce.
- Omitting a foreign key and allowing orphan associations.
- Repeating a fact across rows and creating inconsistent updates.
- Normalizing mechanically without considering the required query model.

## Precision checks

1. Draw the table relationships and identify primary/foreign keys.
2. State which constraints enforce each domain invariant.
3. Explain how the junction table prevents duplicate links.
4. Identify an update anomaly and show how decomposition removes it.
5. State one workload trade-off the schema still needs to support.

Schema design is a model of valid states: keys and constraints should make invalid relationships difficult to store.
## University revision bridge

For the LIG track, connect functional dependencies and normalization to concrete keys, constraints, and relationship tables.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 049 — Constraints, Modeling, and Normalization without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can reproduce the main behavior and diagnose one failure from evidence. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

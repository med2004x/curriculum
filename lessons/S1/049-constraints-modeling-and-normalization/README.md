# Lesson 049 — Constraints, Modeling, and Normalization

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** A relational schema should represent entities and relationships without allowing invalid states or forcing repeated facts to be updated in many rows.

**The contract.** Give each task and tag a stable key. A task can have many tags and a tag can belong to many tasks, so represent the association with a junction table whose foreign keys reference both parent tables. Use a composite primary key `(task_id, tag_id)` to prevent duplicate associations.

**Worked example.** Model `tasks(id, title, status)`, `tags(id, name UNIQUE)`, and `task_tags(task_id REFERENCES tasks, tag_id REFERENCES tags, PRIMARY KEY(task_id, tag_id))`. This stores each tag name once and prevents links to nonexistent tasks or tags.

## 2. Mechanism and vocabulary

**Mechanism.** A primary key identifies a row; a foreign key enforces a reference to a candidate key; `UNIQUE` prevents duplicate values; `NOT NULL` excludes missing values; `CHECK` enforces a row predicate. Functional dependencies express which attributes determine others and help identify repeated facts and update anomalies.

**Second point.** Normalization decomposes relations to reduce redundancy while preserving information and dependencies where possible. A practical 3NF design avoids storing a non-key fact that depends transitively on a key through another non-key attribute; normalization is not a substitute for understanding query and workload requirements.

Do not store a comma-separated tag list in `tasks` if tags must be queried, constrained, or shared. A junction table gives each association a row that can be indexed and constrained.

**Worked example.** Add one boundary case to the # 049 — Constraints, Modeling, and Normalization exercise.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
tags
task_tags(task_id, tag_id)
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Attempt a duplicate tag name, duplicate task/tag association, nonexistent foreign key, and invalid status. Confirm each is rejected by the intended constraint. Then query a task's tags through the junction table.

**Failure mode.** Store repeated tag names in every task row, rename one tag in only some rows, and identify the resulting update anomaly. Replace the repeated text with a tag key and association table.

**Completion criterion.** You can explain the key and relationship choices, show constraints reject invalid data, and justify a normalization step by the anomaly it prevents.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

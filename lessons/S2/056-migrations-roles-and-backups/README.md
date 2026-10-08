# Lesson 056 — Migrations, Roles, Privileges, and Backups

## Threads
**Primary:** Databases/Operations/Security

**Supporting:** schema changes; SQL privileges; restore thinking

**Math companion:** reliability basics

**Time:** 3–4 hours. **Prerequisites:** Lesson 055.

**Time:** 3–4 hours. **Prerequisites:** Lesson 055.

## What you will be able to do

1. Apply ordered schema migrations and explain how migration state is tracked.
2. Design PostgreSQL login/application roles with least-privilege grants.
3. Distinguish a logical dump from a physical/base backup at a high level.
4. Restore a backup into an isolated target and verify data/schema, rather than treating dump creation as proof of recoverability.

## Retrieve first (no notes)

1. Why pass a context to a database operation?
2. How should a repository represent a missing row?
3. Why must transaction commit errors be checked?
4. What makes a test an integration test rather than a fake-backed unit test?

## 1. Problem and model

**What problem this solves.** Schema changes, access control, and backup recovery are operational responsibilities: production data must evolve safely, only authorized roles should access it, and a backup is useful only if it can be restored.

**The contract.** Each migration has a unique ordered version and an explicit change. Deployment must account for application versions that may run before, during, and after the migration. Runtime roles receive only required privileges; schema ownership and migration privileges remain separate. A backup plan states what is backed up, where it is stored, and how restoration is verified.

**Worked example.** Add a nullable `description` column through migration `001`, deploy code that can work with both old and new schema, then backfill and tighten constraints in a later migration if required. Do not make an old application fail immediately by removing or renaming a column during the same compatibility window.

Self-check: distinguish migration rollback capability from data recovery. Reversing DDL cannot necessarily reconstruct data discarded by the forward migration.

## 2. Mechanism and vocabulary

**Mechanism.** Use an expand-migrate-contract sequence for changes requiring compatibility: add compatible schema, deploy code that supports both states, migrate/backfill data, then remove obsolete schema only after old code is gone. Migration tools track applied versions so a migration is not accidentally applied twice.

**Second point.** Separate an owner/migration role from the application login role. Grant only the required schema usage and table/sequence operations. Prefer revoking broad defaults and granting explicit privileges rather than running the service as a superuser.

Logical dumps (for example, `pg_dump`) represent database objects/data for logical restoration. Physical/base backups capture database files and WAL-related state using PostgreSQL's supported backup procedures. They have different restore requirements; a backup is not validated until a restore procedure succeeds.

**Worked example.** Create a migration-only role that owns or alters schema objects and a runtime role that can connect and perform only the application's required CRUD operations. Verify that the runtime role cannot create roles, databases, or unrelated objects.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
migrations/
    001_add_description.up.sql
    001_add_description.down.sql
compatibility-matrix.md
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Record migration version before/after, objects changed, role grants, backup method, restore target, and post-restore checks such as expected schema and representative row counts. Never restore over production as a verification shortcut.

**Failure mode.** Attempt an unauthorized DDL operation as the runtime role and confirm least-privilege rejection. Separately restore a dump to an empty isolated database and identify any missing extension, role, or configuration prerequisite.

**Completion criterion.** You can explain migration ordering, role boundaries, backup formats at a high level, and a restore verification plan.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether a migration should run again after its version is recorded as applied.
2. Predict which application versions remain compatible with an added nullable column.
3. Predict which operation should fail for a least-privilege runtime role.
4. Predict what extra state a physical backup restore may require.
5. Predict which objects may be omitted from a database-only logical dump.
6. Predict why a successful backup command does not prove a successful restore.
7. Predict what is lost if a destructive migration is reversed without a data-preserving plan.
8. State the isolated restore target and checks before relying on the backup.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Write an ordered migration pair adding one backward-compatible column and state its compatibility window.

2. Define a migration role and a restricted application role with explicit grants.

3. Select a logical backup procedure and describe an isolated restore plus verification checklist.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one operational assumption: add a non-null constraint, rotate the runtime role, or restore to a different PostgreSQL instance. Explain the data migration, privilege, or compatibility impact.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does a migration version record?
2. Why use expand-migrate-contract for compatible schema changes?
3. Why separate database ownership from application runtime privileges?
4. How does a logical dump differ from a physical backup?
5. Why must restoration be tested on an isolated target?
6. What does least privilege require for the runtime role?

## Evidence to record

Create `lesson056/EVIDENCE.md` with exactly this structure:

```text
# Lesson 056 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Applying a destructive schema change while old application code still runs.
- Running the production service with the migration owner or superuser role.
- Assuming a down migration can restore deleted data.
- Treating creation of a dump as proof of recoverability.
- Restoring a backup without checking extensions, roles, and configuration prerequisites.

## Precision checks

1. Explain the migration order and supported application versions.
2. Show the runtime role's allowed operations and a denied privileged operation.
3. State the selected backup type and restore prerequisites.
4. Describe restore verification against an isolated database.
5. Explain what data a rollback cannot reconstruct.

Operational safety is demonstrated by compatibility, least privilege, and a verified recovery path, not by migration files alone.
## University revision bridge

For the LIG track, connect schema evolution to compatibility contracts and database roles to security boundaries, then document a reproducible restore procedure.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 056 — Migrations, Roles, Privileges, and Backups without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can reproduce the main behavior and diagnose one failure from evidence. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

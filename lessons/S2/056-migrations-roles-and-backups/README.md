# Lesson 056 — Migrations, Roles, Privileges, and Backups

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Schema changes, access control, and backup recovery are operational responsibilities: production data must evolve safely, only authorized roles should access it, and a backup is useful only if it can be restored.

**The contract.** Each migration has a unique ordered version and an explicit change. Deployment must account for application versions that may run before, during, and after the migration. Runtime roles receive only required privileges; schema ownership and migration privileges remain separate. A backup plan states what is backed up, where it is stored, and how restoration is verified.

**Worked example.** Add a nullable `description` column through migration `001`, deploy code that can work with both old and new schema, then backfill and tighten constraints in a later migration if required. Do not make an old application fail immediately by removing or renaming a column during the same compatibility window.

## 2. Mechanism and vocabulary

**Mechanism.** Use an expand-migrate-contract sequence for changes requiring compatibility: add compatible schema, deploy code that supports both states, migrate/backfill data, then remove obsolete schema only after old code is gone. Migration tools track applied versions so a migration is not accidentally applied twice.

**Second point.** Separate an owner/migration role from the application login role. Grant only the required schema usage and table/sequence operations. Prefer revoking broad defaults and granting explicit privileges rather than running the service as a superuser.

Logical dumps (for example, `pg_dump`) represent database objects/data for logical restoration. Physical/base backups capture database files and WAL-related state using PostgreSQL's supported backup procedures. They have different restore requirements; a backup is not validated until a restore procedure succeeds.

**Worked example.** Create a migration-only role that owns or alters schema objects and a runtime role that can connect and perform only the application's required CRUD operations. Verify that the runtime role cannot create roles, databases, or unrelated objects.

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

## 4. Verification, failure, and evidence

**Verification.** Record migration version before/after, objects changed, role grants, backup method, restore target, and post-restore checks such as expected schema and representative row counts. Never restore over production as a verification shortcut.

**Failure mode.** Attempt an unauthorized DDL operation as the runtime role and confirm least-privilege rejection. Separately restore a dump to an empty isolated database and identify any missing extension, role, or configuration prerequisite.

**Completion criterion.** You can explain migration ordering, role boundaries, backup formats at a high level, and a restore verification plan.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

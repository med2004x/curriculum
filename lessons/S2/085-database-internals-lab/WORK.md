# Lesson 085 — Problem Set

**Topic:** Database Internals Lab.
**Your files:** `code/exercise.c`, `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: If using BusTub directly, complete at least the buffer pool manager project and the B+ tree index project from the official course materials (`https://github.com/cmu-db/bustub`), using its existing test infrastructure to verify correctness — this is the recommended default path, since BusTub's tests are considerably more rigorous than what's practical to con.

2. Extend it with one explicit boundary case: If building your own integrated system instead (a reasonable alternative if you want the components to be ones you fully wrote yourself, end to end), integrate your own Lessons 1-6 implementations into one coherent system with this minimum scope:

3. Rebuild the same behavior with one changed constraint: **Storage layer**: your Lesson 1 page format, with records actually read/written through your Lesson 2 buffer pool (not bypassing it) — every page access in the rest of the system should go through `fetchPage`/`unpinPage`, with correct pin discipline throughout.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### CMU BusTub project

Select the matching BusTub project on pages, buffer pools, B+ trees, query execution, or transactions from the semester that matches the local upstream checkout. Implement the official project API and use the supplied tests. `exercise.c` is a concept prompt only; it is not a miniature BusTub substitute.

## Upstream lab: CMU BusTub

Use the official [BusTub project](https://github.com/cmu-db/bustub) and its [matching semester project guide](https://15445.courses.cs.cmu.edu/). APIs and tests vary by release. Follow that handout's CMake workflow, then run `make check-tests` from the build directory (or the target prescribed for that release). Record the semester/commit, target, output, and one invariant in `EVIDENCE.md`. Keep solutions in a private working copy and follow the course FAQ and academic-integrity rules. The local C/SQL starters are prompts, not a database engine or replacement harness.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

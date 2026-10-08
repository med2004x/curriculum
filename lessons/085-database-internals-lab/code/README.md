# Code map — Lesson 085 — Database Internals Lab

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Database internals lab

Combine pages, indexes, transactions, WAL, and recovery into one small simulation. Write invariants and a failure test for each subsystem.

### Acceptance criteria

- Start from a clean workspace or the latest lesson project without copying a completed solution.
- Implement the required behavior end-to-end.
- Include automated tests where the technology allows it, plus one deliberate failure test.
- Record design assumptions and the main invariant(s) in Markdown.
- Demonstrate the project from a reproducible command sequence.
- Finish with a short post-mortem: what failed, why, and what you changed.

## 8. Debugging and failure lab

1. Take your integrated system (or the BusTub project you completed) and deliberately break one cross-component assumption — for example, if you built your own system, bypass the buffer pool's pin discipline in exactly one code path (directly access a page without pinning it first) while the rest of the system correctly uses pins.

2. Run your workload under concurrent access (multiple goroutines/threads performing operations simultaneously) and attempt to reproduce a corruption or crash caused specifically by this one broken assumption — e.g. a page being evicted while the unprotected code path is still reading it, per Lesson 2's pin/unpin discussion.

3. If you cannot reliably reproduce a visible failure (this specific race can be timing-dependent and hard to force), document your reasoning for why the assumption violation is still a genuine bug even if your specific test run didn't happen to trigger a visible symptom — directly echoing C track Lesson 5's core point about undefined behavior not requiring a v.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the s

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.

## Upstream lab: CMU BusTub

Use the official BusTub project and its project-specific tests for the database component you select. The local C file is an isolated prompt, not a buffer pool, B+ tree, or database engine implementation.

- Upstream source: https://github.com/cmu-db/bustub
- Local audit checkout: `../../../../dart-upstreams/bustub`
- Course project guide: https://15445.courses.cs.cmu.edu/ (select the matching semester; project APIs and tests vary by release).

Follow the matching handout, build with its documented CMake workflow, then run `make check-tests` from the BusTub build directory (or the project-specific target prescribed by that semester). Record the semester/commit, target, test output, and one invariant in `EVIDENCE.md`. Do not recreate the BusTub harness inside DART.

Keep your BusTub solution in a private working copy and follow the course FAQ and academic-integrity rules. The upstream README warns students not to fork or publish solutions; use only the supported local/autograder workflow for your course release.

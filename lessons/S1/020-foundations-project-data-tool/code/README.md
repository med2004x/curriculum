# Code map — Lesson 020 — Foundations Project: In-Memory Data Tool

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Data tool project

Extend the earlier in-memory data tool with filtering, sorting, aggregation, persistence format, and a repeatable regression test suite.

### Acceptance criteria

- Start from a clean workspace or the latest lesson project without copying a completed solution.
- Implement the required behavior end-to-end.
- Include automated tests where the technology allows it, plus one deliberate failure test.
- Record design assumptions and the main invariant(s) in Markdown.
- Demonstrate the project from a reproducible command sequence.
- Finish with a short post-mortem: what failed, why, and what you changed.

## 8. Debugging and failure lab

1. Supply a line missing its second delimiter and verify that no record is appended.

2. Supply a line at the exact maximum field length, then one byte beyond it; check that the chosen boundary policy is consistent.

3. Search an empty file and a file with records; distinguish “no match” from “failed to open.”

4. Plant a parsing or search bug, reduce it to the smallest input that demonstrates it, then explain the correction and evidence.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint: change the delimiter format, allow quoted messages, or replace linear search with sorted-index search. Explain which parser or ordering invariant changes and add a test that would fail under the original assumption.

Your transfer answer is complete when the changed rule is explicit and the test demonstrates the difference.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Closed-book definition:** State the key formal def

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.

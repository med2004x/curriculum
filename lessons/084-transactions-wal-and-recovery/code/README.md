# Code map — Lesson 084 — Transactions, WAL, and Recovery

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### WAL + recovery simulation

Model a sequence of page changes and WAL records, simulate a crash at two points, and determine the recovered state under each point.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Take your step 2 fix (`REPEATABLE READ` preventing the lost update via a commit-time error) and confirm your *application code* actually needs to handle that error explicitly — rerun the scenario but have the "losing" transaction's code simply ignore the serialization error rather than catching it and retrying.

2. Confirm the increment is effectively lost anyway, not because the database allowed the anomaly, but because the application discarded the database's correct rejection instead of retrying the transaction.

3. Explain why choosing a stronger isolation level is necessary but not sufficient on its own — the database can correctly detect and reject the unsafe interleaving, but the application must still contain retry logic for the rejection to actually result in correct final behavior, rather than an error that's silently swallowed and produces the exact same lost up.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one durability setting or crash point in the model. Explain which ack

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.

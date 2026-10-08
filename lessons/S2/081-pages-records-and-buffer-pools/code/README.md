# Code map — Lesson 081 — Pages, Records, and Buffer Pools

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `check.c` | TEST | Checks serialization bytes and rejects an out-of-bounds write without changing the page. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Page layout model

Model database pages and records in memory. Simulate inserts, deletes, and page splits while tracking free space and record locations.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Insert records until the page is nearly full.
2. Delete records from the middle without compacting.
3. Attempt an insertion that fits in total free bytes but not in a contiguous gap.
4. Compact live records, update all slot offsets, and verify every value again.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change the record encoding or page size. Recalculate offsets and show which page invariants remain unchanged.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Model:** State the system model and assumptions before solving.
2. **Trace:** Work a small failure or state-transition example by hand.
3. **Trade-off:** Compare two designs and state the condition under which each wins.
4. **Evidence:** Identify the measurement, test, or trace that would distinguish the designs in practice.

## 11. Self-test

- [ ] What fields does a slot entry need for variable-length records?
- [ ] What prevents slot and record storage from overlapping?
- [ ] What does a pin count protect against?
- [ ] What does a dirty bit indicate?
- [ ] Why can free-space fragmentati

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.


## Automated check

`check.c` exercises record serialization and verifies that an out-of-bounds write is rejected without modifying the page. Run it with `gcc -std=c17 -Wall -Wextra -Wpedantic check.c -o check && ./check`. The test is expected to fail against the untouched TODO starter.

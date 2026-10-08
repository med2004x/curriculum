# Lesson 097 — Problem Set

**Topic:** Rust Ownership, Lifetimes, and Systems Thinking.
**Your files:** `code/exercise.rs`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).

2. Extend it with one explicit boundary case: Confirm the Rust version, using the standard library type, requires no manual `free`/`malloc` equivalent at all — ownership and the `Drop` trait handle cleanup automatically.

3. Rebuild the same behavior with one changed constraint: Deliberately reproduce C track Lesson 4's use-after-free bug's *intent* in Rust (attempt to use a value after it's been moved elsewhere) and confirm the Rust compiler refuses to compile it, with an error message explicitly citing the move — read the actual compiler error text and confirm it correctly identifies both the move point and the invalid subsequent.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Use the exact command in the task specification and record it in `EVIDENCE.md`.

## Stretch problem

### Rust systems exercise

Implement a small resource-owning Rust type and deliberately trigger an ownership/borrow error. Fix it while explaining the lifetime relationship.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

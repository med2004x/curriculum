# Lesson 097 — Rust Ownership, Lifetimes, and Systems Thinking Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.rs`.

## Problems

1. Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).

2. Confirm the Rust version, using the standard library type, requires no manual `free`/`malloc` equivalent at all — ownership and the `Drop` trait handle cleanup automatically.

3. Deliberately reproduce C track Lesson 4's use-after-free bug's *intent* in Rust (attempt to use a value after it's been moved elsewhere) and confirm the Rust compiler refuses to compile it, with an error message explicitly citing the move — read the actual compiler error text and confirm it correctly identifies both the move point and the invalid subsequent.

## Run and check

Use the exact command in the task specification and record it in `EVIDENCE.md`.

## Optional challenge

### Rust systems exercise

Implement a small resource-owning Rust type and deliberately trigger an ownership/borrow error. Fix it while explaining the lifetime relationship.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

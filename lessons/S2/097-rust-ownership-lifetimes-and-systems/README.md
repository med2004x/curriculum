# Lesson 097 — Rust Ownership, Lifetimes, and Systems Thinking

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand Rust's ownership system as a compile-time solution to exactly the memory-safety problems C track Lessons 2-4 covered manually — translate real, unsafe C patterns into safe Rust, and learn to read and fix borrow-checker errors by understanding the rule being enforced, not by trial-and-error editing.

**The contract.** **The core idea: exactly one owner, enforced by the compiler, not by convention.** C track Lesson 4 established that ownership in C is a discipline you maintain manually — nothing stops two pointers from both believing they own the same allocation, leading to double-frees or use-after-free.

**Worked example.** Start with the smallest concrete case described by the source material: Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).

## 2. Mechanism and vocabulary

**Mechanism.** When the owner leaves scope, Rust runs `Drop` and releases owned resources. Moving transfers ownership; borrowing grants temporary access. Lifetimes constrain references so they cannot outlive their referents, while shared and mutable borrows enforce aliasing rules at compile time.

**Second point.** **Move semantics: what happens when ownership transfers.** Assigning a value to a new variable, or passing it to a function, *moves* ownership by default (for types that don't implement `Copy`) — the original variable becomes invalid, and the compiler will refuse to compile any subsequent use of it.

**Worked example.** Confirm the Rust version, using the standard library type, requires no manual `free`/`malloc` equivalent at all — ownership and the `Drop` trait handle cleanup automatically.

## 3. Worked case from the reference material

Worked example: Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).

## 4. Verification, failure, and evidence

**Verification.** Show your actual Rust compiler error messages for steps 2-3 (not paraphrased — the real `rustc` output, which includes specific, often genuinely helpful guidance), and your corrected code for step 4 alongside a one-sentence explanation of which specific ownership rule each fix satisfies.

**Failure mode.** Replace a structurally correct ownership fix with an unnecessary `.clone()`. Show that it compiles, then mutate one copy and inspect the other: cloning creates independent data and may change semantics or cost. Prefer a move, a scoped borrow, or an explicitly owned return type when that matches the contract.

**Completion criterion.** You can explain value ownership and reference scopes, reproduce and interpret move and borrow errors, and choose an owned or borrowed return type that fits the caller?s lifetime.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

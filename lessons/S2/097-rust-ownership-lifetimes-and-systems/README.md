# Lesson 097 — Rust Ownership, Lifetimes, and Systems Thinking

## Threads
**Primary:** Rust/C

**Supporting:** ownership; lifetimes; traits; transfer from C

**Math companion:** type reasoning

**Time:** 3–4 hours. **Prerequisites:** Lesson 096.


## What you will be able to do

1. Translate a small C data structure into idiomatic Rust ownership.
2. Predict moves and borrows from Rust?s aliasing rules.
3. Choose owned values or borrowed views for a caller contract.
4. Use tests for behavior beyond compile-time memory checks.

## Retrieve first (no notes)

1. What makes a workload representative?
2. How does contiguous layout aid a scan?
3. Why must an optimization preserve behavior?
4. Which tool can reveal allocation cost?

## 1. Problem and model

**What problem this solves.** Understand Rust's ownership system as a compile-time solution to exactly the memory-safety problems C track Lessons 2-4 covered manually — translate real, unsafe C patterns into safe Rust, and learn to read and fix borrow-checker errors by understanding the rule being enforced, not by trial-and-error editing.

**The contract.** **The core idea: exactly one owner, enforced by the compiler, not by convention.** C track Lesson 4 established that ownership in C is a discipline you maintain manually — nothing stops two pointers from both believing they own the same allocation, leading to double-frees or use-after-free.

**Worked example.** Start with the smallest concrete case described by the source material: Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** When the owner leaves scope, Rust runs `Drop` and releases owned resources. Moving transfers ownership; borrowing grants temporary access. Lifetimes constrain references so they cannot outlive their referents, while shared and mutable borrows enforce aliasing rules at compile time.

**Second point.** **Move semantics: what happens when ownership transfers.** Assigning a value to a new variable, or passing it to a function, *moves* ownership by default (for types that don't implement `Copy`) — the original variable becomes invalid, and the compiler will refuse to compile any subsequent use of it.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Confirm the Rust version, using the standard library type, requires no manual `free`/`malloc` equivalent at all — ownership and the `Drop` trait handle cleanup automatically.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Show your actual Rust compiler error messages for steps 2-3 (not paraphrased — the real `rustc` output, which includes specific, often genuinely helpful guidance), and your corrected code for step 4 alongside a one-sentence explanation of which specific ownership rule each fix satisfies.

**Failure mode.** Replace a structurally correct ownership fix with an unnecessary `.clone()`. Show that it compiles, then mutate one copy and inspect the other: cloning creates independent data and may change semantics or cost. Prefer a move, a scoped borrow, or an explicitly owned return type when that matches the contract.

**Completion criterion.** You can explain value ownership and reference scopes, reproduce and interpret move and borrow errors, and choose an owned or borrowed return type that fits the caller?s lifetime.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 097 — Rust Ownership, Lifetimes, and Systems Thinking case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).

2. Extend it with one explicit boundary case: Confirm the Rust version, using the standard library type, requires no manual `free`/`malloc` equivalent at all — ownership and the `Drop` trait handle cleanup automatically.

3. Rebuild the same behavior with one changed constraint: Deliberately reproduce C track Lesson 4's use-after-free bug's *intent* in Rust (attempt to use a value after it's been moved elsewhere) and confirm the Rust compiler refuses to compile it, with an error message explicitly citing the move — read the actual compiler error text and confirm it correctly identifies both the move point and the invalid subsequent.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

Return a reference to a local value, then hold an immutable borrow while mutating its owner. Save each compiler diagnostic and fix the design with ownership or non-overlapping scopes.

Save the baseline, command, observed failure, and repaired result. Confirm the original success case still passes after the repair.

## Transfer

Change exactly one constraint in the main case for # 097 — Rust Ownership, Lifetimes, and Systems Thinking: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Which operations move a Vec?
2. When can a value implement Copy?
3. Why can shared references coexist?
4. What rule restricts mutable references?
5. Can a lifetime make a local value live longer?
6. When should an API return String instead of &str?

## Evidence to record

Create `lesson097/EVIDENCE.md` with exactly this structure:

```text
# Lesson 097 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 097 — Rust Ownership, Lifetimes, and Systems Thinking.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## Lab checklist

Use this checklist to decide whether the exercise is reproducible and reviewable:

1. List who owns each allocation in the original C version.
2. Identify every move in the Rust version.
3. Show the scope of each borrow.
4. Save the compiler output for both deliberate errors.
5. Explain the ownership or scope change that resolves each error.
6. Test empty and populated data structures.
7. State one semantic property the compiler cannot prove.

## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the lesson contract, rebuild its core case from a clean workspace, reproduce and repair its failure drill, solve the transfer constraint, and meet this evidence bar: You can explain value ownership and reference scopes, reproduce and interpret move and borrow errors, and choose a return type that fits the caller?s lifetime. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

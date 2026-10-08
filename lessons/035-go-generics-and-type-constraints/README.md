# Lesson 035 — Generics and Type Constraints

## Threads
**Primary:** Go/Algorithms

**Supporting:** type parameters; constraints; when abstraction helps

**Math companion:** logic and sets

**Time:** 3–4 hours. **Prerequisites:** Lesson 034.

**Time:** 3–4 hours. **Prerequisites:** Lesson 034.

## What you will be able to do

1. Explain how type parameters let one function operate over multiple concrete types.
2. Derive the minimum type constraint required by an algorithm's operations.
3. Generalize duplicated `contains` functions for strings and integers using `comparable`.
4. Explain when a generic abstraction improves reuse and when separate concrete code is clearer.

## Retrieve first (no notes)

1. How does a Go type satisfy an interface?
2. Which method set includes a pointer-receiver method?
3. Why can a direct method call compile while an interface assignment fails?
4. What should a caller's interface contain?

## 1. Problem and model

**What problem this solves.** Understand type parameters and constraints, and recognize when generics remove real duplication versus when they only make simple code harder to read.

**The contract.** First write concrete `containsInt` and `containsString` functions. They both scan a slice and compare each element with a target. Since equality is the shared operation, a generic version needs a type parameter whose constraint permits `==` and `!=`.

**Worked example.** Start with the smallest concrete case described by the source material: Implement `contains` for strings and integers first.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A Go type parameter is declared with a constraint. The predeclared `comparable` constraint permits equality comparisons for the admitted types, so `Contains[T comparable](items []T, target T) bool` can use `items[i] == target`.

**Second point.** Constraints should state only operations the implementation needs. If a function only stores, passes, or returns values, it may need no operation-specific constraint; if it compares with `<`, `comparable` is insufficient, and the constraint must provide an ordering operation or an appropriate ordered type set.

Generics do not automatically improve code. Prefer concrete functions when there is no meaningful duplicated algorithm, when behavior differs by type, or when a generic constraint obscures the contract.

**Worked example.** Compare the duplication.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: The reference implementation is in `code/reference/go-worked-examples/lesson-06-generics`.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Exercise the concrete and generic versions with present and absent values, empty slices, and duplicate values. Confirm they agree for both `int` and `string`. The generic version should have one source implementation; test behavior, not execution speed.

**Failure mode.** Replace the equality-based search with a less-than comparison while retaining `comparable`. Explain why the type constraint no longer permits the operation. Then contrast with a type set that admits ordered values and state the additional contract it imposes.

**Completion criterion.** You can move from concrete duplication to a generic function and justify every operation allowed by its constraint.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of `Contains([]int{2, 4}, 4)` and `Contains([]int{2, 4}, 3)`.
2. Predict the result for an empty slice.
3. Predict whether duplicate values change the truth result.
4. Predict which constraint permits `==` for the generic type parameter.
5. Predict whether `comparable` permits sorting with `<`.
6. Predict whether a slice itself is comparable, and distinguish it from its element type.
7. Predict what compile-time error appears if the generic body uses an operation not allowed by its constraint.
8. State one case where the generic version would add complexity without value.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement `contains` for strings and integers first.

2. Test empty, singleton, duplicate, found, and missing-value cases for both types.

3. Write `Contains[T comparable]` and explain why the constraint is necessary for equality.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Make the constraint broader than the algorithm requires.

2. Explain what extra operations you now appear to promise.

3. Change the algorithm to use an operation not permitted by `comparable` and diagnose the type-checking failure.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint: generalize a function from `int` and `string` to another comparable type, or change the operation from equality to ordering. State how the type parameter's admissible operations must change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is a type parameter?
2. What does a constraint specify?
3. Why does `Contains` need `comparable`?
4. Does `comparable` provide ordering with `<`?
5. What test cases distinguish the boundary behavior of `Contains`?
6. When is generic code less clear than concrete implementations?

## Evidence to record

Create `lesson035/EVIDENCE.md` with exactly this structure:

```text
# Lesson 035 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Adding methods or operations to a constraint that the function never uses.
- Assuming `comparable` permits ordering.
- Generalizing before confirming that the concrete algorithms are actually the same.
- Using a generic type parameter when behavior varies materially by type.
- Forgetting that type constraints are compile-time restrictions, not runtime checks.

## Precision checks

1. Show the duplicated concrete functions and identify the common algorithm.
2. State the exact operation that motivates the constraint.
3. Explain why `comparable` is sufficient for equality but not ordering.
4. State when you would keep the concrete versions instead.
5. Explain why the constraint is enforced at compile time.

Generalize only after identifying both the repeated algorithm and the exact operations its type parameter requires.
## University revision bridge

For the LIG track, connect type constraints to sets of admissible types and operations, then assess whether one generic definition accurately captures the shared algorithm.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can derive `Contains[T comparable]` from duplicated concrete algorithms, test its boundaries, explain the exact constraint, and justify when generic reuse is or is not worthwhile.

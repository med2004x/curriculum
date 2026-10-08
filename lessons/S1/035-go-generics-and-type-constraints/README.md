# Lesson 035 — Generics and Type Constraints

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand type parameters and constraints, and recognize when generics remove real duplication versus when they only make simple code harder to read.

**The contract.** First write concrete `containsInt` and `containsString` functions. They both scan a slice and compare each element with a target. Since equality is the shared operation, a generic version needs a type parameter whose constraint permits `==` and `!=`.

**Worked example.** Start with the smallest concrete case described by the source material: Implement `contains` for strings and integers first.

## 2. Mechanism and vocabulary

**Mechanism.** A Go type parameter is declared with a constraint. The predeclared `comparable` constraint permits equality comparisons for the admitted types, so `Contains[T comparable](items []T, target T) bool` can use `items[i] == target`.

**Second point.** Constraints should state only operations the implementation needs. If a function only stores, passes, or returns values, it may need no operation-specific constraint; if it compares with `<`, `comparable` is insufficient, and the constraint must provide an ordering operation or an appropriate ordered type set.

Generics do not automatically improve code. Prefer concrete functions when there is no meaningful duplicated algorithm, when behavior differs by type, or when a generic constraint obscures the contract.

**Worked example.** Compare the duplication.

## 3. Worked case from the reference material

Worked example: The reference implementation is in `code/reference/go-worked-examples/lesson-06-generics`.

## 4. Verification, failure, and evidence

**Verification.** Exercise the concrete and generic versions with present and absent values, empty slices, and duplicate values. Confirm they agree for both `int` and `string`. The generic version should have one source implementation; test behavior, not execution speed.

**Failure mode.** Replace the equality-based search with a less-than comparison while retaining `comparable`. Explain why the type constraint no longer permits the operation. Then contrast with a type set that admits ordered values and state the additional contract it imposes.

**Completion criterion.** You can move from concrete duplication to a generic function and justify every operation allowed by its constraint.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

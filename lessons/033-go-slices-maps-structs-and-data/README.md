# Lesson 033 — Go Slices, Maps, Structs, and Data Structures

## Threads
**Primary:** Go/Algorithms

**Supporting:** slice semantics; maps; structs; aggregates

**Math companion:** complexity

**Time:** 3–4 hours. **Prerequisites:** Lesson 032.

**Time:** 3–4 hours. **Prerequisites:** Lesson 032.

## What you will be able to do

1. Explain how a Go slice relates to an underlying array through length and capacity.
2. Predict how `append` and reslicing can affect aliases that share a backing array.
3. Aggregate records with a map while recognizing that map iteration order is unspecified.
4. Explain the copy semantics of a struct containing scalar fields versus slice or map fields.

## Retrieve first (no notes)

1. What is the difference between a package and a module?
2. Which identifiers are exported from a Go package?
3. What does `go.mod` record?
4. Why does an import cycle fail?

## 1. Problem and model

**What problem this solves.** Choose data representations that fit the operation: slices for ordered sequences, maps for key lookup and aggregation, and structs for records with named fields.

**The contract.** A slice value describes a view into an underlying array and has a length and capacity. Copying a slice copies its descriptor, not its elements; two slices can therefore share storage. A map stores key/value associations with unspecified iteration order. A struct groups fields; assigning it copies the field values, and reference-like fields such as slices still refer to shared underlying data.

**Worked example.** Define `type Account struct { Owner string; Balance int }`, keep an ordered slice of accounts, and build a `map[string]int` of balances by owner. Explain why the map's printed iteration order should not be part of a test expectation.

Self-check: identify which values are copied when an `Account` is passed by value and whether any field refers to shared storage.

## 2. Mechanism and vocabulary

**Mechanism.** `len(s)` is the number of accessible elements; `cap(s)` is the number of elements available from the slice start to the end of its backing array. Reslicing can expose additional elements up to capacity. `append` may reuse the backing array when capacity allows, or allocate a new one when it does not.

**Second point.** Maps are reference-like values: assigning one map variable to another does not duplicate its entries. A nil map can be read but cannot be assigned to; initialize it with `make` or a map literal before insertion.

The earlier `deposit(account Account, amount int) Account` example returns an updated struct because the parameter is a value copy. If the struct contains a slice or map, however, the copied field still refers to the shared backing data/map; “passed by value” does not mean deep copy.

**Worked example.** owner

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
func deposit(account Account, amount int) Account {
    if amount <= 0 {
        return account
    }
    account.Balance += amount
    return account
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Record slice `len` and `cap` before and after reslicing and appending. Create an alias and predict whether an in-capacity append is visible through it. Separately aggregate repeated keys and verify results without assuming map iteration order.

**Failure mode.** Pass an `Account` by value, change its scalar balance, and omit the returned value; then repeat with a slice field and mutate an element. Explain why the caller sees different effects. Also attempt insertion into a nil map and identify the runtime failure.

**Completion criterion.** You can explain slice aliasing and append allocation conditions, map initialization and iteration semantics, and the difference between copying a struct and deep-copying its contents.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict `len` and `cap` after slicing an array-backed slice.
2. Predict whether appending within capacity can affect another slice alias.
3. Predict what changes when append exceeds capacity.
4. Predict the effect of reading from a nil map and assigning to a nil map.
5. Predict whether two iterations over the same map must use the same order.
6. Predict which fields change when a struct value is copied and updated.
7. Predict whether mutating an element through a copied slice field can affect the original.
8. State which behavior is guaranteed and which is unspecified.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Define an `Account` struct and an update function that returns a modified value.

2. Store multiple accounts in a slice and aggregate balances in a map keyed by owner.

3. Construct two slices sharing a backing array; record the results of reslicing and appending within and beyond capacity.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Pass an `Account` to a function, change the local copy, and deliberately forget to return it.

2. Observe that the caller's value does not change.

3. Explain why.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 033 — Go Slices, Maps, Structs, and Data Structures: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What information does a slice value contain conceptually?
2. When can two slices share the same backing array?
3. What may `append` do when capacity is exhausted?
4. Why can a nil map be read but not assigned to?
5. Is Go map iteration order guaranteed?
6. Does passing a struct by value deep-copy slice fields?

## Evidence to record

Create `lesson033/EVIDENCE.md` with exactly this structure:

```text
# Lesson 033 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a slice as the backing array itself.
- Assuming append always allocates or never allocates.
- Mutating a nil map as if it had been initialized.
- Relying on map iteration order in tests or output contracts.
- Assuming a copied struct is a deep copy of referenced data.

## Precision checks

1. State the slice length/capacity relationship.
2. Explain the conditions under which append can reuse storage.
3. Explain map nil behavior and iteration ordering.
4. Distinguish copying a struct value from deep-copying its contents.
5. Choose a slice, map, or struct for a stated data operation and justify the choice.

Representation determines aliasing and update behavior; state which data is shared before reasoning about mutation.
## University revision bridge

For the LIG track, connect representation choices to array, associative-map, and record abstractions, then reason explicitly about Go's value and aliasing semantics.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can predict slice aliasing across append, aggregate records with a map without relying on iteration order, and explain exactly what is and is not copied when a struct containing slices or maps is passed by value.

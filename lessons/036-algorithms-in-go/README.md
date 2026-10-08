# Lesson 036 — Core Algorithms in Go

## Threads
**Primary:** Go/Algorithms

**Supporting:** search/sort; reusable algorithm patterns

**Math companion:** asymptotics

**Time:** 3–4 hours. **Prerequisites:** Lesson 035.

**Time:** 3–4 hours. **Prerequisites:** Lesson 035.

## What you will be able to do

1. Implement linear search and binary search over integer slices in Go.
2. State and preserve the sorted-input precondition for binary search.
3. Implement a sorting algorithm with an explicit permutation-and-order postcondition.
4. Compare operation counts and asymptotic costs against representative input sizes.

## Retrieve first (no notes)

1. What operation does the `comparable` constraint permit?
2. Why should a generic function's constraint admit only operations it uses?
3. When is a concrete implementation clearer than a generic one?
4. What is the difference between a type parameter and a runtime interface value?

## 1. Problem and model

**What problem this solves.** Translate algorithmic contracts into Go code and make preconditions, invariants, and complexity visible in the implementation.

**The contract.** Linear search examines candidates until a match or exhaustion, so it works on unsorted data and costs O(n). Binary search requires sorted data and repeatedly discards half of the remaining interval, giving O(log n) comparisons. A sort must return an ordered permutation: no elements are lost, duplicated, or changed.

**Worked example.** Search `[2, 4, 7, 9, 13]` for `9` with both algorithms. Linear search visits sequentially; binary search examines the middle and keeps the half that can still contain the target. Then repeat with a missing target and with an empty slice.

Self-check: test binary search on an unsorted slice and explain why a plausible answer does not restore its missing precondition.

## 2. Mechanism and vocabulary

**Mechanism.** Use half-open intervals `[low, high)` for binary search: initialize `low=0`, `high=len(items)`, inspect `mid=low+(high-low)/2`, and move a boundary so the interval strictly shrinks. Termination occurs when `low == high`.

**Second point.** For sorting, distinguish comparisons from data movement and extra memory. Insertion sort is simple and effective on small or nearly sorted inputs but has O(n²) worst-case comparisons. Merge sort has O(n log n) comparisons and typically uses O(n) auxiliary storage for slices.

Define boundary behavior explicitly: return both an index and a found flag if index zero must be distinguishable from absence; an empty slice has no valid element index.

**Worked example.** Add one boundary case to the # 036 — Core Algorithms in Go exercise.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
func BinarySearch(items []int, target int) (int, bool) {
    low, high := 0, len(items)
    for low < high {
        mid := low + (high-low)/2
        if items[mid] < target {
            low = mid + 1
        } else {
            high = mid
        }
    }
    if low < len(items) && items[low] == target {
        return low, true
    }
    return 0, false
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

This lower-bound search returns the first position whose value is at least `target`; the final equality check distinguishes a hit from an insertion position. Its precondition is ascending order.

Self-check: prove that each loop iteration strictly decreases `high-low`, including when `mid == low`.

## 4. Verification, failure, and evidence

**Verification.** Test empty, singleton, repeated-value, first-element, last-element, absent-below, absent-between, and absent-above cases. For sorting, verify adjacent order and compare element frequencies before/after to establish permutation preservation.

**Failure mode.** Change `high = mid` to `high = mid - 1` in the half-open lower-bound algorithm. Trace a one-element interval and show how the boundary convention is violated.

**Completion criterion.** You can state search preconditions and postconditions, trace a terminating interval invariant, and explain why one input case is not enough to verify a sort.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Trace binary search for target `7` in `[2, 4, 7, 9, 13]`.
2. Predict the result for a target smaller than every element.
3. Predict the result on an empty slice.
4. State why the interval is half-open and what `high` means.
5. Predict how many comparisons linear search needs for a missing value at size `n`.
6. State the ordering and permutation conditions for a correct sort.
7. Predict which sorting case is worst for insertion sort.
8. Identify the input assumption that makes binary search valid.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement linear search and lower-bound binary search, returning an explicit found result.

2. Test empty, singleton, duplicate, and absent-target cases.

3. Implement insertion sort or merge sort; verify sorted order and permutation preservation.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one precondition: allow duplicate values or search for the first occurrence rather than any occurrence. State how the postcondition and boundary tests change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is linear search's precondition and worst-case cost?
2. What is binary search's precondition and comparison cost?
3. What does the half-open interval `[low, high)` represent?
4. Why must the binary-search interval strictly shrink?
5. What two properties define a correct sorting result?
6. What are insertion sort's worst-case cost and useful input pattern?

## Evidence to record

Create `lesson036/EVIDENCE.md` with exactly this structure:

```text
# Lesson 036 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Applying binary search before sorting or otherwise establishing order.
- Mixing inclusive and half-open boundary conventions.
- Returning index zero to mean both “found at zero” and “not found.”
- Testing sorted order while overlooking lost or duplicated values.
- Claiming measured runtime proves an asymptotic bound.

## Precision checks

1. State the precondition and postcondition for each search.
2. Prove that the binary-search interval decreases.
3. Test boundary cases for empty and singleton inputs.
4. Verify both ordering and permutation after sorting.
5. Distinguish asymptotic analysis from observed timing.

An algorithm is a contract plus a termination argument; tests should target both the contract and its boundaries.
## University revision bridge

For the LIG track, connect loop invariants and asymptotic analysis to executable Go implementations with explicit slice boundary conventions.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 036 — Core Algorithms in Go without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can reproduce the main behavior and diagnose one failure from evidence. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

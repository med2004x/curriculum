# Lesson 036 — Core Algorithms in Go

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Translate algorithmic contracts into Go code and make preconditions, invariants, and complexity visible in the implementation.

**The contract.** Linear search examines candidates until a match or exhaustion, so it works on unsorted data and costs O(n). Binary search requires sorted data and repeatedly discards half of the remaining interval, giving O(log n) comparisons. A sort must return an ordered permutation: no elements are lost, duplicated, or changed.

**Worked example.** Search `[2, 4, 7, 9, 13]` for `9` with both algorithms. Linear search visits sequentially; binary search examines the middle and keeps the half that can still contain the target. Then repeat with a missing target and with an empty slice.

## 2. Mechanism and vocabulary

**Mechanism.** Use half-open intervals `[low, high)` for binary search: initialize `low=0`, `high=len(items)`, inspect `mid=low+(high-low)/2`, and move a boundary so the interval strictly shrinks. Termination occurs when `low == high`.

**Second point.** For sorting, distinguish comparisons from data movement and extra memory. Insertion sort is simple and effective on small or nearly sorted inputs but has O(n²) worst-case comparisons. Merge sort has O(n log n) comparisons and typically uses O(n) auxiliary storage for slices.

Define boundary behavior explicitly: return both an index and a found flag if index zero must be distinguishable from absence; an empty slice has no valid element index.

**Worked example.** Add one boundary case to the # 036 — Core Algorithms in Go exercise.

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

## 4. Verification, failure, and evidence

**Verification.** Test empty, singleton, repeated-value, first-element, last-element, absent-below, absent-between, and absent-above cases. For sorting, verify adjacent order and compare element frequencies before/after to establish permutation preservation.

**Failure mode.** Change `high = mid` to `high = mid - 1` in the half-open lower-bound algorithm. Trace a one-element interval and show how the boundary convention is violated.

**Completion criterion.** You can state search preconditions and postconditions, trace a terminating interval invariant, and explain why one input case is not enough to verify a sort.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

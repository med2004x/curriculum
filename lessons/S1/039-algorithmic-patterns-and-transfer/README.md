# Lesson 039 — Algorithmic Patterns and Transfer

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Array and interval problems often share a few state-maintenance patterns. Transfer depends on recognizing the pattern's preconditions rather than copying its code shape.

**The contract.** Two pointers often require sorted input or a monotone decision rule. A sliding window is valid when extending and shrinking changes the maintained condition predictably; for sum constraints, nonnegative values commonly provide that monotonicity. Interval merging requires defined endpoint semantics and sorting by start.

**Worked example.** In sorted `[1, 2, 4, 7, 11]`, find a pair summing to 9. Start at both ends: if the sum is too small, advance the left pointer; if too large, move the right pointer. Sorted order proves the discarded candidates cannot work.

## 2. Mechanism and vocabulary

**Mechanism.** For a fixed-width window of size `k`, add the entering element and subtract the leaving element, maintaining the aggregate in O(1) per step after initialization. For a variable window with nonnegative values, expand right and shrink left while the constraint is violated.

**Second point.** For closed intervals `[start,end]` sorted by start, merge when `next.start <= current.end`; otherwise emit the current interval. Half-open intervals can use a different touching-endpoint rule, so state the convention first.

Each pattern maintains a small state: pointer positions, window contents/aggregate, or the current merged interval. Its invariant is the reason a candidate can be discarded.

**Worked example.** Add one boundary case to the # 039 — Algorithmic Patterns and Transfer exercise.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
func PairWithSum(sorted []int, target int) (int, int, bool) {
    left, right := 0, len(sorted)-1
    for left < right {
        sum := sorted[left] + sorted[right]
        if sum == target {
            return left, right, true
        }
        if sum < target {
            left++
        } else {
            right--
        }
    }
    return 0, 0, false
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

The function requires ascending input and returns indices. If overflow is possible for the input domain, use a wider sum representation or state a bound.

## 4. Verification, failure, and evidence

**Verification.** Test empty/singleton inputs, duplicates, no solution, and multiple solutions under a documented return policy. For windows, test smallest sizes. For intervals, test nested, touching, disjoint, and unsorted inputs.

**Failure mode.** Run the pair algorithm on unsorted input and identify the invalid elimination step. For variable-sum windows, include a negative number and explain why nonnegative monotonicity no longer holds.

**Completion criterion.** You can state each pattern's preconditions, invariant, safe pointer movement, and complexity before applying it to a new problem.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

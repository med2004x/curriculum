# Lesson 039 — Algorithmic Patterns and Transfer

## Threads
**Primary:** Algorithms/Go

**Supporting:** two pointers; windows; intervals; decomposition

**Math companion:** combinatorics

**Time:** 3–4 hours. **Prerequisites:** Lesson 038.

**Time:** 3–4 hours. **Prerequisites:** Lesson 038.

## What you will be able to do

1. Recognize when sorted two-pointer reasoning can eliminate candidate pairs.
2. Maintain a sliding-window invariant for fixed or variable-width windows.
3. Merge intervals after sorting while stating the endpoint convention.
4. Select a pattern from its preconditions and justify each pointer movement.

## Retrieve first (no notes)

1. What is the difference between memoization and tabulation?
2. What recurrence models binary search?
3. What is a base case and decreasing measure in recursion?
4. Why must Master Theorem hypotheses be checked?

## 1. Problem and model

**What problem this solves.** Array and interval problems often share a few state-maintenance patterns. Transfer depends on recognizing the pattern's preconditions rather than copying its code shape.

**The contract.** Two pointers often require sorted input or a monotone decision rule. A sliding window is valid when extending and shrinking changes the maintained condition predictably; for sum constraints, nonnegative values commonly provide that monotonicity. Interval merging requires defined endpoint semantics and sorting by start.

**Worked example.** In sorted `[1, 2, 4, 7, 11]`, find a pair summing to 9. Start at both ends: if the sum is too small, advance the left pointer; if too large, move the right pointer. Sorted order proves the discarded candidates cannot work.

Self-check: would the same pointer movement be valid if the input were unsorted? State the precondition explicitly.

## 2. Mechanism and vocabulary

**Mechanism.** For a fixed-width window of size `k`, add the entering element and subtract the leaving element, maintaining the aggregate in O(1) per step after initialization. For a variable window with nonnegative values, expand right and shrink left while the constraint is violated.

**Second point.** For closed intervals `[start,end]` sorted by start, merge when `next.start <= current.end`; otherwise emit the current interval. Half-open intervals can use a different touching-endpoint rule, so state the convention first.

Each pattern maintains a small state: pointer positions, window contents/aggregate, or the current merged interval. Its invariant is the reason a candidate can be discarded.

**Worked example.** Add one boundary case to the # 039 — Algorithmic Patterns and Transfer exercise.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: trace why `sum < target` rules out every pair using the current left value and any value at or before `right`.

## 4. Verification, failure, and evidence

**Verification.** Test empty/singleton inputs, duplicates, no solution, and multiple solutions under a documented return policy. For windows, test smallest sizes. For intervals, test nested, touching, disjoint, and unsorted inputs.

**Failure mode.** Run the pair algorithm on unsorted input and identify the invalid elimination step. For variable-sum windows, include a negative number and explain why nonnegative monotonicity no longer holds.

**Completion criterion.** You can state each pattern's preconditions, invariant, safe pointer movement, and complexity before applying it to a new problem.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Trace the pointers for target 9 in `[1, 2, 4, 7, 11]`.
2. Predict what happens when the input is unsorted.
3. State the invariant for a fixed-width sum window.
4. Explain why negative values can invalidate a variable-sum window argument.
5. Predict whether closed intervals `[1,3]` and `[3,5]` merge under the stated rule.
6. Explain what sorting by start contributes to interval merging.
7. Identify why each pointer or boundary moves.
8. State the preconditions and cost of each selected pattern.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement sorted two-pointer pair-sum and state its input contract.

2. Implement a fixed-width sliding-window sum and explain its update invariant.

3. Sort and merge closed intervals, making the touching-endpoint policy explicit.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one precondition: make pair-sum input unsorted, include negative values in a variable-sum window, or switch intervals from closed to half-open. Explain which proof must change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Why does sorted order justify two-pointer elimination?
2. What invariant does a fixed-size sliding window maintain?
3. Why do nonnegative values matter for variable-sum windows?
4. What endpoint rule merges touching closed intervals?
5. Why sort intervals before merging?
6. What proof justifies each pointer update?

## Evidence to record

Create `lesson039/EVIDENCE.md` with exactly this structure:

```text
# Lesson 039 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Applying two-pointer elimination to unsorted input.
- Using a variable-sum window with negative values without proving monotonicity.
- Forgetting to subtract the element leaving a fixed-size window.
- Leaving interval endpoint semantics implicit.
- Moving a pointer without stating which candidates become impossible.

## Precision checks

1. State each pattern's precondition and maintained invariant.
2. Explain the safe pointer movement.
3. Test boundary and return-policy cases.
4. State whether interval endpoints are closed or half-open.
5. Explain whether the original proof survives a changed assumption.

Pattern transfer is proof transfer: carry the invariant and preconditions to the new problem, not just the pointer-shaped code.
## University revision bridge

For the LIG track, connect pointer patterns to monotonicity and interval reasoning, then check how changed assumptions affect the proof.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 039 — Algorithmic Patterns and Transfer without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can reproduce the main behavior and diagnose one failure from evidence. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

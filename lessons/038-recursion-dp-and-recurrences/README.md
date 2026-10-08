# Lesson 038 — Recursion, Dynamic Programming, and Recurrences

## Threads
**Primary:** Algorithms/Math

**Supporting:** overlapping subproblems; memoization; recurrences

**Math companion:** recurrence relations

**Time:** 3–4 hours. **Prerequisites:** Lesson 037.

**Time:** 3–4 hours. **Prerequisites:** Lesson 037.

## What you will be able to do

1. Write recursive solutions with an explicit base case and a decreasing measure.
2. Identify overlapping subproblems and compare naive recursion with memoization and tabulation.
3. Translate a recursive algorithm into a recurrence for its running time.
4. Apply an appropriate Master Theorem case and recognize when the basic form does not apply directly.

## Retrieve first (no notes)

1. What is the BST invariant?
2. What data structure does BFS use?
3. Why does graph traversal need a visited set?
4. What makes a directed graph acyclic?

## 1. Problem and model

**What problem this solves.** Recursion expresses a problem using smaller instances; dynamic programming avoids recomputing overlapping subproblems. Recurrence relations model recursive running time and help compare the cost of algorithm designs.

**The contract.** **Recurrence relations.** A recurrence defines a sequence (or a function's running time) in terms of smaller instances of itself.

**Worked example.** Fibonacci recurrence: `F(0)=0`, `F(1)=1`, `F(n)=F(n-1)+F(n-2)`. The naive recursive version recomputes values such as `F(n-2)` many times. Memoization stores results by `n`; bottom-up tabulation computes `F(0)` through `F(n)` once each.

Self-check: what quantity strictly decreases in every recursive call, and what base case stops the recursion?

## 2. Mechanism and vocabulary

**Mechanism.** A recurrence for merge sort is `T(n)=2T(n/2)+Θ(n)`: two recursive subproblems plus linear merging. For binary search, `T(n)=T(n/2)+Θ(1)`, yielding Θ(log n).

**Second point.** For `T(n)=aT(n/b)+f(n)`, compare `f(n)` with `n^(log_b a)` and check the exact hypotheses of the theorem version being used. In the basic three-case form, `2T(n/2)+n log n` lies between the stated equal-order and polynomially-larger cases; an extended theorem handles logarithmic factors and gives Θ(n log² n).

Memoization is top-down caching of function results; tabulation is bottom-up evaluation in dependency order. Both can reduce the Fibonacci example to O(n) time and O(n) storage; a two-variable iterative version uses O(1) extra storage.

**Worked example.** Identify `a, b, f(n)`, determine which case applies, and state the resulting complexity class.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
func Fibonacci(n int) int {
    if n < 2 {
        return n
    }
    previous, current := 0, 1
    for index := 2; index <= n; index++ {
        previous, current = current, previous+current
    }
    return current
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

The loop preserves the invariant that `previous == F(index-2)` and `current == F(index-1)` before each update. Note that this `int` implementation can overflow for sufficiently large `n`; state an input bound or use a suitable big-integer type if exact larger values are required.

Self-check: trace the invariant for `n=0`, `n=1`, and `n=2`.

## 4. Verification, failure, and evidence

**Verification.** Compare naive recursion, memoization, and tabulation on small `n` for equal answers. For a formal bound, choose constants `c` and `n₀`, then prove the inequality for all `n ≥ n₀`; checking two values is a sanity check, not a proof.

**Failure mode.** Try to force `T(n)=2T(n/2)+n log n` into the basic three-case Master Theorem. Identify exactly which case hypotheses fail; then use a recursion-tree argument or an extended Master Theorem instead of claiming the recurrence has no solution.

**Completion criterion.** You can identify overlapping subproblems, explain the time/space trade-off of dynamic programming, derive a recurrence, and check the hypotheses before applying a theorem.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Trace the recursive calls for `F(4)` and identify repeated subproblems.
2. Predict the number of stored results needed by top-down memoization through `F(n)`.
3. Predict the base cases and decreasing measure for a recursive binary search.
4. Identify `a`, `b`, and `f(n)` for merge sort.
5. Solve `T(n)=T(n/2)+1` asymptotically.
6. State why the basic three-case theorem does not directly cover `2T(n/2)+n log n`.
7. Predict the time and extra space for memoized Fibonacci.
8. State the input bound needed to avoid integer overflow in a fixed-width implementation.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement naive recursive Fibonacci and draw its call tree for a small argument.

2. Add memoization and bottom-up tabulation; explain which repeated computation is removed.

3. Derive and solve the recurrences for binary search and merge sort, stating theorem assumptions.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Remove the memo table or fail to advance the iterative state; identify repeated work or nontermination.

2. Analyze `2T(n/2)+n log n` with a recursion tree or an extended Master Theorem and state why that tool is needed.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint: compute Fibonacci modulo a fixed positive integer, or change the recurrence to a different pair of subproblems. Explain which state and recurrence terms must change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What defines a valid recursive base case?
2. What property identifies overlapping subproblems?
3. How do memoization and tabulation differ?
4. What recurrence models binary search?
5. What recurrence models merge sort?
6. Why must theorem hypotheses be checked before applying a Master Theorem case?

## Evidence to record

Create `lesson038/EVIDENCE.md` with exactly this structure:

```text
# Lesson 038 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Writing recursion without a base case or decreasing measure.
- Assuming dynamic programming always uses less total memory.
- Applying a theorem outside its hypotheses.
- Treating sample inequality checks as a proof for all sufficiently large `n`.
- Ignoring fixed-width integer overflow in an implementation.

## Precision checks

1. State the recursive base case and decreasing measure.
2. Identify overlapping states and the memoization key.
3. Compare time and storage for naive, memoized, and iterative Fibonacci.
4. Derive a recurrence from the algorithm structure.
5. Verify that a theorem's hypotheses match the recurrence.

Dynamic programming changes the computation graph by reusing solved states; recurrence analysis describes the work that remains.
## University revision bridge

For the LIG track, connect recurrence relations and asymptotic proofs to recursive program structure, while distinguishing a proof from numerical spot checks.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 038 — Recursion, Dynamic Programming, and Recurrences without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can identify `a`, `b`, and `f(n)` from a recurrence and correctly apply the Master Theorem's three cases, and you can produce a formal Big-O proof with explicit constants for a new polynomial, rather than asserting the bound by inspection alone. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

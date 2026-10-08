# Lesson 038 — Recursion, Dynamic Programming, and Recurrences

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Recursion expresses a problem using smaller instances; dynamic programming avoids recomputing overlapping subproblems. Recurrence relations model recursive running time and help compare the cost of algorithm designs.

**The contract.** **Recurrence relations.** A recurrence defines a sequence (or a function's running time) in terms of smaller instances of itself.

**Worked example.** Fibonacci recurrence: `F(0)=0`, `F(1)=1`, `F(n)=F(n-1)+F(n-2)`. The naive recursive version recomputes values such as `F(n-2)` many times. Memoization stores results by `n`; bottom-up tabulation computes `F(0)` through `F(n)` once each.

## 2. Mechanism and vocabulary

**Mechanism.** A recurrence for merge sort is `T(n)=2T(n/2)+Θ(n)`: two recursive subproblems plus linear merging. For binary search, `T(n)=T(n/2)+Θ(1)`, yielding Θ(log n).

**Second point.** For `T(n)=aT(n/b)+f(n)`, compare `f(n)` with `n^(log_b a)` and check the exact hypotheses of the theorem version being used. In the basic three-case form, `2T(n/2)+n log n` lies between the stated equal-order and polynomially-larger cases; an extended theorem handles logarithmic factors and gives Θ(n log² n).

Memoization is top-down caching of function results; tabulation is bottom-up evaluation in dependency order. Both can reduce the Fibonacci example to O(n) time and O(n) storage; a two-variable iterative version uses O(1) extra storage.

**Worked example.** Identify `a, b, f(n)`, determine which case applies, and state the resulting complexity class.

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

## 4. Verification, failure, and evidence

**Verification.** Compare naive recursion, memoization, and tabulation on small `n` for equal answers. For a formal bound, choose constants `c` and `n₀`, then prove the inequality for all `n ≥ n₀`; checking two values is a sanity check, not a proof.

**Failure mode.** Try to force `T(n)=2T(n/2)+n log n` into the basic three-case Master Theorem. Identify exactly which case hypotheses fail; then use a recursion-tree argument or an extended Master Theorem instead of claiming the recurrence has no solution.

**Completion criterion.** You can identify overlapping subproblems, explain the time/space trade-off of dynamic programming, derive a recurrence, and check the hypotheses before applying a theorem.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

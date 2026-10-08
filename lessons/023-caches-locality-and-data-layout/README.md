# Lesson 023 — Caches, Locality, and Data Layout

## Threads
**Primary:** Architecture/C

**Supporting:** cache lines; locality; layout consequences

**Math companion:** sequences and matrices intuition

**Time:** 3–4 hours. **Prerequisites:** Lesson 022.

**Time:** 3–4 hours. **Prerequisites:** Lesson 022.

## What you will be able to do

1. Explain temporal locality, spatial locality, cache lines, and working sets.
2. Predict why row-major traversal of a row-major array usually uses cache lines more effectively than column-major traversal.
3. Explain how cache capacity and data layout affect observed performance.
4. Design a fair comparison and state what its measurements can support.

## Retrieve first (no notes)

1. What is the difference between combinational and sequential circuits?
2. What does a full-adder carry-in represent?
3. Why is final carry-out not always signed overflow?
4. What does a truth table specify, and what does it not say about timing?

## 1. Problem and model

**What problem this solves.** Explain the memory hierarchy, why caches exist, and predict when code will be cache-friendly or cache-hostile.

**The contract.** CPUs use storage levels with different latency, capacity, and proximity to the core. Caches transfer data in lines, so nearby bytes may be brought in together. Exact latency and line size depend on the processor and memory system; numeric claims require a named target or a measurement.

**Worked example.** C stores a declared multidimensional array in row-major order: the last subscript varies over contiguous elements. In `a[i][j]`, varying `j` in the inner loop walks adjacent elements; varying `i` jumps between rows by approximately one row's size.

Self-check: hold dimensions, element type, total work, initialization, and compiler options constant when comparing loop orders.

## 2. Mechanism and vocabulary

**Mechanism.** Spatial locality is access to nearby addresses; temporal locality is reuse of recently accessed data. A cache line is the unit transferred between adjacent cache levels, so sequential access can use data fetched alongside the requested element.

**Second point.** If a working set fits in a cache level, traversal order may have little effect. If it exceeds cache capacity, repeated misses or long strides may matter. Prefetching, compiler transformations, cache state, and other machine activity can all affect measured results.

For a row-major array with `C` columns, element `(i,j)` is at `base + (i*C + j) * sizeof(element)`. This address calculation makes the contiguous dimension explicit.

**Worked example.** Sum the same `rows × columns` array with two loops that differ only in nesting order. Make the sum observable so the compiler cannot discard the work. Repeat each measurement, alternate variant order, and report several samples or a median rather than a selected best run.

Self-check: both loops must compute the same sum before their timings are compared.

## 3. Worked case from the reference material

Record dimensions, element size, byte footprint, loop order, compiler options, and repeated elapsed-time samples. Approximate the array's footprint as `rows × columns × sizeof(element)`.

Row-major traversal is expected to benefit from spatial locality in the usual C layout. Do not predict a fixed speedup: machine, compiler, dimensions, and cache state affect the result.

Self-check: elapsed times are observations; an explanation in terms of cache misses is an inference unless supported by counters or other evidence.

## 4. Verification, failure, and evidence

**Verification.** Produce a table with variant, dimensions, compiler settings, repeated timing samples, and cache-miss counts if available. Explain whether the evidence supports the locality hypothesis and note measurement limits.

**Failure mode.** Shrink the array size until row-major and column-major run at the same speed.

**Completion criterion.** You can explain why access order matters for a row-major array, identify the pattern in a loop, and interpret a controlled measurement without treating one machine's timing as universal.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Which subscript varies contiguously in C's `a[i][j]` layout?
2. Estimate the bytes used by your chosen dimensions and element type.
3. Which loop order should usually have stronger spatial locality?
4. What may happen when the working set fits in cache?
5. Why can a large row stride reduce useful line reuse?
6. Why are repeated timing samples better than one result?
7. Does `volatile` make unsynchronized cross-thread access safe? Explain.
8. Mark which parts of your explanation are measured and which are inferred.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Write two correct summation loops over one initialized two-dimensional array, changing only loop order.

2. Compare multiple sizes, including a small working set and one substantially larger working set.

3. Record repeated timings and, when available, cache-miss counters; confirm both variants produce the same sum.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Shrink the array size until row-major and column-major run at the same speed.

2. Find (approximately, by binary search on size) the array size where the effect disappears, and explain why in terms of the array now fitting inside L2 or L3 cache regardless of access order — locality only matters when the working set exceeds cache capacity.

3. Create a misleading use of `volatile` and explain why it does not establish a safe communication protocol between threads.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint: use a one-dimensional array with a configurable stride, or change element size while keeping total bytes approximately constant. Predict the locality effect before measuring it.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Define spatial locality and temporal locality.
2. What is a cache line, and how can it help sequential traversal?
3. Which index varies over contiguous storage in a row-major C array?
4. How do you estimate a two-dimensional array's byte footprint?
5. Why may loop order matter less when the working set fits in cache?
6. What evidence supports a claim that one variant had more cache misses?

## Evidence to record

Create `lesson023/EVIDENCE.md` with exactly this structure:

```text
# Lesson 023 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Attributing every timing difference to cache misses without supporting evidence.
- Comparing different amounts of work or different initialization paths.
- Reporting one timing as a stable result.
- Forgetting that C multidimensional arrays are row-major.
- Using `volatile` as a substitute for atomics or synchronization between threads.

## Precision checks

1. State the array layout and address calculation used.
2. Show that both variants compute the same result.
3. Report dimensions, element size, and compiler settings with timing samples.
4. Separate observed timing from cache-based interpretation.
5. Explain why `volatile` does not establish inter-thread synchronization.

Performance measurements need a reproducible workload and a restrained conclusion: locality predicts a tendency, not a fixed speedup.
## University revision bridge

For the LIG track, connect cache hierarchy and data layout to architecture. Treat timing as an experiment and state the platform assumptions behind numerical claims.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain both forms of locality, calculate approximate array footprint, identify contiguous versus strided access, and describe what a controlled timing or cache-counter comparison can support.

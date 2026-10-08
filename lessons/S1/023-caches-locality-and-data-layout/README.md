# Lesson 023 — Caches, Locality, and Data Layout

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Explain the memory hierarchy, why caches exist, and predict when code will be cache-friendly or cache-hostile.

**The contract.** CPUs use storage levels with different latency, capacity, and proximity to the core. Caches transfer data in lines, so nearby bytes may be brought in together. Exact latency and line size depend on the processor and memory system; numeric claims require a named target or a measurement.

**Worked example.** C stores a declared multidimensional array in row-major order: the last subscript varies over contiguous elements. In `a[i][j]`, varying `j` in the inner loop walks adjacent elements; varying `i` jumps between rows by approximately one row's size.

## 2. Mechanism and vocabulary

**Mechanism.** Spatial locality is access to nearby addresses; temporal locality is reuse of recently accessed data. A cache line is the unit transferred between adjacent cache levels, so sequential access can use data fetched alongside the requested element.

**Second point.** If a working set fits in a cache level, traversal order may have little effect. If it exceeds cache capacity, repeated misses or long strides may matter. Prefetching, compiler transformations, cache state, and other machine activity can all affect measured results.

For a row-major array with `C` columns, element `(i,j)` is at `base + (i*C + j) * sizeof(element)`. This address calculation makes the contiguous dimension explicit.

**Worked example.** Sum the same `rows × columns` array with two loops that differ only in nesting order. Make the sum observable so the compiler cannot discard the work. Repeat each measurement, alternate variant order, and report several samples or a median rather than a selected best run.

## 3. Worked case from the reference material

Record dimensions, element size, byte footprint, loop order, compiler options, and repeated elapsed-time samples. Approximate the array's footprint as `rows × columns × sizeof(element)`.

Row-major traversal is expected to benefit from spatial locality in the usual C layout. Do not predict a fixed speedup: machine, compiler, dimensions, and cache state affect the result.

## 4. Verification, failure, and evidence

**Verification.** Produce a table with variant, dimensions, compiler settings, repeated timing samples, and cache-miss counts if available. Explain whether the evidence supports the locality hypothesis and note measurement limits.

**Failure mode.** Shrink the array size until row-major and column-major run at the same speed.

**Completion criterion.** You can explain why access order matters for a row-major array, identify the pattern in a loop, and interpret a controlled measurement without treating one machine's timing as universal.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

# Lesson 096 — Benchmarking, Profiling, and Locality

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Produce measurements that can distinguish a real performance change from noise.

**The contract.** Keep workload, build, and environment explicit. Warm-up, repeated runs, variance, throughput, latency, dead-code elimination, and workload drift all affect what a benchmark can establish.

**Worked example.** Start with the smallest concrete case described by the source material: Benchmark two implementations of the same operation.

## 2. Mechanism and vocabulary

**Mechanism.** **Why profiling, not intuition, should drive optimization effort.** Linux tools Lesson 5 established the core principle: intuition about where time is spent is often wrong.

**Second point.** Go `pprof` can sample CPU call paths and attribute allocations to sites; Linux profiling tools can also expose hardware behavior such as cache misses. Use profiles to select a candidate cost, then benchmark the same correct workload before and after one change.

**Worked example.** Repeat enough times to report a distribution, not one favorite number.

## 3. Worked case from the reference material

Worked example: Benchmark two implementations of the same operation.

## 4. Verification, failure, and evidence

**Verification.** Save repeated baseline and candidate measurements with units, input size, compiler flags, and environment. Include a CPU or allocation profile that identifies a candidate hotspot, then show correctness checks and whether the measured change exceeds run-to-run variation.

**Failure mode.** Make a *second*, unrelated change to your function at the same time as your intended fix (e.g. also reorder some unrelated logic, or add an unrelated micro-optimization you assumed was harmless) and re-benchmark.

**Completion criterion.** You can reproduce a stated baseline, profile a representative workload, make one change, and report whether the repeated results support a bounded performance claim.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

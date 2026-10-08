# Lesson 019 — Searching, Sorting, and Measuring Work

## Threads
**Primary:** C/Algorithms

**Supporting:** search; sorting; measurement and benchmarking

**Math companion:** asymptotic cost; logarithms; stable ordering

**Time:** 3–4 hours. **Prerequisites:** Lesson 018.

## What you will be able to do

1. Explain the difference between linear search and binary search.
2. Explain the difference between easy-to-implement sorts and efficient sorts.
3. Measure a task with more than one data point and explain what a benchmark does and does not show.
4. Interpret the relationship between algorithmic cost and empirical runtime.

## Retrieve first (no notes)

1. What is a hash table?
2. What is a collision?
3. What is the load factor?
4. Why do hash tables resize?
5. What is a bucket or slot?
6. Why is expected lookup near constant in a well-distributed table?

## 1. Problem and model

The goal of searching is to find whether an item exists and, if desired, where it is stored. The basic algorithms are:

- linear search: scan from start to end, O(n)
- binary search: requires sorted input, repeatedly halves the search space, O(log n)

The goal of sorting is to arrange elements in order. Common families include:

- insertion sort: simple and adaptive, O(n^2) worst case
- merge sort: divide and conquer, O(n log n)
- quicksort: fast in practice but can hit poor partitions, O(n^2) worst case
- heapsort: O(n log n) and in-place

Benchmarking is not just “time the program once.” A real measurement must isolate the workload and report enough samples to distinguish noise from systematic differences.

## 2. Mechanism and vocabulary

The core vocabulary is:

- sorted order: the required order in which elements are arranged
- comparison sort: a sort that decides order using comparisons between elements
- stable sort: if two records have the same key, their relative order is preserved
- benchmark: a reproducible workload used to measure runtime
- warm-up: running the code before measuring to let caches, JITs, or other state settle
- distribution: the spread of several measurements

A benchmark is not credible if it runs one tiny case and claims a general result. A good benchmark should control input size, data pattern, and measurement method.

## 3. Worked case: search and sorting on a list

Example list:

```text
[3, 7, 1, 9, 2]
```

Linear search for value `9` checks each element in sequence until it finds it, which takes 4 comparisons in this case, and 5 in the worst case for 5 elements.

Binary search requires sorted input:

```text
[1, 2, 3, 7, 9]
```

The search compares the target to the middle value, then halves the remaining range. That reduces the candidate set logarithmically.

Sorting the same list with insertion sort repeatedly inserts each new value into the correct place. Merge sort instead divides the list into halves, sorts each half, and merges the results.

## 4. Verification, failure, and evidence

Evidence for an algorithmic claim should include:

- the input size
- the algorithm used
- the observed number of comparisons or operations
- the measured runtime or timing sample
- the conclusion supported by the data

Common benchmarking mistakes:

- measuring one run when a single run is noisy
- changing the input while comparing algorithms
- including startup costs that are unrelated to the algorithm itself
- comparing only the best result instead of a distribution of results

A valid conclusion is narrow: “for this workload and this machine, algorithm A was faster than algorithm B.” It is not a universal statement unless the inputs and setup are controlled.

## Predict, then observe

1. Predict whether binary search can work on unsorted input.
2. Predict the worst-case cost of insertion sort on `n` items.
3. Predict the worst-case cost of merge sort on `n` items.
4. Predict whether a benchmark of one timed run is enough to conclude a performance difference.
5. Predict what changes when input size doubles in a linear-time algorithm.
6. Predict what changes when input size doubles in a logarithmic-time algorithm.

## Build from blank

1. Implement a linear search and a binary search on an integer array.
2. Implement insertion sort and merge sort on a small integer array.
3. Measure the runtime of each algorithm over several input sizes.
4. Record the input sizes and observed timing pattern.

## Failure drill

1. Sort an array with duplicate keys and explain whether the sort is stable.
2. Time a function after a warm-up and compare it to the first run.
3. Compare an already-sorted array to a reverse-sorted array with insertion sort.
4. Explain what each benchmark result does and does not prove.

## Transfer

Change exactly one condition: sort a list of strings instead of integers, or search for a missing value as well as a present value. Explain whether the algorithm remains correct, and whether the cost model changes.

## Retrieval questions

1. What is a comparison sort?
2. What is a stable sort?
3. What is the cost of linear search?
4. What is the cost of binary search on sorted data?
5. What is a benchmark?
6. Why is repeated measurement better than a single timing?

## Evidence to record

Create `lesson019/EVIDENCE.md` with exactly this structure:

```text
# Lesson 019 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Using binary search on unsorted data.
- Treating a single timing as a universal truth.
- Confusing algorithmic complexity with the runtime of one machine.
- Forgetting that sort stability matters for multi-key data.

## Precision checks

1. State the cost of linear search in one sentence.
2. State the cost of binary search in one sentence.
3. Explain why insertion sort is often slower than merge sort on large inputs.
4. Explain one reason a benchmark may be misleading.
5. State the difference between a stable and unstable sort.

## University revision bridge

The academic model is asymptotic cost: we reason about how the algorithm scales with input size. The engineering model is benchmarking: we measure one machine under one workload and then connect that evidence to the theoretical model.

## Assigned material

- `07-data-structures-and-algorithms/lesson-10-linear-search/README.md`
- `07-data-structures-and-algorithms/lesson-12-sorting/README.md`
- `07-data-structures-and-algorithms/lesson-17-binary-search/README.md`
- `18-performance-engineering/lesson-01-benchmarking/README.md`

## Mastery gate

You pass when you can explain the basic search and sort families, compare their asymptotic cost, and describe how a benchmark should be designed so that the measured difference corresponds to the actual algorithmic difference rather than just machine noise.

# Lesson 019 — Searching, Sorting, and Measuring Work

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

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

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

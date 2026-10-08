# S1 COURSE EXERCISE - ADS1
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Algorithmique et structures de données 1
#
# **Semester:** S1
# **Attached engineering lesson:** 004 — Conditions, Logic, and Proof Thinking
#
# ## What you must be able to do
#
# Write an algorithm in precise pseudocode, trace it by hand, justify its correctness on ordinary and boundary inputs, then implement it in C. You must be comfortable with variables, types, expressions, input/output, sequences, conditions, loops, counters, accumulators, arrays, matrices, strings, records/structures, procedures/functions, parameter passing, linear search, basic sorting, and elementary complexity reasoning.
#
# ## Course sequence
#
# 1. Algorithm representation: problem statement, input, output, constraints, invariants.
# 2. Basic control: assignment, sequence, condition, nested condition, loop.
# 3. Systematic patterns: counter, accumulator, min/max, flag, sentinel, frequency.
# 4. Arrays and matrices: traversal, filtering, aggregation, insertion/removal.
# 5. Strings and records: character traversal, structured records, arrays of records.
# 6. Functions/procedures: parameters, return values, local state, decomposition.
# 7. Search and sort: linear search, insertion sort, selection sort, stability and cost.
# 8. Correctness and complexity: invariants, termination, best/worst cases, O(1), O(n), O(n²) intuition.
#
# ## Teaching notes
#
# An algorithm is a finite, unambiguous procedure transforming an input satisfying stated constraints into the required output. A good answer starts with the specification, not with C syntax.
#
# For a loop, identify the loop variable, the progress measure, the invariant, and the termination condition. A counter counts events. An accumulator combines values. A flag records whether a predicate has been seen. A sentinel terminates input when the problem defines a sentinel.
#
# For an array `a[0..n-1]`, valid indices are `0` through `n-1`. Most first-year algorithm mistakes are boundary mistakes disguised as arithmetic.
#
# Linear search examines elements until the key is found or the sequence is exhausted. Its worst-case cost is proportional to `n`. Insertion sort grows a sorted prefix one item at a time. Its worst-case number of shifts/comparisons is quadratic, while an already sorted input is close to linear.
#
# Correctness is not the same thing as passing a few examples. You should be able to state why the algorithm preserves the relevant invariant and why the loop terminates.
#
# ## Worked examples to reproduce by hand
#
# ### Example A — Minimum and position
#
# Given `12, 4, 9, 4, 15`, trace an algorithm that keeps the smallest value and its first position. Write the state after each element before coding it.
#
# ### Example B — Frequency
#
# Given digits `1, 2, 1, 3, 2, 1`, build a frequency table. State the invariant after processing the first `k` values.
#
# ### Example C — Insertion sort
#
# Sort `[5, 2, 4, 2]` using insertion sort. Show the array after every insertion. Then count comparisons and shifts.
#
# ## Exercises
#
# 1. Write pseudocode to read `n` integers and output minimum, maximum, sum, and average.
# 2. Count how many values are even, odd, negative, and zero in one pass.
# 3. Given an array, return the first position of a key or `-1` if absent.
# 4. Reverse an array in place and state the invariant of the swap loop.
# 5. Compute every row sum and column sum of a matrix.
# 6. Build an array of records for 10 students and return the student with the highest average.
# 7. Insert one value into an already sorted array while preserving order.
# 8. Trace insertion sort on an input with duplicates and state whether equal items keep their relative order.
# 9. Rewrite one solution as a function and identify which values must be parameters versus local state.
# 10. Give the worst-case operation count class for each previous solution.
#
# ## Problem-solving lab
#
# ### Lab A — Smallest missing positive-ish value
#
# Given an array of integers, find the smallest non-negative value that does not occur. First solve by scanning with a frequency structure. Then design a solution using only the representations allowed by the problem statement. Compare time and memory.
#
# ### Lab B — Exam-style trace
#
# An algorithm scans `a[0..n-1]`, updates `m`, `p`, and `c`, and contains nested conditions. Create your own six-element trace table, execute every branch, and explain the final values. Do not rely on running code.
#
# ## Coding workspace
#
# Use the C starters in the linked engineering lessons when a task calls for implementation. The mathematics and written traces remain required evidence. No general-purpose input/output checker is supplied for arbitrary algorithm statements: define a function contract and write ordinary, boundary, and adversarial tests for each implementation.
#
# ## Exam preparation
#
# You should recognize these question forms immediately: write pseudocode from a statement; complete a partially given algorithm; trace variable state; find and repair an off-by-one error; calculate a loop count; choose an appropriate data representation; implement a basic search/sort; justify a loop invariant; classify complexity.
#
# ## Mastery gate
#
# From a blank page, solve one array problem, one matrix problem, one record problem, one search/sort problem, and one trace problem. For each, provide pseudocode, boundary cases, and a brief correctness argument. Then implement one without copying from a reference.

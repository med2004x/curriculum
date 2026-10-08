/*
 * STUDENT PROBLEM SET — Lesson 019.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Implement a linear search and a binary search on an integer array.
 * 2. Implement insertion sort and merge sort on a small integer array.
 * 3. Measure the runtime of each algorithm over several input sizes.
 * 4. Record the input sizes and observed timing pattern.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 019: Searching, Sorting, and Measuring Work
 * Your task: sort in place and state the loop invariant.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

static void insertion_sort(int *a, size_t n) {
    /* TODO: implement the assigned behavior. */
}

int main(void) {
    int a[] = {5, 2, 4, 1, 3};
    insertion_sort(a, 5);
    for (size_t i = 0; i < 5; i++) printf("%d%c", a[i], i == 4 ? '\n' : ' ');
    return 0;
}

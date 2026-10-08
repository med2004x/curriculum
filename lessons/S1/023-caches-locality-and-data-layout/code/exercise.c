/*
 * STUDENT PROBLEM SET — Lesson 023.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Write two correct summation loops over one initialized two-dimensional array, changing only loop order.
 * 2. Compare multiple sizes, including a small working set and one substantially larger working set.
 * 3. Record repeated timings and, when available, cache-miss counters; confirm both variants produce the same sum. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 023: Caches, Locality, and Data Layout
 * Your task: implement contiguous row-major traversal.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

static long row_major_sum(const int *a, size_t rows, size_t cols) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    int a[2][3] = {{1,2,3},{4,5,6}};
    printf("%ld\n", row_major_sum(&a[0][0], 2, 3));
    return 0;
}

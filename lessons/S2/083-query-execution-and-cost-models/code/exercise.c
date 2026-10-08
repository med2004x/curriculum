/*
 * STUDENT PROBLEM SET — Lesson 083.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Build the smallest useful version: Implement the iterator interface (`Open()`, `Next() (row, hasMore)`, `Close()`, or the Go-idiomatic equivalent) and a `SeqScan` operator that iterates over an in-memory slice of rows (standing in for reading table pages, per Lesson 1) one at a time.
 * 2. Extend it with one explicit boundary case: Implement a `Filter` operator that wraps any child iterator and a predicate function, calling the child's `Next()` and only returning rows where the predicate is true — confirm it correctly stops (returns `hasMore = false`) once the child is exhausted, not just once it happens to find no more matching rows in a finite lookahead.
 * 3. Rebuild the same behavior with one changed constraint: Implement a `Project` operator that wraps a child iterator and a list of column indices/names to keep, transforming each row before passing it up. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 083: Query Execution and Cost Models
 * Your task: encode the cost assumptions from the lecture, then compare two plans.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>

typedef struct {
    double rows;
    double cost;
} Plan;

static Plan estimate(double rows, double selectivity, double page_cost) {
    /* TODO: implement the assigned behavior. */
    return (Plan){0};
}

int main(void) {
    Plan p = estimate(100000, 0.01, 1.0);
    printf("rows=%.0f cost=%.2f\n", p.rows, p.cost);
    return 0;
}

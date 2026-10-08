/*
 * STUDENT PROBLEM SET — Lesson 020.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Define `LogRecord`, the line grammar, field bounds, and malformed-line policy in a short design sketch.
 * 2. Implement a parser that accepts one valid record and rejects one malformed record without modifying the output record on failure.
 * 3. Add file loading, an in-memory array with an explicit capacity policy, a search command, and clear stdout/stderr behavior.
 * 4. Split the program into at least two `.c` files and a header; add automated tests for parser and search behavior. Save the design sketch before coding. Keep memory allocation bounded and justified by the file-size/record-count policy.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 020: Foundations Project: In-Memory Data Tool
 * Your task: define the empty-input contract and compute the minimum.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

static int min_value(const int *a, size_t n) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    int values[] = {7, 1, 8, 2};
    printf("%d\n", min_value(values, 4));
    return 0;
}

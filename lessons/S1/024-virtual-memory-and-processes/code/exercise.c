/*
 * STUDENT PROBLEM SET — Lesson 024.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Create a Linux program that maps or allocates a buffer and records its initial memory metrics.
 * 2. Touch one byte per page in a selected range, then separately touch the whole range.
 * 3. Record `/proc/self/status` `VmSize` and `VmRSS` before and after each phase; explain what these process metrics indicate and their limitations. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 024: Virtual Memory and Address Spaces
 * Your task: compute how many pages are touched by a byte range.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

static size_t touched_pages(size_t bytes, size_t page_size) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    printf("%zu\n", touched_pages(8192, 4096));
    return 0;
}

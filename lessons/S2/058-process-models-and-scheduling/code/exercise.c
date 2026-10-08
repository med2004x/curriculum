/*
 * STUDENT PROBLEM SET — Lesson 058.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Build the smallest useful version: Use `strace -c` (from Linux tools Lesson 5) on a simple program (e.g. one that reads a file and prints its contents) and identify which specific syscalls correspond to which C library calls you wrote (`fopen`→likely `openat`, `fread`→`read`, `printf`→eventually `write`).
 * 2. Extend it with one explicit boundary case: Note any syscalls you didn't expect (e.g.
 * 3. Rebuild the same behavior with one changed constraint: `mmap` calls from the C library's own startup or buffering). Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 058: Processes, Threads, and Scheduling Revisited
 * Your task: simulate one scheduling tick and state the scheduling policy.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

typedef struct {
    int remaining;
    int period;
} Process;

static void run_one_tick(Process *processes, size_t n) {
    /* TODO: implement the assigned behavior. */
}

int main(void) {
    Process processes[] = {{5, 3}, {2, 4}};
    run_one_tick(processes, 2);
    return 0;
}

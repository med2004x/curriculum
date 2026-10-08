/*
 * STUDENT PROBLEM SET — Lesson 061.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Build the smallest useful version: Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.
 * 2. Extend it with one explicit boundary case: Ensure the child closes the write end when done, and the parent's read loop correctly detects EOF (a `read()` returning 0) and exits its loop rather than blocking forever.
 * 3. Rebuild the same behavior with one changed constraint: Modify step 1 so the parent (consumer) deliberately reads slowly (add a small `sleep()` between reads) while the child (producer) writes a large volume of data quickly with no delay. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 061: IPC, Pipes, and Process Composition
 * Your task: create a pipe, fork, exchange bytes, close unused ends, and wait for the child.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <unistd.h>

int main(void) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

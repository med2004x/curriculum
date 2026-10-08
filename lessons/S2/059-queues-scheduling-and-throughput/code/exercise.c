/*
 * STUDENT PROBLEM SET — Lesson 059.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Build the smallest useful version: Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.
 * 2. Extend it with one explicit boundary case: Simulate SJF (non-preemptive — once a process starts, it runs to completion, but among *ready* processes at each decision point, pick the shortest) for the same workload, and compute the same metrics.
 * 3. Rebuild the same behavior with one changed constraint: Compare average turnaround and response times, and explain why SJF's all-jobs-ready optimality statement does not apply blindly to every arrival schedule. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 059: Queues, Scheduling, and Throughput
 * Your task: implement FIFO behavior and define empty-queue semantics.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

typedef struct {
    int value;
} QueueItem;

static int dequeue(QueueItem *queue, size_t *length, int *out) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    QueueItem queue[8] = {{1}, {2}, {3}};
    size_t length = 3;
    int out = 0;
    printf("ok=%d out=%d length=%zu\n", dequeue(queue, &length, &out), out, length);
    return 0;
}

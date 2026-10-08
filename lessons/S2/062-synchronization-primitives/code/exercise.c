/*
 * STUDENT PROBLEM SET — Lesson 062.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Implement a protected counter and ensure every access follows the mutex protocol.
 * 2. Rewrite the counter as channel-owned state and identify the send/receive synchronization relation.
 * 3. Add a `sync.Cond` predicate wait and explain lock release/reacquisition plus the required loop. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 062: Synchronization Primitives in the Kernel and in Go
 * Your task: protect the shared update with the appropriate synchronization primitive.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <pthread.h>
#include <stdio.h>

static int shared;

static void *worker(void *arg) {
    /* TODO: implement the assigned behavior. */
}

int main(void) {
    pthread_t a, b;
    pthread_create(&a, NULL, worker, NULL);
    pthread_create(&b, NULL, worker, NULL);
    pthread_join(a, NULL);
    pthread_join(b, NULL);
    printf("shared=%d\n", shared);
    return 0;
}

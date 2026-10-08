/*
 * STUDENT PROBLEM SET — Lesson 084.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Draw the order of a WAL record, durable log flush, commit record, and later data-page write.
 * 2. Explain how checkpointing bounds recovery work without replacing WAL.
 * 3. Separately trace the prescribed two-session isolation anomaly and record the exact PostgreSQL result/error. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 084: Transactions, WAL, and Recovery
 * Your task: make replay deterministic, detect a partial tail, and avoid double-apply.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stddef.h>

typedef struct {
    unsigned long lsn;
    int value;
} Record;

static int replay(const Record *log, size_t n, int *state) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    int state = 0;
    Record log[] = {{1, 10}, {2, 5}};
    return replay(log, 2, &state);
}

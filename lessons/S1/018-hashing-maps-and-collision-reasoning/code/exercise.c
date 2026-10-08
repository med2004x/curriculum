/*
 * STUDENT PROBLEM SET — Lesson 018.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Build a small integer-key hash table using an array of buckets.
 * 2. Implement `insert`, `lookup`, and `delete`.
 * 3. Add a collision scenario where at least two keys map to the same slot.
 * 4. Explain what happens in both chaining and open addressing.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 018: Hashing, Maps, and Collision Reasoning
 * Your task: map keys to buckets and explain collisions.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>

static unsigned hash_int(int key, unsigned bucket_count) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    printf("%u\n", hash_int(42, 8));
    return 0;
}

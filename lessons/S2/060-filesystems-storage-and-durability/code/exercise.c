/*
 * STUDENT PROBLEM SET — Lesson 060.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Create a file and hard link; compare inode identity and link count.
 * 2. Remove one name and verify the other still accesses the same file object.
 * 3. Create and remove a symlink target; explain why the link can become dangling. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 060: Filesystems, Storage, and Durability
 * Your task: write safely, flush/sync, then atomically replace the target.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#define _POSIX_C_SOURCE 200809L
#include <stdio.h>

static int durable_write(const char *path, const char *data) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    return durable_write("state.db", "hello\n");
}

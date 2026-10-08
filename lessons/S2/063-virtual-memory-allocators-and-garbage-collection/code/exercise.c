/*
 * STUDENT PROBLEM SET — Lesson 063.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Inspect `/proc/<pid>/maps` for a selected running process and classify several mappings by permissions and backing.
 * 2. Record virtual-size and resident-memory metrics for a process, and explain why they are not interchangeable.
 * 3. Compare C allocation behavior with Go heap/GC observations; state what each tool measures and what it cannot establish. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 063: Virtual Memory, Allocators, and Garbage Collection
 * Your task: align allocations and reject requests that exceed remaining space.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stddef.h>
#include <stdint.h>

typedef struct {
    unsigned char *base;
    size_t size;
    size_t used;
} Arena;

static void *arena_alloc(Arena *arena, size_t size) {
    /* TODO: implement the assigned behavior. */
}

int main(void) {
    unsigned char storage[128];
    Arena arena = {storage, sizeof storage, 0};
    (void)arena_alloc(&arena, 16);
    return 0;
}

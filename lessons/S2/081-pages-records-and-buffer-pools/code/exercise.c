/*
 * STUDENT PROBLEM SET — Lesson 081.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Define a page header and slot record format.
 * 2. Implement insertion and retrieval for variable-length records with bounds checks.
 * 3. Test empty, exact-fit, overflow, deletion, and compaction cases.
 * 4. Add a small buffer-pool model with page ID, pin count, dirty state, and an eviction decision. Do not rely on host struct layout as a portable on-disk format; define byte order and field encoding explicitly.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 081: Pages, Records, and Buffer Pools
 * Your task: enforce bounds and serialize a fixed-size record into the page.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stddef.h>
#include <stdint.h>

typedef struct {
    unsigned char bytes[4096];
} Page;

static int encode_record(Page *page, size_t offset, uint32_t value) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    Page page = {0};
    return encode_record(&page, 0, 42);
}

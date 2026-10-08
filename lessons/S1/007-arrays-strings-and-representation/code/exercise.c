/*
 * STUDENT PROBLEM SET — Lesson 007.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. `array_stats.c`: print array size and element count, then sum five integers.
 * 2. `string_report.c`: report `sizeof`, `strlen`, and the final-byte value for a C string.
 * 3. `safe_index.c`: reject indexes outside the valid range before accessing the array.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 007: Arrays, Strings, and Representation
 * Your task: walk the string without reading past the terminator.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

static size_t count_vowels(const char *s) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    const char *input = "systems";
    printf("%zu\n", count_vowels(input));
    return 0;
}

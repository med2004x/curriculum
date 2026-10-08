/*
 * STUDENT PROBLEM SET — Lesson 015.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Write a program that triggers signed overflow.
 * 2. Compile it without a sanitizer and record the output.
 * 3. Compile it with `-fsanitize=undefined` and record the diagnostic.
 * 4. Repeat the same exercise with a null pointer or invalid index case if you want a second example. Do not copy a finished answer before you have produced your own minimal example.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 015: Undefined Behavior, Sanitizers, and Defensive Debugging
 * Your task: detect overflow before performing the addition.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <limits.h>
#include <stdio.h>

static int safe_add(int a, int b, int *out) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    int out = 0;
    printf("ok=%d out=%d\n", safe_add(40, 2, &out), out);
    return 0;
}

/*
 * STUDENT PROBLEM SET — Lesson 004.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. `truth.c`: print the full truth table for `(A && B) || !C`, all 8 rows, using nested loops or a counter over 0..7.
 * 2. `demorgan.c`: for all four `(A, B)` pairs, assert `!(A && B) == (!A || !B)` and the dual law; print `ok` only if every row passes.
 * 3. `classify.c`: given an integer read from `argv[1]` (use `atoi`), print whether it is negative, zero, or positive, and whether it is even or odd, correctly for negative numbers.
 * 4. `leap.c`: a year is a leap year if divisible by 4, except centuries, unless divisible by 400. Write the condition, then test 1900, 2000, 2023, 2024. State your test cases *before* running.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 004: Conditions, Boolean Logic, and Proof Thinking
 * Your task: implement the condition logic from the lesson.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>

static int classify(int year) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    int year = 2024;
    printf("%d\n", classify(year));
    return 0;
}

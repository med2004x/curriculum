/*
 * STUDENT PROBLEM SET — Lesson 010.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. `main.c`, `util.c`, `util.h`: define `add`, declare it in the header, and build the executable in explicit stages.
 * 2. `macro.c`: verify `BAD_SQ(1+2) = 5` and `SQ(1+2) = 9`.
 * 3. `symbols.md`: record `nm` results and explain every relevant `U` and `T` symbol.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 010: Compilation, Linking, and Toolchains
 * Your task: keep this declaration, then supply the definition from exercise_math.c.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>

/* TODO: keep this declaration, then supply the definition from exercise_math.c. */
int add(int a, int b);

int main(void) {
    printf("%d\n", add(7, 5));
    return 0;
}

/*
 * STUDENT PROBLEM SET — Lesson 013.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Declare one initialized global and one zero-initialized global; predict their initial values.
 * 2. Add a function with a `static` local and explain why its address remains valid after return.
 * 3. Add one automatic local and one checked `malloc` allocation; print all addresses using `%p` and cast to `(void *)`.
 * 4. Free the heap object exactly once and never read it afterward.
 * 5. Build the separate invalid-lifetime experiment with AddressSanitizer and save the compiler warning and runtime report, if produced. Use a debugger breakpoint inside the function if available. Inspect the frame and locals, then distinguish debugger observations from guarantees made by the C language.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 013: Pointers, Addresses, and Memory Layout
 * Your task: swap the pointed-to integers without losing either value.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>

static void swap(int *a, int *b) {
    /* TODO: implement the assigned behavior. */
}

int main(void) {
    int a = 10, b = 20;
    swap(&a, &b);
    printf("%d %d\n", a, b);
    return 0;
}

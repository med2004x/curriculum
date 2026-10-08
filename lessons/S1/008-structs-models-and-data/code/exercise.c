/*
 * STUDENT PROBLEM SET — Lesson 008.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. `order.c`: model an order with an enum status and print three states.
 * 2. `layout.c`: measure the verified struct layouts with `sizeof` and `offsetof`.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 008: Structs, Enums, and Modeling Data
 * Your task: return the maximum score and define the empty-input behavior.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>
#include <stddef.h>

typedef struct {
    char name[32];
    int score;
} Student;

static int best_score(const Student *students, size_t n) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    Student students[] = {{"A", 72}, {"B", 91}, {"C", 84}};
    printf("%d\n", best_score(students, 3));
    return 0;
}

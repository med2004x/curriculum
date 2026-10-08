#include <stdio.h>
#include <stddef.h>

typedef struct {
    char name[32];
    int score;
} Student;

static int best_score(const Student *students, size_t n) {
    /* TODO: return the maximum score and define the empty-input behavior. */
    return 0;
}

int main(void) {
    Student students[] = {{"A", 72}, {"B", 91}, {"C", 84}};
    printf("%d\n", best_score(students, 3));
    return 0;
}

#include <stdio.h>
#include <stddef.h>

static int min_value(const int *a, size_t n) {
    /* TODO: define the empty-input contract and compute the minimum. */
    (void)a; (void)n;
    return 0;
}

int main(void) {
    int values[] = {7, 1, 8, 2};
    printf("%d\n", min_value(values, 4));
    return 0;
}

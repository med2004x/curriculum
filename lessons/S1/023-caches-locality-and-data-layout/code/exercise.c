#include <stdio.h>
#include <stddef.h>

static long row_major_sum(const int *a, size_t rows, size_t cols) {
    /* TODO: implement contiguous row-major traversal. */
    (void)a; (void)rows; (void)cols;
    return 0;
}

int main(void) {
    int a[2][3] = {{1,2,3},{4,5,6}};
    printf("%ld\n", row_major_sum(&a[0][0], 2, 3));
    return 0;
}

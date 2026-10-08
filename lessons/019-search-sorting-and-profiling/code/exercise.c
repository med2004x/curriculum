#include <stdio.h>
#include <stddef.h>

static void insertion_sort(int *a, size_t n) {
    /* TODO: sort in place and state the loop invariant. */
    (void)a; (void)n;
}

int main(void) {
    int a[] = {5, 2, 4, 1, 3};
    insertion_sort(a, 5);
    for (size_t i = 0; i < 5; i++) printf("%d%c", a[i], i == 4 ? '\n' : ' ');
    return 0;
}

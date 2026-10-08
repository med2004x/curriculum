#include <stdio.h>

static void swap(int *a, int *b) {
    /* TODO: swap the pointed-to integers without losing either value. */
    (void)a;
    (void)b;
}

int main(void) {
    int a = 10, b = 20;
    swap(&a, &b);
    printf("%d %d\n", a, b);
    return 0;
}

#include <stdio.h>

int add(int a, int b) {
    /* CHALLENGE TODO: move the implementation to a separate translation unit. */
    return a + b;
}

int main(void) {
    printf("%d\n", add(7, 5));
    return 0;
}

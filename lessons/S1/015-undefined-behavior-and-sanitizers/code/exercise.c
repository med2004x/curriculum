#include <limits.h>
#include <stdio.h>

static int safe_add(int a, int b, int *out) {
    /* TODO: detect overflow before performing the addition. */
    (void)a; (void)b; (void)out;
    return 0;
}

int main(void) {
    int out = 0;
    printf("ok=%d out=%d\n", safe_add(40, 2, &out), out);
    return 0;
}

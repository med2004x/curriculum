#include <stdio.h>

static int compute(int x) {
    /* TODO: build with optimisation on/off and inspect the generated assembly. */
    return x * 3 + 1;
}

int main(void) {
    printf("%d\n", compute(7));
    return 0;
}

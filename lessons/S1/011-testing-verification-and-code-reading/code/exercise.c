#include <stdio.h>

static int score(int total, int passed) {
    /* TODO: implement the contract and test boundary cases. */
    return total + passed;
}

int main(void) {
    printf("%d\n", score(10, 3));
    return 0;
}

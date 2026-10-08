#include <stdio.h>

static int logic_xor(int a, int b) {
    /* TODO: implement the gate using only the allowed primitive operations. */
    return a ^ b;
}

int main(void) {
    for (int a = 0; a <= 1; ++a)
        for (int b = 0; b <= 1; ++b)
            printf("%d xor %d = %d\n", a, b, logic_xor(a, b));
    return 0;
}

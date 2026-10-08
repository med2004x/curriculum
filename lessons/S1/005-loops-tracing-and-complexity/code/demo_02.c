#include <stdio.h>

int main(void) {
    int n = 0;
    while (n > 0) puts("while body");
    do { puts("do body"); } while (n > 0);
    return 0;
}

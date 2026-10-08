#include <stdio.h>

void swap_bad(int a, int b) {
    int tmp = a;
    a = b;
    b = tmp;
}

int main(void) {
    int x = 1;
    int y = 2;
    swap_bad(x, y);
    printf("%d %d\n", x, y);
    return 0;
}

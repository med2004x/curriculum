#include <limits.h>
#include <stdio.h>

int main(void) {
    int x = INT_MAX;
    int y = x + 1; // UB: signed overflow
    printf("%d\n", y);
    return 0;
}

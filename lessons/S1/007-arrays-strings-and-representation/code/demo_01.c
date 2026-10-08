#include <stdio.h>

int main(void) {
    int a[5] = {10, 20, 30, 40, 50};
    printf("%zu\n", sizeof a);
    printf("%zu\n", sizeof a / sizeof a[0]);
    return 0;
}

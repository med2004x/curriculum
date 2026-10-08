#include <stdio.h>

int main(void) {
    int sum = 0;
    for (int i = 1; i <= 5; i++) {
        sum += i;
        printf("%d %d\n", i, sum);
    }
    return 0;
}

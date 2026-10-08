#include <stdio.h>

int main(void) {
    for (int a = 0; a <= 1; a++) {
        for (int b = 0; b <= 1; b++) {
            for (int c = 0; c <= 1; c++) {
                int result = (a && b) || !c;
                printf("%d %d %d -> %d\n", a, b, c, result);
            }
        }
    }
    return 0;
}

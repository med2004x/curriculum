#include <assert.h>
#include <stdio.h>

int main(void) {
    for (int a = 0; a <= 1; a++) {
        for (int b = 0; b <= 1; b++) {
            assert((!(a && b)) == (!a || !b));
            assert((!(a || b)) == (!a && !b));
        }
    }
    puts("ok");
    return 0;
}

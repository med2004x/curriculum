#include <stdio.h>

#define BAD_SQ(x) x*x
#define SQ(x) ((x)*(x))

int main(void) {
    printf("%d %d\n", BAD_SQ(1 + 2), SQ(1 + 2));
    return 0;
}

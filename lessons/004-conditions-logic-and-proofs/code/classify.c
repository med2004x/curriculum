#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
    if (argc != 2) {
        fprintf(stderr, "usage: %s INTEGER\n", argv[0]);
        return 2;
    }
    int n = atoi(argv[1]);
    if (n < 0) puts("negative odd/even:");
    else if (n == 0) puts("zero even");
    else puts("positive odd/even:");
    printf("%s %s\n", n < 0 ? "negative" : n > 0 ? "positive" : "zero", n % 2 == 0 ? "even" : "odd");
    return 0;
}

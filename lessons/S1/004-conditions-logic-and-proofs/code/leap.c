#include <stdio.h>
#include <stdlib.h>

static int is_leap(int year) {
    return year % 400 == 0 || (year % 4 == 0 && year % 100 != 0);
}

int main(int argc, char **argv) {
    if (argc != 2) {
        fprintf(stderr, "usage: %s YEAR\n", argv[0]);
        return 2;
    }
    int year = atoi(argv[1]);
    printf("%d: %s\n", year, is_leap(year) ? "leap" : "common");
    return 0;
}

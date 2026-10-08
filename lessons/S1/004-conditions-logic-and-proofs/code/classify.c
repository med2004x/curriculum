/*
 * STUDENT TASK: given one integer in argv[1], print its sign and parity.
 * Use labels `negative`, `zero`, `positive` and `odd`, `even` separated by
 * one space. Remember that C's remainder for negative values may be negative.
 */
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
    if (argc != 2) {
        fprintf(stderr, "usage: %s INTEGER\n", argv[0]);
        return 2;
    }
    int n = atoi(argv[1]);
    /* TODO: classify n and print one exact label pair. */
    return 0;
}

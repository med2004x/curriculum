#include <errno.h>
#include <stdio.h>
#include <string.h>

int main(void) {
    FILE *fp = fopen("nope.txt", "r");
    if (fp == NULL) {
        fprintf(stderr, "fopen(nope.txt): %s (errno=%d)\n", strerror(errno), errno);
        return 1;
    }
    fclose(fp);
    return 0;
}

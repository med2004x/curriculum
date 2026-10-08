#define _POSIX_C_SOURCE 200809L
#include <stdio.h>

static int durable_write(const char *path, const char *data) {
    /* TODO: write safely, flush/sync, then atomically replace the target. */
    (void)path;
    (void)data;
    return 0;
}

int main(void) {
    return durable_write("state.db", "hello\n");
}

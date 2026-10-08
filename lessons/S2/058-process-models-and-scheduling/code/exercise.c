#include <stdio.h>
#include <stddef.h>

typedef struct {
    int remaining;
    int period;
} Process;

static void run_one_tick(Process *processes, size_t n) {
    /* TODO: simulate one scheduling tick and state the scheduling policy. */
    (void)processes;
    (void)n;
}

int main(void) {
    Process processes[] = {{5, 3}, {2, 4}};
    run_one_tick(processes, 2);
    return 0;
}

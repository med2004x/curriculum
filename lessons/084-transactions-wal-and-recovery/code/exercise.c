#include <stddef.h>

typedef struct {
    unsigned long lsn;
    int value;
} Record;

static int replay(const Record *log, size_t n, int *state) {
    /* TODO: make replay deterministic, detect a partial tail, and avoid double-apply. */
    (void)log;
    (void)n;
    (void)state;
    return 0;
}

int main(void) {
    int state = 0;
    Record log[] = {{1, 10}, {2, 5}};
    return replay(log, 2, &state);
}

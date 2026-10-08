#include <stddef.h>
#include <stdlib.h>

typedef struct {
    int *data;
    size_t len;
    size_t cap;
} Vector;

static int vector_push(Vector *v, int value) {
    /* TODO: grow when required, preserve existing elements, and update len. */
    (void)v;
    (void)value;
    return 0;
}

static void vector_free(Vector *v) {
    /* TODO: release owned memory and leave the vector in a safe state. */
    (void)v;
}

int main(void) {
    Vector v = {0};
    vector_push(&v, 42);
    vector_free(&v);
    return 0;
}

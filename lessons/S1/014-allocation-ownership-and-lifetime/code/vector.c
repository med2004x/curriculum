/* STUDENT TASK: implement this file from the assignment in WORK.md. */
#include "vector.h"
#include <stdint.h>
#include <stdlib.h>

int vector_push(Vector *vector, int value) {
    /* TODO: grow safely, preserve data, then update len. */
    (void)vector;
    (void)value;
    return 0;
}

int vector_get(const Vector *vector, size_t index, int *out) {
    /* TODO: reject invalid input and copy the requested value to out. */
    (void)vector;
    (void)index;
    (void)out;
    return 0;
}

void vector_free(Vector *vector) {
    /* TODO: free the owned buffer and reset the owner to its empty state. */
    (void)vector;
}

#include "vector.h"
#include <stdint.h>
#include <stdlib.h>

int vector_push(Vector *vector, int value) {
    if (vector == NULL) return 0;
    if (vector->len == vector->cap) {
        size_t next;
        if (vector->cap == 0) {
            next = 4;
        } else {
            if (vector->cap > SIZE_MAX / 2) return 0;
            next = vector->cap * 2;
        }
        if (next > SIZE_MAX / sizeof *vector->data) return 0;
        int *grown = realloc(vector->data, next * sizeof *vector->data);
        if (grown == NULL) return 0;
        vector->data = grown;
        vector->cap = next;
    }
    vector->data[vector->len] = value;
    vector->len++;
    return 1;
}

int vector_get(const Vector *vector, size_t index, int *out) {
    if (vector == NULL || out == NULL || index >= vector->len) return 0;
    *out = vector->data[index];
    return 1;
}

void vector_free(Vector *vector) {
    if (vector == NULL) return;
    free(vector->data);
    vector->data = NULL;
    vector->len = 0;
    vector->cap = 0;
}

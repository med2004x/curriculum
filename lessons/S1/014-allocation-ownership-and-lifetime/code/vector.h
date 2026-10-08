#ifndef VECTOR_H
#define VECTOR_H
#include <stddef.h>

typedef struct {
    int *data;
    size_t len;
    size_t cap;
} Vector;

int vector_push(Vector *vector, int value);
int vector_get(const Vector *vector, size_t index, int *out);
void vector_free(Vector *vector);

#endif

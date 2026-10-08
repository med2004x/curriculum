#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int *data;
    size_t len;
    size_t cap;
} Vector;

static int vector_push(Vector *v, int value) {
    if (v->len == v->cap) {
        size_t next = v->cap == 0 ? 4 : v->cap * 2;
        int *grown = realloc(v->data, next * sizeof *grown);
        if (grown == NULL) return 0;
        v->data = grown;
        v->cap = next;
    }
    v->data[v->len++] = value;
    return 1;
}

static void vector_free(Vector *v) {
    free(v->data);
    *v = (Vector){0};
}

int main(void) {
    Vector v = {0};
    vector_push(&v, 10);
    vector_push(&v, 20);
    vector_push(&v, 30);
    for (size_t i = 0; i < v.len; ++i) printf("%d%c", v.data[i], i + 1 == v.len ? '\n' : ' ');
    vector_free(&v);
    return 0;
}

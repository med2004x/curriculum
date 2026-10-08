/*
 * TEST SUITE: run with the sanitizer command in WORK.md.
 * This is the teacher; implement the student code in vector.c/vector.h.
 */
#include "vector.h"
#include <assert.h>

int main(void) {
    Vector values = {0};
    for (int i = 0; i < 9; i++) assert(vector_push(&values, i * 10));
    assert(values.len == 9);
    assert(values.cap >= values.len);
    for (size_t i = 0; i < values.len; i++) {
        int value = -1;
        assert(vector_get(&values, i, &value));
        assert(value == (int)i * 10);
    }
    int ignored;
    assert(!vector_get(&values, values.len, &ignored));
    vector_free(&values);
    assert(values.data == NULL && values.len == 0 && values.cap == 0);
    vector_free(&values);
    return 0;
}

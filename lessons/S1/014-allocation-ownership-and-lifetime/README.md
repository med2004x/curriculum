# Lesson 014 — Allocation, Ownership, and Lifetime

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Define ownership before writing allocation code

The vector owns its backing array. `vector_push` may replace that array; callers use the vector API and do not free `data` directly. `vector_free` releases the array exactly once. A successful `realloc` may move the array, so pointers into `data` must not be kept across a push that can grow it.

Maintain this invariant after every public operation: `len <= cap`; `data == NULL` when `cap == 0`; otherwise `data` points to storage for at least `cap` integers. A failed push leaves the prior vector valid and unchanged.

## 2. The interface and implementation

Save the shared type and declarations in `vector.h`:

```c
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
```

Implement the ownership rules in `vector.c`:

```c
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
```

Never write `vector->data = realloc(vector->data, bytes)` directly. If growth fails, `realloc` returns `NULL` and the original allocation is still live; overwriting the only pointer would leak it. Store the result in `grown` and update the vector only after success.

## 3. Test the boundary and the invariant

The initial capacity is 4. Pushing the fifth element forces growth; pushing nine forces another capacity change. Check values after each growth because a moved buffer must preserve all prior elements.

```c
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
```

The final second `vector_free` is safe because the API resets the owner to its empty state. It is not a double-free: the second call passes `NULL` to `free`.

## 4. Compile and verify

Build the test and implementation together:

```sh
gcc -std=c17 -Wall -Wextra -Wpedantic -fsanitize=address,undefined -g test_vector.c vector.c -o test_vector
./test_vector
echo $?
```

Expected: no assertion or sanitizer report, then status 0. Leak detection is platform-dependent; where supported, run with `ASAN_OPTIONS=detect_leaks=1`. Record compiler and sanitizer versions because diagnostics and leak support vary.

Trace the fifth push: old `cap` is 4, `next` becomes 8, allocation requests `8 * sizeof(int)` bytes, and only after successful `realloc` do `data` and `cap` change. `len` changes from 4 to 5 after the new value is stored.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

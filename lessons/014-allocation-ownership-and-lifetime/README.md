# Lesson 014 — Allocation, Ownership, and Lifetime

## Threads

**Primary:** C

**Supporting:** `malloc`; `realloc`; `free`; ownership; capacity; failure handling. **Math companion:** geometric growth and amortized work

**Time:** 3–4 hours. **Prerequisites:** Lesson 013.

## What you will be able to do

1. State who owns a dynamically allocated buffer and who must free it.
2. Grow a vector without losing its old allocation when `realloc` fails.
3. Enforce the invariant `0 <= len <= cap` and check indexed reads.
4. Test growth, bounds, allocation cleanup, and failure paths.
5. Distinguish a memory leak, double-free, and use-after-free.

## Retrieve first (no notes)

1. When does an automatic local from Lesson 013 stop being alive?
2. What must be true before a pointer may be dereferenced?
3. What did the address program show about `malloc` storage versus static storage?
4. What can AddressSanitizer report, and what does no report fail to prove?

Write answers first. Lesson 013 covered object lifetime; this lesson applies the same rule to storage whose lifetime is controlled by `malloc` and `free`.

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

## Predict, then observe

1. Predict `len` and `cap` after pushes one through five.
2. Predict the result of reading index 8 after nine pushes.
3. Predict the result of reading index `len`.
4. Predict the vector state if `realloc` returns `NULL`.
5. Predict why the local variable `grown` is necessary.
6. Predict whether calling the correct `vector_free` twice frees twice.

Write the invariant values before running the test. Do not assume that `realloc` moves the buffer on every growth; correctness must hold whether it moves or extends in place.

## Build from blank

1. Define `Vector` and its invariant in a header.
2. Implement push from an empty vector, then add growth by doubling capacity.
3. Add bounds-checked `vector_get` and an idempotent cleanup function.
4. Push nine known values, verify every value, and check index `len` fails.
5. Compile with both sanitizers and save the exact command and exit status.

Document that the vector owns its buffer and that callers must not shallow-copy a populated `Vector`; two copies would both appear to own the same pointer and could free it twice.

## Failure drill

1. Remove the `vector_free` call in a test scope and run LeakSanitizer where available; identify the allocation that remains reachable only as a leak.
2. In a disposable broken copy, remove the `data = NULL` reset and call `vector_free` twice; run under AddressSanitizer and record the double-free report.
3. Assign `realloc` directly to `vector->data`; explain how allocation failure loses the owner’s only pointer.
4. Read index `len` without checking the return value; use the invariant to explain why it is out of bounds.

Do not rely on the unsanitized program crashing. Undefined behavior and allocator diagnostics vary; restore the correct implementation and rerun the clean test afterward.

## Transfer

Change the vector element type from `int` to a small `struct Point { int x; int y; }`. Recalculate the allocation size using `sizeof *vector->data`, preserve the same owner/invariant rules, and test growth across capacity 4. Explain why the byte count changes but the capacity is still measured in elements.

## Retrieval questions

1. Which object owns the allocation after a successful `vector_push`?
2. Why must `realloc` return into a temporary pointer?
3. What does the invariant say about `len`, `cap`, and `data`?
4. Why is the `vector_free` function safe to call twice?
5. What bug does LeakSanitizer detect in the first failure drill?
6. Why is a shallow copy of a populated `Vector` an ownership error?

## Evidence to record

Create `lesson014/EVIDENCE.md` with:

```text
Compiler and sanitizer versions:

Build command and clean-run status:

Capacity trace through push five:

Leak report or platform limitation:

Double-free report:

Ownership rule in one sentence:
```

## Common mistakes

- Losing the original pointer when `realloc` fails.
- Multiplying capacity without checking for `size_t` overflow.
- Forgetting that a successful `realloc` can invalidate interior pointers.
- Freeing the buffer in both the vector and its caller.
- Keeping a shallow copy that appears to own the same allocation.
- Treating sanitizer silence as proof that ownership is correct.

## Precision checks

1. State who allocates and who frees the buffer.
2. Explain how a failed `realloc` leaves the original block allocated.
3. Show why `index == len` is invalid even when `cap > len`.
4. Distinguish a leak from a double-free by the invalid ownership transition.
5. Explain what the sanitizer proves and what remains the program's responsibility.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can implement the vector from blank, preserve its invariant across both growth and allocation failure, test its bounds and cleanup, and explain who owns each allocation. Reproduce a leak and a double-free with sanitizer evidence, then restore the clean implementation and run the tests again.

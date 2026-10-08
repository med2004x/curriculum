#include <stddef.h>
#include <stdint.h>

typedef struct {
    unsigned char *base;
    size_t size;
    size_t used;
} Arena;

static void *arena_alloc(Arena *arena, size_t size) {
    /* TODO: align allocations and reject requests that exceed remaining space. */
    (void)arena;
    (void)size;
    return NULL;
}

int main(void) {
    unsigned char storage[128];
    Arena arena = {storage, sizeof storage, 0};
    (void)arena_alloc(&arena, 16);
    return 0;
}

#include <stdio.h>

typedef struct {
    int frame[4];
    unsigned valid;
} BufferPool;

static int fetch(BufferPool *pool, int page_id) {
    /* TODO: implement hit/miss tracking and justify a replacement policy. */
    (void)pool;
    (void)page_id;
    return -1;
}

int main(void) {
    BufferPool pool = {0};
    printf("frame=%d\n", fetch(&pool, 7));
    return 0;
}

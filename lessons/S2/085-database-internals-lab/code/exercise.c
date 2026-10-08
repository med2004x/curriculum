/* STUDENT TASK: implement the four-frame fetch behavior described in WORK.md. */
#include <stdio.h>

typedef struct {
    int frame[4];
    unsigned valid;
    unsigned hits;
    unsigned misses;
    unsigned next_victim;
} BufferPool;

static int fetch(BufferPool *pool, int page_id) {
    /* TODO: find a hit or load the page and update the pool state. */
    return 0;
}

int main(void) {
    BufferPool pool = {0};
    printf("frame=%d\n", fetch(&pool, 7));
    printf("frame=%d\n", fetch(&pool, 7));
    printf("hits=%u misses=%u\n", pool.hits, pool.misses);
    return 0;
}

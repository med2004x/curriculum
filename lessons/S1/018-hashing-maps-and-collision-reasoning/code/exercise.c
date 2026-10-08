#include <stdio.h>

static unsigned hash_int(int key, unsigned bucket_count) {
    /* TODO: map keys to buckets and explain collisions. */
    (void)bucket_count;
    return (unsigned)key;
}

int main(void) {
    printf("%u\n", hash_int(42, 8));
    return 0;
}

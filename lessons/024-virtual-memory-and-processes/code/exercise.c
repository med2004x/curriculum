#include <stdio.h>
#include <stddef.h>

static size_t touched_pages(size_t bytes, size_t page_size) {
    /* TODO: compute how many pages are touched by a byte range. */
    (void)bytes; (void)page_size;
    return 0;
}

int main(void) {
    printf("%zu\n", touched_pages(8192, 4096));
    return 0;
}

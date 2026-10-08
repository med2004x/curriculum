#include <stddef.h>
#include <stdint.h>

typedef struct {
    unsigned char bytes[4096];
} Page;

static int encode_record(Page *page, size_t offset, uint32_t value) {
    /* TODO: enforce bounds and serialize a fixed-size record into the page. */
    (void)page;
    (void)offset;
    (void)value;
    return 0;
}

int main(void) {
    Page page = {0};
    return encode_record(&page, 0, 42);
}

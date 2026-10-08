#include <assert.h>
#include <stdint.h>
#include <string.h>

#define main lesson_demo_main
#include "exercise.c"
#undef main

int main(void) {
    Page page = {{0}};
    const uint32_t value = UINT32_C(0x12345678);
    assert(encode_record(&page, 12, value) == 0);
    assert(memcmp(page.bytes + 12, &value, sizeof value) == 0);

    unsigned char before[sizeof page.bytes];
    memcpy(before, page.bytes, sizeof before);
    assert(encode_record(&page, sizeof page.bytes - sizeof value + 1, value) < 0);
    assert(memcmp(page.bytes, before, sizeof before) == 0);
    return 0;
}

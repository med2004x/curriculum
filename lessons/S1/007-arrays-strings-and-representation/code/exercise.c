#include <stdio.h>
#include <stddef.h>

static size_t count_vowels(const char *s) {
    /* TODO: walk the string without reading past the terminator. */
    return 0;
}

int main(void) {
    const char *input = "systems";
    printf("%zu\n", count_vowels(input));
    return 0;
}

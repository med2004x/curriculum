#include <stdio.h>
#include <stdlib.h>

int initialized_global = 17;
int zero_global;

static int *static_object(void) {
    static int value = 23;
    return &value;
}

int main(void) {
    int automatic = 31;
    int *heap = malloc(sizeof *heap);
    if (heap == NULL) {
        perror("malloc");
        return 1;
    }
    *heap = 47;
    printf("global=%p value=%d\n", (void *)&initialized_global, initialized_global);
    printf("zero=%p value=%d\n", (void *)&zero_global, zero_global);
    printf("static=%p value=%d\n", (void *)static_object(), *static_object());
    printf("automatic=%p value=%d\n", (void *)&automatic, automatic);
    printf("heap=%p value=%d\n", (void *)heap, *heap);
    free(heap);
    return 0;
}

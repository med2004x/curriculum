#include <stdio.h>
#include <stddef.h>

typedef struct {
    int value;
} QueueItem;

static int dequeue(QueueItem *queue, size_t *length, int *out) {
    /* TODO: implement FIFO behavior and define empty-queue semantics. */
    (void)queue;
    (void)length;
    (void)out;
    return 0;
}

int main(void) {
    QueueItem queue[8] = {{1}, {2}, {3}};
    size_t length = 3;
    int out = 0;
    printf("ok=%d out=%d length=%zu\n", dequeue(queue, &length, &out), out, length);
    return 0;
}

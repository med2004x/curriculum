#include <pthread.h>
#include <stdio.h>

static int shared;

static void *worker(void *arg) {
    /* TODO: protect the shared update with the appropriate synchronization primitive. */
    (void)arg;
    shared++;
    return NULL;
}

int main(void) {
    pthread_t a, b;
    pthread_create(&a, NULL, worker, NULL);
    pthread_create(&b, NULL, worker, NULL);
    pthread_join(a, NULL);
    pthread_join(b, NULL);
    printf("shared=%d\n", shared);
    return 0;
}

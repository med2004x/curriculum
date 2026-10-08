/* TEST HARNESS: A supplied check driver; keep it unchanged while implementing the task. */
#include <stdio.h>
#include <stdlib.h>

enum Result { RESULT_INVALID = 0, RESULT_FAIL = 1, RESULT_PASS = 2 };

enum Result classify(int score) {
    if (score < 0 || score > 100) return RESULT_INVALID;
    return score >= 50 ? RESULT_PASS : RESULT_FAIL;
}

int main(int argc, char **argv) {
    if (argc != 2) return 2;
    int score = atoi(argv[1]);
    printf("%d\n", classify(score));
    return 0;
}

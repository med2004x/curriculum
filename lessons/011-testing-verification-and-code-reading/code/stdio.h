#include <stdio.h>

#ifndef PASS_MARK
#define PASS_MARK 50
#endif

enum result { INVALID = -1, FAIL = 0, PASS = 1 };

static enum result classify(int score) {
    if (score < 0 || score > 100) {
        return INVALID;
    }
    if (score >= PASS_MARK) {
        return PASS;
    }
    return FAIL;
}

static int check_case(int score, enum result expected) {
    enum result actual = classify(score);
    if (actual != expected) {
        printf("FAIL score=%d expected=%d actual=%d\n",
               score, expected, actual);
        return 1;
    }
    printf("PASS score=%d result=%d\n", score, actual);
    return 0;
}

int main(void) {
    int failures = 0;
    failures += check_case(0, FAIL);
    failures += check_case(49, FAIL);
    failures += check_case(50, PASS);
    failures += check_case(100, PASS);
    failures += check_case(-1, INVALID);
    failures += check_case(101, INVALID);
    return failures != 0;
}

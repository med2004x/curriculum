# Lesson 011 — Testing, Verification, and Reading Code

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. A test is a checkable contract

A test case has four parts: a chosen input, an expected result, the actual result, and a rule that decides whether they agree. “Try some scores” is not a test plan until each value and expected classification is written down.

We will test this contract:

- Scores from 0 through 100 are valid.
- A valid score below 50 is `FAIL`.
- A valid score from 50 through 100 is `PASS`.
- Values outside that range are `INVALID`.

The endpoints matter. Testing 20 and 80 alone would not reveal whether 50 is included in the passing interval or whether 100 is accepted.

## 2. Design cases before reading the implementation

| Input | Reason for choosing it | Expected result |
|---:|---|---|
| 0 | Lowest valid score | FAIL |
| 49 | Immediately below the decision boundary | FAIL |
| 50 | At the decision boundary | PASS |
| 100 | Highest valid score | PASS |
| -1 | Just outside the lower bound | INVALID |
| 101 | Just outside the upper bound | INVALID |

This is boundary-value analysis: the cases cluster around places where the contract changes. The inputs 49 and 50 are more diagnostic than two arbitrary values well inside each region.

Before running anything, predict the result for each row. In particular, decide whether the checks should accept 100 and reject 101. If either prediction is wrong, the contract is not yet clear enough to test.

## 3. Read the implementation as a trace

The function first rejects values outside the valid interval. Only then does it compare a valid score with the passing threshold. For input 50, the range check is false, the threshold check is true, and the result is `PASS`. For input 49, both range checks are false, while the threshold check is false, so the result is `FAIL`.

```c
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
```

The enum values make the printed evidence readable: `INVALID` is -1, `FAIL` is 0, and `PASS` is 1. The process exit status has a different role: `main` returns 0 only when every check passed, and 1 when at least one check failed.

## 4. Run the checks and interpret the evidence

Save the code as `score_test.c`. From its directory, build and run it:

```sh
gcc -std=c17 -Wall -Wextra -Wpedantic -g score_test.c -o score_test
./score_test
echo $?
```

Expected: six `PASS` lines, followed by exit status `0`. The status is the status of `./score_test`; do not run another command before `echo $?`, because `$?` always refers to the immediately preceding command.

Now change the comparison from `score >= PASS_MARK` to `score > PASS_MARK`. Rebuild and run again. The test for 50 must print a mismatch (`expected=1 actual=0`), and the program must exit nonzero. Restore the original comparison after recording the failure.

This is a regression test: it catches a specific defect at the boundary. The test is useful because it fails for the changed implementation and passes for the contracted one, not because the output looks reassuring.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

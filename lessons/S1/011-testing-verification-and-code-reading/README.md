# Lesson 011 — Testing, Verification, and Reading Code

## Threads

**Primary:** C

**Supporting:** test cases; assertions; code reading; evidence

**Math companion:** logic review

**Time:** 3-4 hours. **Prerequisites:** Lesson 010.

## What you will be able to do

1. Turn an input/output contract into normal, boundary, and invalid test cases.
2. Trace a C function on paper before compiling it.
3. Read a failing test as evidence about a specific expected and actual value.
4. Distinguish a test failure from a compiler warning, linker error, and runtime failure.

## Retrieve first (no notes)

1. What file does `gcc -E main.c -o main.i` produce, and is it an executable?
2. What does `U add` in `nm main.o` say about that object file?
3. Why can `gcc main.o -o app` fail even when compiling `main.c` succeeded?
4. Name one difference between a warning and an unresolved-symbol linker failure.

Write these from memory before continuing. Check your answers against Lesson 010 after finishing today's first test run.

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

## Predict, then observe

1. Predict every expected classification in the table before running the program.
2. Predict the actual result at 50 after changing `>=` to `>`.
3. Predict the process exit status when one check fails.
4. Predict whether `101` reaches the threshold comparison or returns at the range check.
5. Predict what happens if the test for 49 is deleted while the implementation wrongly uses `>`.
6. Predict which evidence distinguishes a bad classification from a failed compilation.
7. Run the original and mutated versions; record prediction beside observation.

For item 6, the compiled test's `FAIL score=...` line means the program ran and compared values. A compiler diagnostic means no such test result has yet been produced.

## Build from blank

1. In a fresh directory, write the enum and `classify` function from the contract without copying the code block.
2. Write a test function that compares actual and expected values and counts mismatches.
3. Add all six cases from the table, including both sides of the threshold and both invalid inputs.
4. Make the executable return zero only if the mismatch count is zero.
5. Compile with the command above, run it, and record the process status.

Before compiling, trace scores 49, 50, 100, and 101 by hand. Your trace should identify the branch each input takes, not only its final label.

## Failure drill

1. Change `>=` to `>` and confirm that the score-50 case catches it.
2. Change the expected result for 49 from `FAIL` to `PASS`; confirm that the test harness reports a mismatch even though the implementation is unchanged.
3. Remove the lower-bound check; confirm that -1 is no longer classified as `INVALID`.
4. For each failure, save the diagnostic line and exit status before restoring the source.

Classify each observation: compiler diagnostic, test assertion/mismatch, or process status. They answer different questions. A clean compile proves neither that the contract is correct nor that every test passes.

## Transfer

Change the contract so a valid score passes at 60 instead of 50. Keep the valid range at 0 through 100. From a blank copy, choose cases 59, 60, and 61 plus both range endpoints and both invalid neighbors. Write expected outcomes before editing the code.

Do not merely change the macro and reuse the original tests: add the new boundary pair so the suite would reject both `score >= 59` and `score > 60`. Explain which case catches each defect, then run the suite and record its exit status.

## Retrieval questions

1. State the valid score range and the exact passing rule.
2. Why are 49 and 50 a more useful pair than 20 and 80 for checking the threshold?
3. What is the expected result for 100? For 101?
4. What does the check-case function compare?
5. What process status means that every test passed in this program?
6. How does the score-50 test expose the `>` mutation?

## Evidence to record

Create `lesson011/EVIDENCE.md` with these fields:

```text
Prediction for six cases:

Build command:

Observed output summary:

Exit status for original:

Mutation made and failing case:

Exit status for mutation:

Why the boundary test caught it:
```

## Common mistakes

- Testing ordinary values but omitting values at or adjacent to a boundary.
- Writing expected values by looking at the implementation instead of the contract.
- Changing both the code and the expected result to make a broken test appear green.
- Treating a successful compile as proof that the program meets its contract.
- Reading `$?` after a command other than the test executable.
- Forgetting to restore the deliberate mutation before the transfer task.

## Precision checks

1. State the score contract without using the words “obvious” or “normal.”
2. Point to the exact branch taken by 50 and by 101.
3. Explain why the suite must fail when the threshold operator is mutated.
4. Distinguish the mismatch line from the program's exit status.
5. Rebuild from the saved source and check that the unmutated suite exits with status 0.

## University revision bridge

For the LIG track, separate the formal statement of a predicate from its implementation. Express the passing rule as `50 <= score <= 100`, then explain how the range check and threshold branch implement that predicate. A test table samples the domain; it is not a proof for every possible integer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when, without notes, you can derive test cases from the score contract, build the classifier and test harness from blank, reproduce and diagnose the threshold mutation, and design the 60-point transfer tests. Your evidence must include the commands, observed failure, and both exit statuses. Do not move on because the six original cases looked familiar; rebuild and make the test fail for the intended reason.

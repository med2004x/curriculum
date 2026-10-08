/*
 * STUDENT PROBLEM SET — Lesson 022.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Derive NOT, AND, and OR from NAND and verify their truth tables by hand.
 * 2. Derive a half-adder and full-adder, then connect four stages into a ripple-carry adder.
 * 3. Trace `0110 + 0101` and `1111 + 0001`; distinguish result bits, carry-out, and signed overflow. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 022: Logic Gates and the CPU Datapath
 * Your task: implement the gate using only the allowed primitive operations.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stdio.h>

static int logic_xor(int a, int b) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    for (int a = 0; a <= 1; ++a)
        for (int b = 0; b <= 1; ++b)
            printf("%d xor %d = %d\n", a, b, logic_xor(a, b));
    return 0;
}

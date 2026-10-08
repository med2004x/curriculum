# Lesson 022 — Logic Gates and the CPU Datapath Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Derive NOT, AND, and OR from NAND and verify their truth tables by hand.

2. Derive a half-adder and full-adder, then connect four stages into a ripple-carry adder.

3. Trace `0110 + 0101` and `1111 + 0001`; distinguish result bits, carry-out, and signed overflow.

## Run and check

```powershell
Set-Location lessons/S1/022-logic-gates-and-cpu-datapath/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Tiny ALU model

Model a four-operation ALU from Boolean building blocks or code that explicitly mirrors gates. Verify truth tables and explain the datapath.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

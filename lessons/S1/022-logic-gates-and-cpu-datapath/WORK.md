# Lesson 022 — Problem Set

**Topic:** Logic Gates and the CPU Datapath.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Derive NOT, AND, and OR from NAND and verify their truth tables by hand.

2. Derive a half-adder and full-adder, then connect four stages into a ripple-carry adder.

3. Trace `0110 + 0101` and `1111 + 0001`; distinguish result bits, carry-out, and signed overflow.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/022-logic-gates-and-cpu-datapath/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Tiny ALU model

Model a four-operation ALU from Boolean building blocks or code that explicitly mirrors gates. Verify truth tables and explain the datapath.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

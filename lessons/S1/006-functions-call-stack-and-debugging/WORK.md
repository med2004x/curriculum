# Lesson 006 — Problem Set

**Topic:** Functions, the Call Stack, and Debugging.
**Your files:** `code/factorial.c`, `code/swap.c`, `code/trace_calls.c`, `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. `factorial.c`: implement recursive factorial and verify 5 and 12.
2. `swap.c`: demonstrate the unchanged caller values and explain pass-by-value.
3. `trace_calls.c`: create three nested functions and inspect the call stack in a debugger.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/006-functions-call-stack-and-debugging/code
gcc -std=c17 -Wall -Wextra -Wpedantic factorial.c -o factorial.exe
.\factorial.exe
gcc -std=c17 -Wall -Wextra -Wpedantic swap.c -o swap.exe
.\swap.exe
gcc -std=c17 -Wall -Wextra -Wpedantic trace_calls.c -o trace_calls.exe
.\trace_calls.exe
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Call-trace utility

Implement a recursive or nested function workflow whose output can be reconstructed from the call stack. Add one deliberate bug and document the debugging path.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

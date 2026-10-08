# Lesson 006 — Functions, the Call Stack, and Debugging Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/factorial.c`, `code/swap.c`, `code/trace_calls.c`, `code/exercise.c`.

## Problems

1. `factorial.c`: implement recursive factorial and verify 5 and 12.
2. `swap.c`: demonstrate the unchanged caller values and explain pass-by-value.
3. `trace_calls.c`: create three nested functions and inspect the call stack in a debugger.

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

## Optional challenge

### Call-trace utility

Implement a recursive or nested function workflow whose output can be reconstructed from the call stack. Add one deliberate bug and document the debugging path.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

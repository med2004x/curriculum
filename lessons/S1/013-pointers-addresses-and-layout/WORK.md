# Lesson 013 — Problem Set

**Topic:** Pointers, Addresses, and Memory Layout.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Declare one initialized global and one zero-initialized global; predict their initial values.
2. Add a function with a `static` local and explain why its address remains valid after return.
3. Add one automatic local and one checked `malloc` allocation; print all addresses using `%p` and cast to `(void *)`.
4. Free the heap object exactly once and never read it afterward.
5. Build the separate invalid-lifetime experiment with AddressSanitizer and save the compiler warning and runtime report, if produced.

Use a debugger breakpoint inside the function if available. Inspect the frame and locals, then distinguish debugger observations from guarantees made by the C language.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/013-pointers-addresses-and-layout/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Pointer map

Implement an array transformation using pointers, then print addresses and offsets to explain exactly which object each pointer refers to.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

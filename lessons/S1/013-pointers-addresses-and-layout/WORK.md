# Lesson 013 — Pointers, Addresses, and Memory Layout Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Declare one initialized global and one zero-initialized global; predict their initial values.
2. Add a function with a `static` local and explain why its address remains valid after return.
3. Add one automatic local and one checked `malloc` allocation; print all addresses using `%p` and cast to `(void *)`.
4. Free the heap object exactly once and never read it afterward.
5. Build the separate invalid-lifetime experiment with AddressSanitizer and save the compiler warning and runtime report, if produced.

Use a debugger breakpoint inside the function if available. Inspect the frame and locals, then distinguish debugger observations from guarantees made by the C language.

## Run and check

```powershell
Set-Location lessons/S1/013-pointers-addresses-and-layout/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Pointer map

Implement an array transformation using pointers, then print addresses and offsets to explain exactly which object each pointer refers to.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

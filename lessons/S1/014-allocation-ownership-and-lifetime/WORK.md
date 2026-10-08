# Lesson 014 — Allocation, Ownership, and Lifetime Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/vector.h` and `code/vector.c`. `code/test_vector.c` is the supplied check suite; do not edit it.

## Problems

1. Define `Vector` and its invariant in a header.
2. Implement push from an empty vector, then add growth by doubling capacity.
3. Add bounds-checked `vector_get` and an idempotent cleanup function.
4. Push nine known values, verify every value, and check index `len` fails.
5. Compile with both sanitizers and save the exact command and exit status.

Document that the vector owns its buffer and that callers must not shallow-copy a populated `Vector`; two copies would both appear to own the same pointer and could free it twice.

## Run and check

```powershell
Set-Location lessons/S1/014-allocation-ownership-and-lifetime/code
gcc -std=c17 -Wall -Wextra -Wpedantic -fsanitize=address,undefined -g test_vector.c vector.c -o test_vector.exe
.\test_vector.exe
```

## Optional challenge

### Lifetime-safe container

Build a dynamically sized C container with explicit ownership rules. Add creation, insertion, removal, and destruction paths and verify allocation balance.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

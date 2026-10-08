# Lesson 008 — Structs, Enums, and Modeling Data Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/layout.c`, `code/order.c`, `code/exercise.c`.

## Problems

1. `order.c`: model an order with an enum status and print three states.
2. `layout.c`: measure the verified struct layouts with `sizeof` and `offsetof`.

## Run and check

```powershell
Set-Location lessons/S1/008-structs-models-and-data/code
gcc -std=c17 -Wall -Wextra -Wpedantic layout.c -o layout.exe
.\layout.exe
gcc -std=c17 -Wall -Wextra -Wpedantic order.c -o order.exe
.\order.exe
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Record store

Implement a small in-memory record store using structs and enums with create, update, and summary operations. Define invariants for valid records.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

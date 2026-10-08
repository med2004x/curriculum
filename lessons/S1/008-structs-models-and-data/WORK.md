# Lesson 008 — Problem Set

**Topic:** Structs, Enums, and Modeling Data.
**Your files:** `code/layout.c`, `code/order.c`, `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. `order.c`: model an order with an enum status and print three states.
2. `layout.c`: measure the verified struct layouts with `sizeof` and `offsetof`.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

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

## Stretch problem

### Record store

Implement a small in-memory record store using structs and enums with create, update, and summary operations. Define invariants for valid records.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

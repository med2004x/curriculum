# Lesson 010 — Compilation, Linking, and Toolchains

## Threads

**Primary:** C/Linux

**Supporting:** preprocess/compile/assemble/link; reproducible builds

**Math companion:** sequences and functions review

**Time:** 4 hours. **Prerequisites:** Lesson 009.

## What you will be able to do

1. Name the translation stages from C source to an executable.
2. Use object files and symbol tables to explain a linker failure.
3. Explain why preprocessing, compilation, and linking answer different questions.
4. Build a small multi-file program with explicit, inspectable stages.

## Retrieve first (no notes)

1. What does `argc` count?
2. Why is root unsuitable for a permission-denied demonstration?
3. What errno was verified for a missing file?
4. What changes when stderr is redirected?

## 1. Preprocessing expands source directives

Preprocessing happens before ordinary compilation. The verified command is:

```sh
gcc -E main.c -o main.i
```

For the verified 5-line `main.c` that included `<stdio.h>` and `"util.h"`, `main.i` contained 821 lines.

Self-check: is `main.i` source text, assembly, or an executable?

## 2. Compilation produces assembly

The compiler translates the preprocessed program into target assembly. The verified command is:

```sh
gcc -S main.i -o main.s
```

The verified assembly file was 51 lines and contained `call add@PLT` and `call printf@PLT`.

Self-check: why can assembly mention a function whose body is not in this source file?

## 3. Object files hold incomplete program pieces

Use:

```sh
gcc -c main.c -o main.o
gcc -c util.c -o util.o
```

The two objects can be built separately. Final symbol resolution happens later.

Self-check: which command creates `main.o` without making the final executable?

## 4. Linking resolves symbols

The verified link command is:

```sh
gcc main.o util.o -o app
```

The linked program printed `5 9`. Linking `main.o` alone failed with `undefined reference to 'add'`, then `collect2: error: ld returned 1 exit status`; no `app` was produced.

Self-check: which stage owns an `undefined reference` failure?

## 5. `nm` shows symbol state

The verified object symbols were `main.o -> U add, T main, U printf` and `util.o -> T add`. Here `U` means undefined in that object and `T` means defined in text.

Worked shell example:

```sh
nm main.o
nm util.o
```

Self-check: which object supplies the definition of `add`?

## 6. Macro grouping is a preprocessing problem

The verified macro trap is:

```c
#include <stdio.h>

#define BAD_SQ(x) x*x
#define SQ(x) ((x)*(x))

int main(void) {
    printf("%d %d\n", BAD_SQ(1 + 2), SQ(1 + 2));
    return 0;
}
```

The verified output is `5 9`.

Self-check: expand `BAD_SQ(1+2)` by hand.

## 7. Diagnostics belong to different stages

The verified missing prototype produced `warning: implicit declaration of function 'add' [-Wimplicit-function-declaration]`. A misspelled function produced a warning plus a linker failure. UBSan on signed `INT_MAX + 1` produced `signed integer overflow: 2147483647 + 1 cannot be represented in type 'int'`.

Worked command:

```sh
gcc -Wall -Wextra -std=c11 -fsanitize=address,undefined -g overflow.c -o overflow
```

Self-check: classify warning, linker failure, and runtime sanitizer report separately.

## Predict, then observe

1. Predict which command creates `main.i`.
2. Predict which command creates `main.s`.
3. Predict which command creates an object file.
4. Predict what happens when `util.o` is omitted at link time.
5. Predict the `nm` state of `add` in `main.o`.
6. Predict the two macro results.
7. Predict which stage reports an unresolved symbol.
8. Predict which tool reports signed overflow when enabled.

Run each prediction before checking the result.

## Build from blank

1. `main.c`, `util.c`, `util.h`: define `add`, declare it in the header, and build the executable in explicit stages.
2. `macro.c`: verify `BAD_SQ(1+2) = 5` and `SQ(1+2) = 9`.
3. `symbols.md`: record `nm` results and explain every relevant `U` and `T` symbol.

## Failure drill

1. Link `main.o` without `util.o`; record the verified `undefined reference to 'add'`.
2. Misspell `add`; record the warning and link failure as separate diagnostics.
3. Remove the prototype; record the verified implicit-declaration warning.
4. Run the signed-overflow example under UBSan and record its runtime report.

## Transfer

The constraint changes: use a four-stage build script that leaves the preprocessed file, assembly, and object files visible instead of using a one-shot compile-and-link command.

## Retrieval questions

1. What does `gcc -E` produce?
2. What does `gcc -S` produce?
3. What does `U` mean in the verified `nm` output?
4. Why does linking `main.o` alone fail?
5. What are the two verified macro results?
6. What report does UBSan give for signed `INT_MAX + 1`?

## Evidence to record

Create `lesson010/EVIDENCE.md` with exactly:

```text
Lesson 010 evidence
===================

Stages:
gcc -E main.c -o main.i
gcc -S main.i -o main.s
gcc -c main.c -o main.o
gcc main.o util.o -o app

Verified:
main.i = 821 lines
main.s = 51 lines
main.s contains call add@PLT
main.s contains call printf@PLT

Link failure:
undefined reference to 'add'
collect2: error: ld returned 1 exit status
no app produced

nm:
main.o -> U add, T main, U printf
util.o -> T add

Macros:
BAD_SQ(1+2) = 5
SQ(1+2) = 9
```

## Common mistakes

- Treating compiler and linker errors as the same stage.
- Assuming a successful object build means symbols are resolved.
- Ignoring `nm` during a linker failure.
- Writing function-like macros without parentheses.
- Calling every diagnostic a compiler error.
- Treating signed overflow as a portable result.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can build a multi-file C program in explicit stages, use `nm` to explain symbol ownership, and classify warnings, linker failures, and sanitizer reports correctly. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

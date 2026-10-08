# Lesson 010 — Compilation, Linking, and Toolchains

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Preprocessing expands source directives

Preprocessing happens before ordinary compilation. The verified command is:

```sh
gcc -E main.c -o main.i
```

For the verified 5-line `main.c` that included `<stdio.h>` and `"util.h"`, `main.i` contained 821 lines.

## 2. Compilation produces assembly

The compiler translates the preprocessed program into target assembly. The verified command is:

```sh
gcc -S main.i -o main.s
```

The verified assembly file was 51 lines and contained `call add@PLT` and `call printf@PLT`.

## 3. Object files hold incomplete program pieces

Use:

```sh
gcc -c main.c -o main.o
gcc -c util.c -o util.o
```

The two objects can be built separately. Final symbol resolution happens later.

## 4. Linking resolves symbols

The verified link command is:

```sh
gcc main.o util.o -o app
```

The linked program printed `5 9`. Linking `main.o` alone failed with `undefined reference to 'add'`, then `collect2: error: ld returned 1 exit status`; no `app` was produced.

## 5. `nm` shows symbol state

The verified object symbols were `main.o -> U add, T main, U printf` and `util.o -> T add`. Here `U` means undefined in that object and `T` means defined in text.

Worked shell example:

```sh
nm main.o
nm util.o
```

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

## 7. Diagnostics belong to different stages

The verified missing prototype produced `warning: implicit declaration of function 'add' [-Wimplicit-function-declaration]`. A misspelled function produced a warning plus a linker failure. UBSan on signed `INT_MAX + 1` produced `signed integer overflow: 2147483647 + 1 cannot be represented in type 'int'`.

Worked command:

```sh
gcc -Wall -Wextra -std=c11 -fsanitize=address,undefined -g overflow.c -o overflow
```

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

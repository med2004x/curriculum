# Lesson 013 — Pointers, Addresses, and Memory Layout

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. A pointer is a value, not an object

For `int value = 31; int *pointer = &value;`, `pointer` stores an address and `*pointer` designates the `int` at that address. `&value` produces a pointer to `value`; `*pointer` accesses the pointed-to object. Changing `*pointer` changes `value`.

Pointer equality asks whether two pointers contain the same address. `pointer_a == pointer_b` is not the same question as `*pointer_a == *pointer_b`, which compares the pointed-to values. Never dereference a pointer until you know it points to a live object of the right type.

## 2. Lifetime and the common memory-region model

C specifies object lifetimes and storage durations; it does not require a literal stack or heap layout. On a typical Linux process, automatic locals are implemented using a stack, dynamically allocated objects using a heap, and static-storage objects in program data segments.

- A global object and a `static` local have static storage duration: they exist for the whole program execution.
- A block-scope non-`static` local has automatic storage duration: its lifetime ends when execution leaves its block.
- An object returned by `malloc` exists until it is released with `free`.
- Instructions occupy a code/text mapping, usually protected from writes by the operating system.

For a typical ELF executable, zero-initialized globals are represented in `.bss`; initialized globals are represented in `.data`. These names and address layouts are platform details, not portable C promises. ASLR can change addresses between process runs.

## 3. Observe addresses only while objects are alive

Save this as `layout.c`. It prints the addresses of one initialized global, one zero-initialized global, a `static` local, an automatic local, and a heap object. The heap allocation is checked and freed before exit.

```c
#include <stdio.h>
#include <stdlib.h>

int initialized_global = 17;
int zero_global;

static int *static_object(void) {
    static int value = 23;
    return &value;
}

int main(void) {
    int automatic = 31;
    int *heap = malloc(sizeof *heap);
    if (heap == NULL) {
        perror("malloc");
        return 1;
    }
    *heap = 47;
    printf("global=%p value=%d\n", (void *)&initialized_global, initialized_global);
    printf("zero=%p value=%d\n", (void *)&zero_global, zero_global);
    printf("static=%p value=%d\n", (void *)static_object(), *static_object());
    printf("automatic=%p value=%d\n", (void *)&automatic, automatic);
    printf("heap=%p value=%d\n", (void *)heap, *heap);
    free(heap);
    return 0;
}
```

Compile with warnings and run twice:

```sh
gcc -std=c17 -Wall -Wextra -Wpedantic -g layout.c -o layout
./layout
./layout
```

The values should remain 17, 0, 23, 31, and 47. Addresses are valid to observe during each run, but exact addresses can change with ASLR. Do not order unrelated C pointers with `<` to prove a universal region order; that comparison is not defined for arbitrary objects.

## 4. The lifetime boundary

This function returns an address to an automatic object:

```c
int *bad_pointer(void) {
    int local = 7;
    return &local;
}
```

When the function returns, `local`'s lifetime ends. The returned pointer cannot be used to read that object. Dereferencing it has undefined behavior: C does not promise a crash, a changed number, a diagnostic, or any particular output. A plausible printed value is not evidence that the access is valid.

This is different from returning the address of a `static` local, whose lifetime lasts until program termination, or returning allocated storage that remains live until its owner calls `free`. Each design has a different lifetime contract.

## 5. Catch the invalid access

Put `bad_pointer` in a separate test program and dereference its result only in a sanitizer build. Do not use the unsanitized output as a result to memorize; it is undefined behavior and may vary by compiler and optimization level.

```sh
gcc -std=c17 -Wall -Wextra -Wpedantic -fsanitize=address -fno-omit-frame-pointer -g bad.c -o bad
ASAN_OPTIONS=detect_stack_use_after_return=1 ./bad
```

The compiler may warn that the function returns the address of a local. If AddressSanitizer detects the invalid read, record the exact report and source line. Sanitizer support and report wording vary by platform; absence of a report does not make the C access valid. The language rule comes from object lifetime, not from the tool's output.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

# Lesson 013 — Pointers, Addresses, and Memory Layout

## Threads

**Primary:** C

**Supporting:** address-of; dereference; object lifetime; static, automatic, and allocated storage

**Math companion:** functions and graphs

**Time:** 3–4 hours. **Prerequisites:** Lesson 012.

## What you will be able to do

1. Distinguish a pointer value from the object it points to.
2. Trace `&` and `*` and identify when a pointer is valid to dereference.
3. Relate C object lifetime to the common static, stack, and heap model.
4. Explain why returning a pointer to an automatic local creates undefined behavior.
5. Use address output and a sanitizer as evidence without treating an observed address as a language guarantee.

## Retrieve first (no notes)

1. In Lesson 012, what fields did one log record contain, and which separator split them?
2. Why was a writable line buffer passed to `parse_record`?
3. What should the CLI send to stderr when a record is malformed?
4. Which command linked `main.c` with `log.c` into one executable?

Answer from memory. Then connect the lifetime of the `line` buffer in Lesson 012 to the pointer rules in this lesson: a pointer is useful only while the pointed-to object exists.

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

## Predict, then observe

1. Predict which objects in `layout.c` remain alive after `main` exits.
2. Predict which printed values can change between runs and which should not.
3. Predict what `*pointer = 44` does when `pointer` points to the live automatic `automatic` variable.
4. Predict whether `bad_pointer`'s returned address is valid after the call.
5. Predict what evidence a sanitizer may report, and what its absence cannot prove.
6. Predict why `free(heap)` ends the allocated object's lifetime.

Record addresses as observations from one run, not as portable constants. Label every dereference as valid or invalid before running it.

## Build from blank

1. Declare one initialized global and one zero-initialized global; predict their initial values.
2. Add a function with a `static` local and explain why its address remains valid after return.
3. Add one automatic local and one checked `malloc` allocation; print all addresses using `%p` and cast to `(void *)`.
4. Free the heap object exactly once and never read it afterward.
5. Build the separate invalid-lifetime experiment with AddressSanitizer and save the compiler warning and runtime report, if produced.

Use a debugger breakpoint inside the function if available. Inspect the frame and locals, then distinguish debugger observations from guarantees made by the C language.

## Failure drill

1. Return a pointer to an automatic local and explain the exact lifetime boundary.
2. Read through that pointer only in the sanitizer experiment; classify any output as undefined behavior, not as a stable result.
3. Remove the `malloc` null check, then explain which failure path became unsafe.
4. Read through the heap pointer after `free`; identify the use-after-free before running a sanitizer.
5. Compare `p == q` with `*p == *q` using two live integers that have equal values but distinct addresses.

Restore the safe `layout.c` program and verify that the heap object is freed once and no dead object is dereferenced.

## Transfer

In Go, return a pointer to a local variable and run `go build -gcflags="-m"`. Find the escape-analysis message. Explain why Go can move the object so it outlives the function, while C requires the programmer to choose storage whose lifetime is long enough.

## Retrieval questions

1. What does `&value` produce, and what does `*pointer` access?
2. When does the lifetime of a non-static local end?
3. Why is dereferencing a pointer to a returned local undefined behavior?
4. Which storage durations keep an object alive beyond its creating function?
5. What can an ASan report establish, and what can silence not establish?
6. Why should arbitrary object pointers not be ordered to infer a portable memory layout?

## Evidence to record

Create `lesson013/EVIDENCE.md` with:

```text
Prediction of storage durations:

Build command:

Values and address observations:

One address change between runs:

Sanitizer command and report, if available:

Why a plausible stale value is not valid evidence:
```

## Common mistakes

- Confusing a pointer's numeric address with the lifetime of the object there.
- Assuming the stack and heap arrangement is specified by C.
- Treating one successful stale-pointer read as defined behavior.
- Forgetting that `free` ends the allocated object's lifetime.
- Comparing unrelated pointers to infer a portable address order.
- Assuming a sanitizer must diagnose every undefined behavior on every platform.

## Precision checks

1. State the lifetime rule for each object in `layout.c`.
2. Explain why `%p` receives a `(void *)` argument.
3. Separate a platform's observed address layout from a C guarantee.
4. Identify the exact operation that becomes invalid after `bad_pointer` returns.
5. Explain why no sanitizer report cannot prove the access was valid.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can distinguish pointer values from pointed-to objects, trace the five live objects in the address program, explain the exact lifetime violation in `bad_pointer`, and compare the C failure with Go escape analysis. Your evidence must separate portable language rules from one platform's address and sanitizer observations.

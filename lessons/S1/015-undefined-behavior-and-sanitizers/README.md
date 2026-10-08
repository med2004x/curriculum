# Lesson 015 — Undefined Behavior, Sanitizers, and Defensive Debugging

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

C is a language with a precise contract, but the contract is only guaranteed for well-defined programs. Once a program steps outside those rules, the standard gives no requirement for the result. That is the heart of undefined behavior: the program is no longer in a defined state, and the compiler is free to assume that the undefined case never occurs.

This matters because compilers use this assumption to optimize. A common example is signed overflow:

```c
#include <limits.h>
#include <stdio.h>

int main(void) {
    int x = INT_MAX;
    int y = x + 1; // UB: signed overflow
    printf("%d\n", y);
    return 0;
}
```

The expression `x + 1` may be folded away or transformed because the compiler may assume the earlier arithmetic was valid. The program is not “wrong in a portable way”; it is out of the language’s contract.

## 2. Mechanism and vocabulary

The important terms are:

- undefined behavior (UB): behavior that is not specified by the C standard; anything can happen
- implementation-defined behavior: the standard allows multiple valid outcomes, and the implementation must document which one it chooses
- sanitizer: compiler- or runtime-inserted checks that detect dangerous programming patterns
- optimization assumption: compilers often assume UB cannot occur and generate code based on that assumption

The classic UB examples are:

- signed integer overflow
- out-of-bounds array access
- null pointer dereference
- use-after-free or dangling pointer access
- reading an uninitialized local variable

A sanitizer does not “fix” the program; it turns silent corruption into a crash, diagnostic, or clear violation message before the program continues into nonsense.

## 3. Worked case: signed overflow and sanitizer output

Write a minimal program that adds one to `INT_MAX` and prints the value. Then compile it once without sanitizers and once with a sanitizer:

```bash
gcc -O2 -Wall -Wextra overflow.c -o overflow
gcc -O2 -fsanitize=undefined -g overflow.c -o overflow_ubsan
```

The unsanitized version may appear to work, or may produce a value that looks plausible. The sanitized version should emit an error pointing to the exact source line and the offending operation. This is the key lesson: UB is not just “bad code”; it is a state that makes optimizer assumptions invalid.

## 4. Verification, failure, and evidence

The evidence you record should include:

- the source line that triggers UB
- the compilation flags used
- the observed output or crash
- the sanitizer diagnosis

For example, a good note is:

- `gcc -O2 -Wall -Wextra` produced a plausible but wrong-looking result
- `gcc -O2 -fsanitize=undefined -g` reported signed integer overflow at line 6
- `-fsanitize=address` is useful for invalid memory access, not just arithmetic overflow

A defensive C habit is: always assume an invalid state is possible until proven otherwise. Check all pointer values, array bounds, and integer arithmetic before using an untrusted value.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

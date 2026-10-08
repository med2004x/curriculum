# Lesson 007 — Arrays, Strings, and Representation

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Array size and element count

A declared array contains a fixed number of elements. For `int a[5]`, the verified `sizeof a` is 20, the element count expression is 5, and adjacent elements differ by 4 bytes in the measured layout.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int a[5] = {10, 20, 30, 40, 50};
    printf("%zu\n", sizeof a);
    printf("%zu\n", sizeof a / sizeof a[0]);
    return 0;
}
```

## 2. Bounds are part of the representation

The valid positions in an array of length 5 are 0 through 4. `a[5]` is outside the object.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int a[5] = {10, 20, 30, 40, 50};
    printf("%d\n", a[4]);
    return 0;
}
```

The last valid index is 4.

## 3. C strings use a NUL terminator

For `char s[] = "hello"`, the verified `sizeof s` is 6, `strlen(s)` is 5, and `s[5]` is zero. The extra byte marks the end of the string.

Worked example:

```c
#include <stdio.h>
#include <string.h>

int main(void) {
    char s[] = "hello";
    printf("%zu %zu %d\n", sizeof s, strlen(s), s[5] == '\0');
    return 0;
}
```

## 4. Not every character array is a string

`char t[5] = {'h','e','l','l','o'};` has size 5 and no terminator. Calling `strlen(t)` is undefined because the function searches beyond the intended array for a zero byte.

Worked example:

```c
#include <stdio.h>

int main(void) {
    char t[5] = {'h', 'e', 'l', 'l', 'o'};
    printf("%zu\n", sizeof t);
    return 0;
}
```

## 5. `sizeof` can expose the representation

`sizeof` measures the object named by its operand. In the verified comparison, `sizeof s` is 6 for `char s[] = "hello"`, while `sizeof lit` is 8 for `char *lit = "abc"` on the verified x86-64 environment. Pointer semantics are deferred to Lesson 013.

## 6. Sanitizers expose boundary violations

The verified out-of-bounds access `a[5]` compiled with no warning under `-Wall -Wextra`. With AddressSanitizer enabled, it reported `stack-buffer-overflow ... READ of size 4`.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int a[5] = {10, 20, 30, 40, 50};
    printf("%d\n", a[5]);
    return 0;
}
```

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

# Lesson 008 — Structs, Enums, and Modeling Data

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Structs group related values

Worked example:

```c
#include <stdio.h>

struct User {
    int id;
    int active;
};

int main(void) {
    struct User u = {42, 1};
    printf("%d %d\n", u.id, u.active);
    return 0;
}
```

## 2. Structs can carry invariants

Worked example:

```c
#include <stdio.h>

struct Account { int balance; int open; };

int main(void) {
    struct Account a = {100, 1};
    if (a.open && a.balance >= 0) puts("valid");
    return 0;
}
```

## 3. Struct assignment is value assignment

When two objects have the same struct type, assignment copies the struct value.

Worked example:

```c
#include <stdio.h>

struct Item { int i; };

int main(void) {
    struct Item x = {7};
    struct Item y = x;
    y.i = 99;
    printf("%d %d\n", x.i, y.i);
    return 0;
}
```

The verified values are 7 and 99.

## 4. Layout includes padding

The compiler may insert padding between fields. On the verified x86-64 environment, `struct A {char c; int i; char d;}` has size 12 with offsets 0, 4, 8. Reordering to `struct B {int i; char c; char d;}` gives size 8 with offsets 0, 4, 5.

Worked example:

```c
#include <stdio.h>
#include <stddef.h>

struct A { char c; int i; char d; };

int main(void) {
    printf("%zu %zu %zu %zu\n", sizeof(struct A),
           offsetof(struct A, c), offsetof(struct A, i), offsetof(struct A, d));
    return 0;
}
```

Verified output: `12 0 4 8`.

## 5. Enums name finite states

An enum gives names to integer constants. With `RED`, `GREEN = 5`, `BLUE`, the verified values are 0, 5, and 6.

Worked example:

```c
#include <stdio.h>

enum Color { RED, GREEN = 5, BLUE };

int main(void) {
    printf("%d %d %d\n", RED, GREEN, BLUE);
    return 0;
}
```

## 6. Modeling is explicit state design

A struct groups fields; an enum narrows a category to named values. Together they replace magic numbers with named states.

Worked example:

```c
#include <stdio.h>

enum Status { PENDING, PAID, CANCELLED };
struct Order { int id; enum Status status; };

int main(void) {
    struct Order order = {7, PAID};
    if (order.status == PAID) puts("ready");
    return 0;
}
```

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

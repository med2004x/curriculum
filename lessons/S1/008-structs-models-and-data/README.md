# Lesson 008 — Structs, Enums, and Modeling Data

## Threads

**Primary:** C

**Supporting:** records; invariants; domain modeling

**Math companion:** sets, functions, relations

**Time:** 4 hours. **Prerequisites:** Lesson 007.

## What you will be able to do

1. Define a struct that groups related fields into one domain value.
2. Measure struct size and member offsets with `sizeof` and `offsetof`.
3. Explain how field order changes padding and total size.
4. Use enums and struct assignment to represent finite states.

## Retrieve first (no notes)

1. Why does a C string need a NUL byte?
2. What is the last valid index of `int a[5]`?
3. What is `strlen("hello")`?

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

Self-check: name the type, object, and fields.

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

Self-check: which part of the condition proves the account is open?

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

Self-check: which object changed after `y.i = 99`?

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

Self-check: why does `i` start at 4?

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

Self-check: why is `BLUE` one more than the previous explicit value?

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

Self-check: list all three legal status labels before writing a branch.

## Predict, then observe

1. Predict the values of `u.id` and `u.active`.
2. Predict the result after assigning `y = x` and changing only `y`.
3. Predict `sizeof(struct A)`.
4. Predict the three member offsets in `struct A`.
5. Predict `sizeof(struct B)`.
6. Predict `RED`, `GREEN`, and `BLUE`.

Run each prediction before checking the result.

## Build from blank

1. `order.c`: model an order with an enum status and print three states.
2. `layout.c`: measure the verified struct layouts with `sizeof` and `offsetof`.

## Failure drill

1. Reorder `struct A` and measure the size change.
2. Remove `GREEN = 5`; predict the new enum sequence before compiling.
3. Assign a status value outside your intended domain and state the invariant it violates.

## Transfer

The constraint changes: replace a boolean `active` field with a four-state priority. Choose an enum representation, name every legal state, and write one invariant that rejects an invalid state.

## Retrieval questions

1. What is the verified size of `struct A`?
2. What are the verified offsets in `struct A`?
3. What is the verified size of `struct B`?
4. What is the verified size of `struct C`?
5. What are the verified values of RED, GREEN, and BLUE?
6. What does `offsetof` measure?

## Evidence to record

Create `lesson008/EVIDENCE.md` with exactly:

```text
Lesson 008 evidence
===================
A: size=12, offsets=0,4,8
B: size=8, offsets=0,4,5
C: size=24, _Alignof(double)=8
Enum: RED=0, GREEN=5, BLUE=6
Assignment: x.i=7, y.i=99
```

## Common mistakes

- Treating a struct as only formatting.
- Assuming field order cannot affect size.
- Using magic integers when a finite state can be named.
- Using `%d` for `sizeof` or `offsetof`.
- Assuming one ABI's layout is universal.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can design a small C data model, measure its layout, explain padding, and use enums to express finite states without magic numbers. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

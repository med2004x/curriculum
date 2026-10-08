# C Quick Reference

Use this page only to unblock syntax. It is not a substitute for the lessons.

## Compile

```bash
gcc -std=c17 -Wall -Wextra -Wpedantic -O0 -g -o app main.c
./app
```

For memory diagnostics:

```bash
gcc -std=c17 -Wall -Wextra -Wpedantic -g -fsanitize=address,undefined -o app main.c
./app
```

## Program Shape

```c
#include <stdio.h>

int main(void) {
    printf("hello\n");
    return 0;
}
```

Prefer `int main(void)` when the program accepts no arguments.

## Variables And Types

```c
int count = 3;
double price = 12.5;
char grade = 'A';
const int limit = 100;
```

Initialization and assignment are different operations. A local variable that was not initialized does not acquire a useful value merely by existing.

## Conditions

```c
if (count > 0 && count < limit) {
    printf("valid\n");
} else {
    printf("invalid\n");
}
```

Short-circuiting means the right side of `&&` or `||` may not execute.

## Loops

```c
for (int i = 0; i < count; i++) {
    printf("%d\n", i);
}

while (count > 0) {
    count--;
}
```

Trace the state changes before running unfamiliar loops.

## Functions

```c
int add(int a, int b) {
    return a + b;
}
```

A function prototype can appear before its definition:

```c
int add(int a, int b);
```

## Arrays And Strings

```c
int values[4] = {10, 20, 30, 40};
char name[] = "Med";
```

A C string is an array of characters terminated by `\0`. Its buffer capacity is not the same thing as its current string length.

## Structs

```c
struct Product {
    int id;
    int price_cents;
};

struct Product product = { .id = 1, .price_cents = 1299 };
```

Use a struct when several fields describe one domain object.

## Pointers

```c
int value = 10;
int *p = &value;
*p = 20;
```

```text
value  -> the integer
&value -> the address of that integer
p      -> a pointer containing that address
*p     -> the integer reached through p
```

## Files

```c
FILE *file = fopen("data.txt", "r");
if (file == NULL) {
    return 1;
}

/* use file */

fclose(file);
```

Every successful resource acquisition needs a corresponding release on every control-flow path.

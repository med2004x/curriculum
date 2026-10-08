# Lesson 01: Program Structure and Values

## Objective

Understand the smallest executable Go program, how `main` is reached, how variables are declared, how Go infers types, and what zero values mean.

## Prerequisites

None.

## Learn

Start with one complete program:

```go
package main

import "fmt"

func main() {
    var name string
    count := 3
    fmt.Println(name, count)
}
```

Reason from the declarations before memorizing syntax:

```text
package main -> executable program
import      -> dependency on fmt
main()      -> entry point
name        -> explicit string with zero value ""
count       -> inferred int with value 3
```

Go has predictable zero values. A newly declared `int` is `0`, `bool` is `false`, and `string` is `""`. This lets data have a defined initial state without pretending initialization did not happen.

## Attempt

Create a program that stores a product name, price in cents, stock count, and published flag. Print all four values.

Then change only the declarations and predict the output before running it.

## Guided change

Replace one explicit declaration with `:=`. Explain what information Go now gets from the initializer.

## Verify

Predict the output on paper. Then run it. Add one uninitialized variable of each basic type and verify its zero value.

## Failure drill

Use a variable before declaring it. Then assign a string to an integer variable. Read the compiler error and classify each failure: name resolution or type checking.

## Transfer

Write the same program for a bank account and identify which values are state and which are just labels.

## Done when

You can write a small Go program from an empty file, explain every declaration, and predict zero values without running the code.

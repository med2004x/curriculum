# Lesson 02: Conditions, Loops, and Boolean Reasoning

## Objective

Use `if`, comparison operators, boolean operators, and `for` loops to express decisions and repeated work.

## Prerequisites

Lesson 01.

## Learn

A condition is a boolean expression. A loop is a rule for repeating work. The important skill is translating a verbal rule into a sequence of state changes.

Example:

```go
for i := 1; i <= 5; i++ {
    if i%2 == 0 {
        fmt.Println(i, "even")
    }
}
```

Trace it as a table before execution:

```text
i | i <= 5 | i%2==0 | action
1 | true   | false  | none
2 | true   | true   | print
...
```

## Attempt

Write a program that scans integers from 1 to 20, prints multiples of 3, and counts how many were found.

## Guided change

Change the condition to multiples of both 3 and 5. Predict the count first.

## Verify

Produce a trace for at least the first 8 loop iterations.

## Failure drill

Create a loop whose condition never becomes false. Stop it deliberately and explain which state variable failed to make progress.

## Transfer

Apply the same reasoning to scanning product inventory and counting out-of-stock items.

## Done when

You can trace a loop, explain every boolean condition, and identify why a loop terminates.

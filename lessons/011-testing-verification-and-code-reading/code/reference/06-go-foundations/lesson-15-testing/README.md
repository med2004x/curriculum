# Lesson 15: Testing as Evidence

## Objective

Write tests that describe behavior, use failures as evidence, and test boundary cases instead of only the happy path.

## Prerequisites

Lessons 01–14.

## Learn

A test has four pieces:

```text
input -> behavior -> expected result -> comparison
```

Start with a small pure function because it keeps the source of failure visible. Use table-driven tests when the same rule needs several cases.

## Worked example

Open the reference example in `../../supplementary/go-worked-examples/lesson-03-testing` only after you have written one test yourself.

## Attempt

Write tests for a shipping-cost function with:

- normal input
- zero
- minimum boundary
- invalid input
- one large input

## Guided change

Deliberately change one expected value and read the failure output before fixing it.

## Verify

Run:

```powershell
go test -v ./...
```

Add one test that fails first, then restore the implementation.

## Failure drill

Write a test that checks the implementation's internal variable instead of observable behavior. Explain why that test is brittle.

## Transfer

Test one of your earlier aggregation functions with empty, boundary, and ordinary inputs.

## Done when

You can write a useful failing test from a behavior description without copying a test template.

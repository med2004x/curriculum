# Lesson 11: Aggregation with Structs and Slices

## Objective

Combine structs, slices, loops, and return values to calculate derived information. Methods are not used yet.

## Prerequisites

Lessons 01–10.

## Learn

A collection problem usually has two layers:

```text
data model -> Student{name, scores}
algorithm  -> scan scores and accumulate a total
result     -> total / count
```

Keep those layers visible. A method would hide one relationship that you have not learned yet.

## Attempt

Create at least three students. Each student has a name and `[]int` scores. Write:

```go
func average(scores []int) float64
func passed(student Student) bool
```

`passed` should use an average threshold of 50.

## Guided change

Handle an empty score slice explicitly. Decide what result represents “no scores” rather than relying on integer division by accident.

## Verify

Test a student with all scores above 50, all below 50, one score exactly 50, and no scores.

## Failure drill

Remove the empty-input check and observe the behavior. Then change integer division to floating-point division and explain why the conversion is needed.

## Transfer

Calculate the average price of products and determine whether the catalog exceeds a chosen stock threshold.

## Done when

You can explain the accumulator, the empty-input rule, and the difference between integer and floating-point division.

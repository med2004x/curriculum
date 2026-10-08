# Lesson 24: Generics After Concrete Repetition

## Objective

Understand type parameters and constraints, and recognize when generics remove real duplication versus when they only make simple code harder to read.

## Prerequisites

Lessons 15–23.

## Learn

Start with two concrete functions that implement the same algorithm for different types. Only then generalize.

The key question is not “can this be generic?” It is “is the algorithm genuinely identical, and does the generic version preserve clarity?”

A constraint describes which operations a type parameter must support. For a set keyed by values, `comparable` is enough.

## Worked example

The reference implementation is in `../../supplementary/go-worked-examples/lesson-06-generics`.

Read it only after writing a concrete `contains` function for at least one type.

## Attempt

Implement `contains` for strings and integers first. Compare the duplication. Then write one generic version with the smallest valid constraint.

## Verify

Compile both versions. Use the compiler to test a type that violates your constraint.

## Failure drill

Make the constraint broader than the algorithm requires. Explain what extra operations you now appear to promise.

## Transfer

Generalize a second small algorithm only if your concrete versions are structurally identical. Otherwise keep them concrete and explain why.

## Done when

You can justify a type parameter from concrete repetition rather than reaching for generics because they exist.

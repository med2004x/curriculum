# Lesson 09: Structs and State

## Objective

Group related values into a struct and write plain functions that operate on that data. Methods are deliberately postponed to later lessons.

## Prerequisites

Lessons 01–08.

## Learn

Start with separate values:

```go
owner := "Sara"
balance := 70
```

When those values belong to one domain object, model the relationship:

```go
type Account struct {
    Owner   string
    Balance int
}
```

A struct gives related data one type. Functions can operate on it without introducing receiver semantics yet:

```go
func deposit(account Account, amount int) Account {
    if amount <= 0 {
        return account
    }
    account.Balance += amount
    return account
}
```

The function returns a new state value. That gives you a clean place to learn mutation, copies, and explicit state transitions before methods complicate the picture.

## Attempt

Build a bank account with:

- owner
- balance
- deposit function
- withdraw function
- summary function

Deposit 100, withdraw 30, and print the resulting account.

## Guided change

Add a rule that a withdrawal larger than the balance leaves the state unchanged. Trace one valid and one invalid withdrawal.

## Verify

Write down the balance after each operation. The printed result is not enough: your trace must show the transition.

## Failure drill

Pass an `Account` to a function, change the local copy, and deliberately forget to return it. Observe that the caller's value does not change. Explain why.

## Transfer

Model an inventory item with `stock` and implement `addStock` and `removeStock` as plain functions.

## Done when

You can explain why a struct exists, what its fields represent, and how a plain function can compute an updated state without receiver syntax.

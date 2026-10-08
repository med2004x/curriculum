# Lesson 033 — Go Slices, Maps, Structs, and Data Structures

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Choose data representations that fit the operation: slices for ordered sequences, maps for key lookup and aggregation, and structs for records with named fields.

**The contract.** A slice value describes a view into an underlying array and has a length and capacity. Copying a slice copies its descriptor, not its elements; two slices can therefore share storage. A map stores key/value associations with unspecified iteration order. A struct groups fields; assigning it copies the field values, and reference-like fields such as slices still refer to shared underlying data.

**Worked example.** Define `type Account struct { Owner string; Balance int }`, keep an ordered slice of accounts, and build a `map[string]int` of balances by owner. Explain why the map's printed iteration order should not be part of a test expectation.

## 2. Mechanism and vocabulary

**Mechanism.** `len(s)` is the number of accessible elements; `cap(s)` is the number of elements available from the slice start to the end of its backing array. Reslicing can expose additional elements up to capacity. `append` may reuse the backing array when capacity allows, or allocate a new one when it does not.

**Second point.** Maps are reference-like values: assigning one map variable to another does not duplicate its entries. A nil map can be read but cannot be assigned to; initialize it with `make` or a map literal before insertion.

The earlier `deposit(account Account, amount int) Account` example returns an updated struct because the parameter is a value copy. If the struct contains a slice or map, however, the copied field still refers to the shared backing data/map; “passed by value” does not mean deep copy.

**Worked example.** owner

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
func deposit(account Account, amount int) Account {
    if amount <= 0 {
        return account
    }
    account.Balance += amount
    return account
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Record slice `len` and `cap` before and after reslicing and appending. Create an alias and predict whether an in-capacity append is visible through it. Separately aggregate repeated keys and verify results without assuming map iteration order.

**Failure mode.** Pass an `Account` by value, change its scalar balance, and omit the returned value; then repeat with a slice field and mutate an element. Explain why the caller sees different effects. Also attempt insertion into a nil map and identify the runtime failure.

**Completion criterion.** You can explain slice aliasing and append allocation conditions, map initialization and iteration semantics, and the difference between copying a struct and deep-copying its contents.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

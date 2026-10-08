# Lesson 034 — Go Methods, Interfaces, and Abstraction

## Threads
**Primary:** Go/Design

**Supporting:** method sets; interfaces; substitution

**Math companion:** sets/functions as abstraction

**Time:** 3–4 hours. **Prerequisites:** Lesson 033.

**Time:** 3–4 hours. **Prerequisites:** Lesson 033.

## What you will be able to do

1. Explain how Go methods attach behavior to a named receiver type.
2. Distinguish value and pointer receiver method sets, including interface satisfaction.
3. Define a small consumer-side interface and use two concrete implementations through it.
4. Use a compile-time assertion to make an intended interface contract explicit.

## Retrieve first (no notes)

1. What are a slice's length and capacity?
2. How can two slice values alias the same backing array?
3. Why is map iteration order unspecified?
4. Does passing a struct containing a slice deep-copy the slice elements?

## 1. Problem and model

**What problem this solves.** Methods group operations with a receiver type; interfaces let code depend on behavior rather than a concrete implementation. Used carefully, this allows substitution without forcing an inheritance hierarchy.

**The contract.** A method is a function with a receiver. An interface is satisfied implicitly when a type's method set contains the required methods with matching signatures. A function accepting a small interface can work with any implementation that satisfies that interface.

**Worked example.** Define a `Payer` interface with `Pay(int) bool`, then create card and cash payment types that implement that method. Write one checkout function that accepts `Payer`; it should not need to know which concrete type it receives.

Self-check: does the interface describe the operation the caller needs, or is it merely duplicating the methods of one implementation?

## 2. Mechanism and vocabulary

**Mechanism.** A method declared with receiver `T` belongs to the method sets of both `T` and `*T`. A method declared with receiver `*T` belongs to the method set of `*T`, but not `T`. This matters for interface implementation.

**Second point.** For an addressable value `x` of type `T`, Go may permit the call `x.PointerMethod()` by taking `&x` automatically. That call convenience does not add pointer-receiver methods to the method set of `T`; therefore `T` still may not satisfy an interface requiring that method.

Prefer a narrow interface at the point of use. A compile-time assertion such as `var _ Payer = CardPayment{}` documents that the value type satisfies the contract; use `var _ Payer = (*Account)(nil)` when the pointer type is intended to satisfy it.

**Worked example.** Add one boundary case to the # 034 — Go Methods, Interfaces, and Abstraction exercise.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Use this small interface contract and implementation shape:

```go
type Payer interface {
	Pay(amount int) bool
}

type CardPayment struct {
	Approved bool
}

func (card CardPayment) Pay(amount int) bool {
	return card.Approved && amount > 0
}

var _ Payer = CardPayment{}
```

The value receiver means both `CardPayment` and `*CardPayment` have this method in their method sets. Change the receiver to `*CardPayment` and reason about which of those types still satisfies `Payer`.

Worked example check: identify the concrete value passed to the interface and the exact method signature required for satisfaction.

Self-check: does automatic address-taking for a method call change the method set used by interface satisfaction?

## 4. Verification, failure, and evidence

**Verification.** Add a second `Payer` implementation and pass both through the same checkout function. Add compile-time assertions for the intended concrete type or pointer type. Confirm that a signature mismatch fails at compile time.

**Failure mode.** Change `Pay` to a pointer receiver, then assign a non-pointer `CardPayment` to `Payer`. Explain the method-set error and why a direct method call on an addressable variable can still compile.

**Completion criterion.** You can state the method-set rule, select the intended receiver, and explain how implicit interface satisfaction enables substitution.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether `T` or `*T` satisfies an interface when the method has receiver `T`.
2. Predict the same for a method with receiver `*T`.
3. Predict whether calling a pointer-receiver method on an addressable `T` changes its method set.
4. Predict the result of a compile-time interface assertion for a mismatched signature.
5. Predict whether a value or pointer should be passed when a method mutates receiver state.
6. Predict how adding a second implementation changes the checkout function.
7. Predict whether an interface should expose all methods of the concrete type.
8. Record the interface and receiver method sets before checking assignments.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Define a small `Payer` interface and a value-receiver implementation.

2. Add a second implementation and make one caller depend only on `Payer`.

3. Change the receiver to a pointer receiver and predict interface satisfaction for `T` and `*T`.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one design point: make the interface accept a different method, or make a receiver pointer-based because it mutates state. Explain the resulting method set and which implementation must be passed.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is a method receiver?
2. Which method sets include a value-receiver method?
3. Which method sets include a pointer-receiver method?
4. Why does automatic address-taking not change interface satisfaction?
5. How does Go determine whether a type implements an interface?
6. What does a compile-time assertion document?

## Evidence to record

Create `lesson034/EVIDENCE.md` with exactly this structure:

```text
# Lesson 034 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Assuming a method call that compiles guarantees the value type satisfies an interface.
- Confusing `T` and `*T` method sets.
- Defining a large interface that mirrors one concrete implementation.
- Choosing a value receiver when mutation or non-copyable state requires pointer semantics.
- Relying on runtime tests instead of compile-time checks for interface satisfaction.

## Precision checks

1. Write the method set for `T` and `*T` for both receiver forms.
2. Explain implicit interface satisfaction.
3. State the compile-time contract assertion for the intended type.
4. Explain when value and pointer receivers are appropriate.
5. Show how one caller works with two concrete implementations.

Interfaces describe caller needs; method sets determine which concrete values meet that contract.
## University revision bridge

For the LIG track, connect methods to receiver semantics and interfaces to behavioral abstraction, while treating Go's method-set rules as the precise substitution boundary.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can define a useful interface, implement it with multiple concrete types, state the method sets of `T` and `*T`, and predict interface satisfaction for value and pointer receivers.

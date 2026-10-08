# Lesson 034 — Go Methods, Interfaces, and Abstraction

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Methods group operations with a receiver type; interfaces let code depend on behavior rather than a concrete implementation. Used carefully, this allows substitution without forcing an inheritance hierarchy.

**The contract.** A method is a function with a receiver. An interface is satisfied implicitly when a type's method set contains the required methods with matching signatures. A function accepting a small interface can work with any implementation that satisfies that interface.

**Worked example.** Define a `Payer` interface with `Pay(int) bool`, then create card and cash payment types that implement that method. Write one checkout function that accepts `Payer`; it should not need to know which concrete type it receives.

## 2. Mechanism and vocabulary

**Mechanism.** A method declared with receiver `T` belongs to the method sets of both `T` and `*T`. A method declared with receiver `*T` belongs to the method set of `*T`, but not `T`. This matters for interface implementation.

**Second point.** For an addressable value `x` of type `T`, Go may permit the call `x.PointerMethod()` by taking `&x` automatically. That call convenience does not add pointer-receiver methods to the method set of `T`; therefore `T` still may not satisfy an interface requiring that method.

Prefer a narrow interface at the point of use. A compile-time assertion such as `var _ Payer = CardPayment{}` documents that the value type satisfies the contract; use `var _ Payer = (*Account)(nil)` when the pointer type is intended to satisfy it.

**Worked example.** Add one boundary case to the # 034 — Go Methods, Interfaces, and Abstraction exercise.

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

## 4. Verification, failure, and evidence

**Verification.** Add a second `Payer` implementation and pass both through the same checkout function. Add compile-time assertions for the intended concrete type or pointer type. Confirm that a signature mismatch fails at compile time.

**Failure mode.** Change `Pay` to a pointer receiver, then assign a non-pointer `CardPayment` to `Payer`. Explain the method-set error and why a direct method call on an addressable variable can still compile.

**Completion criterion.** You can state the method-set rule, select the intended receiver, and explain how implicit interface satisfaction enables substitution.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

# Lesson 022 — Logic Gates and the CPU Datapath

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Explain how logic gates combine into circuits that compute, store, and control state.

**The contract.** A circuit's output follows its gate-level Boolean function. Combinational circuits depend on current inputs; sequential circuits also depend on stored state. Transistors implement switching behavior, but a transistor alone is not a complete model of a logic circuit.

**Worked example.** NAND outputs 0 only when both inputs are 1. Thus `NOT(A) = NAND(A,A)`, `AND(A,B) = NOT(NAND(A,B))`, and `OR(A,B) = NAND(NOT(A),NOT(B))`. Verify each identity against all four input rows.

## 2. Mechanism and vocabulary

**Mechanism.** A half-adder has `Sum = A XOR B` and `Carry = A AND B`. A full-adder can be written as `P = A XOR B`, `Sum = P XOR Cin`, and `Cout = (A AND B) OR (P AND Cin)`.

**Second point.** A ripple-carry adder connects each bit's `Cout` to the next more significant bit's `Cin`. The final carry-out is information beyond the fixed-width result; it is not synonymous with signed two's-complement overflow.

Sequential circuits add state. A latch can be transparent while enabled and retain its prior state while disabled; a flip-flop samples according to a clock event. Truth tables specify logic values, not propagation delay.

**Worked example.** Build NOT, AND, OR, XOR from NAND only.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
NAND truth table:
A B | Out
0 0 | 1
0 1 | 1
1 0 | 1
1 1 | 0
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** For at least five input pairs, show both operands, each stage's `Cin`, every sum bit, final `Cout`, and the four-bit result. Include an unsigned carry-out case and a signed overflow case, then explain why those conditions differ.

**Failure mode.** Chain the full adders without connecting Cin from one stage to Cout of the previous stage — leave each Cin tied to 0.

**Completion criterion.** You can verify derived gates by truth table, trace all ripple carries, and explain the distinction between combinational output and stored state.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

# Lesson 022 — Logic Gates and the CPU Datapath

## Threads
**Primary:** Architecture/C

**Supporting:** bits as hardware; combinational/sequential logic

**Math companion:** boolean algebra

**Time:** 3–4 hours. **Prerequisites:** Lesson 021.

**Time:** 3–4 hours. **Prerequisites:** Lesson 021.

## What you will be able to do

1. Construct NOT, AND, OR, and XOR using NAND gates.
2. Derive half-adder and full-adder outputs from Boolean expressions and truth tables.
3. Trace a four-bit ripple-carry addition, including every carry bit.
4. Explain the difference between combinational logic and state-holding sequential logic.

## Retrieve first (no notes)

1. What does an ISA specify?
2. Under System V AMD64, where does an integer return value appear?
3. What does a Boolean truth table record?
4. Why does a ripple-carry stage need the preceding stage's carry-out?

## 1. Problem and model

**What problem this solves.** Explain how logic gates combine into circuits that compute, store, and control state.

**The contract.** A circuit's output follows its gate-level Boolean function. Combinational circuits depend on current inputs; sequential circuits also depend on stored state. Transistors implement switching behavior, but a transistor alone is not a complete model of a logic circuit.

**Worked example.** NAND outputs 0 only when both inputs are 1. Thus `NOT(A) = NAND(A,A)`, `AND(A,B) = NOT(NAND(A,B))`, and `OR(A,B) = NAND(NOT(A),NOT(B))`. Verify each identity against all four input rows.

Self-check: a correct Boolean identity must match its target truth table for every possible input combination.

## 2. Mechanism and vocabulary

**Mechanism.** A half-adder has `Sum = A XOR B` and `Carry = A AND B`. A full-adder can be written as `P = A XOR B`, `Sum = P XOR Cin`, and `Cout = (A AND B) OR (P AND Cin)`.

**Second point.** A ripple-carry adder connects each bit's `Cout` to the next more significant bit's `Cin`. The final carry-out is information beyond the fixed-width result; it is not synonymous with signed two's-complement overflow.

Sequential circuits add state. A latch can be transparent while enabled and retain its prior state while disabled; a flip-flop samples according to a clock event. Truth tables specify logic values, not propagation delay.

**Worked example.** Build NOT, AND, OR, XOR from NAND only.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** For at least five input pairs, show both operands, each stage's `Cin`, every sum bit, final `Cout`, and the four-bit result. Include an unsigned carry-out case and a signed overflow case, then explain why those conditions differ.

**Failure mode.** Chain the full adders without connecting Cin from one stage to Cout of the previous stage — leave each Cin tied to 0.

**Completion criterion.** You can verify derived gates by truth table, trace all ripple carries, and explain the distinction between combinational output and stored state.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict NAND's output for all four input pairs.
2. Predict `NAND(A,A)` for `A=0` and `A=1`.
3. Predict `Sum` and `Cout` for `A=1, B=1, Cin=1`.
4. Trace `0110 + 0101` from least-significant to most-significant bit.
5. Predict the four-bit result and final carry of `1111 + 0001`.
6. Predict what is lost if each stage's `Cin` is tied to zero.
7. Does `0111 + 0001` produce signed overflow, unsigned carry-out, both, or neither?
8. Record the truth table and carry trace before checking a simulator.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Derive NOT, AND, and OR from NAND and verify their truth tables by hand.

2. Derive a half-adder and full-adder, then connect four stages into a ripple-carry adder.

3. Trace `0110 + 0101` and `1111 + 0001`; distinguish result bits, carry-out, and signed overflow.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Chain the full adders without connecting Cin from one stage to Cout of the previous stage — leave each Cin tied to 0.

2. Compute `0110 + 0101` again with this broken wiring.

3. Explain exactly which additions now produce a wrong 4-bit result and why, and connect this to what "ripple carry" actually means — the propagation delay per stage is the source of the ripple-carry adder's speed limit, which is why real CPUs use carry-lookahead adders instead.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint: widen the adder from four bits to eight, or change one input pair. State how many full-adder stages are needed and trace the new carry path.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is NAND's Boolean expression and truth table?
2. How can NAND construct NOT, AND, and OR?
3. Write the `Sum` and `Cout` equations for a full-adder.
4. What connects neighboring stages in a ripple-carry adder?
5. How does unsigned carry-out differ from signed overflow?
6. How does a latch differ from combinational logic?

## Evidence to record

Create `lesson022/EVIDENCE.md` with exactly this structure:

```text
# Lesson 022 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating carry-out as synonymous with signed overflow.
- Forgetting to connect one stage's carry-out to the next stage's carry-in.
- Checking only one row of a derived gate's truth table.
- Describing stored state as a purely combinational output.
- Assuming a truth table establishes timing behavior.

## Precision checks

1. Verify NAND and its derived gates against their truth tables.
2. State both full-adder output equations.
3. Trace every carry in a four-bit addition.
4. Give one example of unsigned carry-out and one of signed overflow.
5. Explain which component stores state and what controls its update.

Truth tables establish the Boolean contract; the datapath exercise shows how small circuits compose into arithmetic behavior.
## University revision bridge

For the LIG track, connect Boolean algebra identities to circuit construction and distinguish combinational logic from sequential state.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can derive basic gates from NAND, compute full-adder outputs, trace a four-bit ripple-carry addition, distinguish carry-out from signed overflow, and explain how a latch retains state.

# Lesson 021 — Assembly and the Compiler View

## Threads
**Primary:** C/Architecture

**Supporting:** read compiler output; registers; calling convention intuition

**Math companion:** algebra review

**Time:** 3–4 hours. **Prerequisites:** Lesson 020.

**Time:** 3–4 hours. **Prerequisites:** Lesson 020.

## What you will be able to do

1. Explain that an ISA specifies the software-visible instruction contract, not a CPU's internal implementation.
2. Trace argument, temporary, and return values in simple x86-64 assembly.
3. Explain the System V AMD64 integer argument and return-value registers.
4. Compare compiler output at different optimization levels without treating assembly text as portable.

## Retrieve first (no notes)

1. What is the difference between an ISA and a CPU microarchitecture?
2. What does a calling convention define?
3. Under System V AMD64, where do the first two integer arguments arrive?
4. Which register carries an integer return value under that ABI?

## 1. Problem and model

**What problem this solves.** Read and write basic x86-64 (or ARM64) assembly, and explain what an ISA is: the contract between hardware and software.

**The contract.** The ISA is the interface, not the implementation. Assembly output depends on the target architecture, ABI, compiler, options, and surrounding code; use it to inspect one build, not as a portable source-level guarantee.

**Worked example.** For `int add(int a, int b) { return a + b; }`, predict the ABI locations. Under System V AMD64, the first two integer arguments arrive in `edi` and `esi`, and the integer result is returned in `eax` (the corresponding 64-bit registers are `rdi`, `rsi`, and `rax`).

Self-check: separate ABI-defined argument/result locations from compiler-specific decisions such as spills and stack-frame layout.

## 2. Mechanism and vocabulary

**Mechanism.** It defines the instructions, registers, memory model, and calling convention a CPU exposes — different CPUs (an Intel chip, an AMD chip, an Apple M-series chip) implement the same x86-64 or ARM64 ISA completely differently in hardware while running the same compiled binary correctly.

**Second point.** This is the same idea as an API contract in software: the caller doesn't need to know the implementation, only the interface.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** int add(int a, int b) {

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```c
int add(int a, int b) {
       return a + b;
   }
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Produce a hand-traced table for the loop function in the assigned exercise: instruction, relevant register values, and branch decision. Treat exact register allocation as compiler-specific; verify the source-level result separately and explain which ABI facts apply to your target.

**Failure mode.** Take the `-O0` assembly from step 1, manually flip the `add` instruction's operand order or change the immediate value, reassemble it by hand-editing (or note the change if you can't reassemble), and predict the wrong output before checking.

**Completion criterion.** You can trace a small function's data flow, identify the ABI contract for arguments and return values, and explain `call` and `ret` without claiming every compiler emits the same prologue.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the argument and result registers for `add` under System V AMD64.
2. Predict one likely difference between `-O0` and `-O2` output and what must remain semantically equivalent.
3. For `sum_to_n(3)`, predict the accumulator after each iteration.
4. Predict the result when the loop bound is zero.
5. Explain what `call` stores and what `ret` consumes.
6. Predict whether an optimized leaf function must have a stack frame.
7. Identify which claim is ABI-defined and which is compiler-output-specific.
8. Record the trace before comparing it with generated assembly.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Write the two-argument `add` function and annotate its ABI-defined inputs and output.

2. Add a `sum_to_n` loop and trace its control flow and accumulator for `n = 3`.

3. Compare output at `-O0` and `-O2`; explain a transformation without claiming its textual form is guaranteed.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take the `-O0` assembly from step 1, manually flip the `add` instruction's operand order or change the immediate value, reassemble it by hand-editing (or note the change if you can't reassemble), and predict the wrong output before checking.

2. This confirms you're reading the semantics of the instruction, not pattern-matching the syntax.

3. Take the two-file program from step 4 and change `helper.h`'s function signature (e.g. change a parameter type) without recompiling `helper.c`, only recompiling and relinking against the stale `helper.o`.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint: target ARM64 instead of x86-64, or change the optimization level. State which observations should change (instructions, registers, or layout) and which source-level result should remain the same for defined code.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does an ISA define, and what does it leave to microarchitecture?
2. Under System V AMD64, where are the first six integer or pointer arguments passed?
3. Which register carries a scalar integer return value under that ABI?
4. What do `call` and `ret` do with the return address?
5. Why can optimized assembly differ from `-O0` output?
6. Which register claims are not portable across ABIs?

## Evidence to record

Create `lesson021/EVIDENCE.md` with exactly this structure:

```text
# Lesson 021 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating one compiler's assembly spelling as part of the C language contract.
- Confusing the ISA with a processor's internal pipeline or execution units.
- Applying the System V AMD64 convention to a different ABI.
- Tracing a branch without updating the condition or branch decision.
- Assuming optimization preserves source statement order in generated text.

## Precision checks

1. Name the target ISA and ABI for every register claim.
2. Trace argument values through the function to its return register.
3. Explain why the same C source can produce different assembly on two targets.
4. Distinguish ABI guarantees from observed register allocation.
5. Explain one transformation made possible by optimization.

Assembly reading is a model-building exercise: annotate what each instruction changes, then check the whole function against the source-level contract.
## University revision bridge

For the LIG track, connect the ISA abstraction to compiler and architecture topics. Keep the selected ABI and target explicit in every register-level explanation.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can trace the `add` and loop examples, identify System V AMD64 argument and return registers, explain the return-address role of `call`/`ret`, and distinguish ABI facts from compiler-specific output.

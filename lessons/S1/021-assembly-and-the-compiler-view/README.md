# Lesson 021 — Assembly and the Compiler View

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Read and write basic x86-64 (or ARM64) assembly, and explain what an ISA is: the contract between hardware and software.

**The contract.** The ISA is the interface, not the implementation. Assembly output depends on the target architecture, ABI, compiler, options, and surrounding code; use it to inspect one build, not as a portable source-level guarantee.

**Worked example.** For `int add(int a, int b) { return a + b; }`, predict the ABI locations. Under System V AMD64, the first two integer arguments arrive in `edi` and `esi`, and the integer result is returned in `eax` (the corresponding 64-bit registers are `rdi`, `rsi`, and `rax`).

## 2. Mechanism and vocabulary

**Mechanism.** It defines the instructions, registers, memory model, and calling convention a CPU exposes — different CPUs (an Intel chip, an AMD chip, an Apple M-series chip) implement the same x86-64 or ARM64 ISA completely differently in hardware while running the same compiled binary correctly.

**Second point.** This is the same idea as an API contract in software: the caller doesn't need to know the implementation, only the interface.

**Worked example.** int add(int a, int b) {

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```c
int add(int a, int b) {
       return a + b;
   }
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Produce a hand-traced table for the loop function in the assigned exercise: instruction, relevant register values, and branch decision. Treat exact register allocation as compiler-specific; verify the source-level result separately and explain which ABI facts apply to your target.

**Failure mode.** Take the `-O0` assembly from step 1, manually flip the `add` instruction's operand order or change the immediate value, reassemble it by hand-editing (or note the change if you can't reassemble), and predict the wrong output before checking.

**Completion criterion.** You can trace a small function's data flow, identify the ABI contract for arguments and return values, and explain `call` and `ret` without claiming every compiler emits the same prologue.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

# FSEGS LIG — Systèmes logiques et architecture des ordinateurs

**Semester:** S1  
**Attached engineering lesson:** 022 — Logic Gates and the CPU Datapath

## Scope

Number representation; binary/decimal/hexadecimal conversion; signed integers; Boolean algebra; logic gates; truth tables; canonical forms; combinational circuits; multiplexers/decoders; adders; registers; CPU datapath; instruction execution; memory hierarchy basics.

## Core teaching

Convert between bases by repeated division or positional expansion. Be explicit about the number of bits when representing integers.

Boolean operations use `AND`, `OR`, `NOT`, and derived operations such as `XOR`. Use truth tables to verify equivalence rather than trusting visual similarity.

For signed integers at first-year level, two's complement is the key representation. Negating a fixed-width value means invert bits and add one. Overflow depends on the fixed width and sign pattern.

A half adder combines two bits into sum and carry. A full adder also includes an incoming carry. A CPU datapath combines registers, arithmetic/logic units, control signals, memory access, and instruction sequencing.

## Worked problems

1. Convert `45₁₀` to binary and hexadecimal, then back.
2. Build the truth table for `(A AND B) OR (NOT A AND B)` and simplify it.
3. Represent `-13` in 8-bit two's complement and verify the result by addition.
4. Draw a one-bit full adder truth table and explain the carry equation.

## Exercises

1. Convert ten values among decimal, binary, and hexadecimal.
2. Perform binary addition/subtraction with fixed width and report overflow.
3. Complete a missing-row truth table.
4. Simplify five Boolean expressions using algebraic identities.
5. Prove two Boolean expressions equivalent by truth table.
6. Design a 2-to-1 multiplexer truth table and logic expression.
7. Design a four-bit ripple-carry adder on paper.
8. Trace a simple instruction through fetch, decode, execute, memory, write-back.
9. Explain why cache locality matters without confusing cache with RAM.
10. Given a datapath diagram, identify where operands, results, and control signals travel.

## Lab

Use [`lig_logic_checker.py`](code/lig_logic_checker.py) to check implication tables and fixed-width two's-complement encoding. Perform base conversion and circuit design by hand; tests cover only these small helpers.

## Exam patterns

Conversions, two's-complement arithmetic, Boolean simplification, truth tables, gate-level reasoning, small combinational circuits, instruction-cycle tracing, architecture definitions, cache/locality explanations.

## Mastery gate

Solve fresh conversion, Boolean, fixed-width arithmetic, circuit, and datapath questions without external tools. Explain every bit-level step.

# STUDENT PROBLEM SET — Lesson 021.
# Read the lecture, then solve these tasks before opening any worked examples:
# 1. Write the two-argument `add` function and annotate its ABI-defined inputs and output.
# 2. Add a `sum_to_n` loop and trace its control flow and accumulator for `n = 3`.
# 3. Compare output at `-O0` and `-O2`; explain a transformation without claiming its textual form is guaranteed. Do not open the reference implementation until your own attempt has been saved.
# Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
# PROBLEM STARTER — Lesson 021: Assembly and the Compiler View
# Your task: build with optimisation on/off and inspect the generated assembly.
# Read ../README.md for the concept; follow ../WORK.md for the problem and checks.
.section .text
.global _start
_start:
    mov $60, %rax
    xor %rdi, %rdi
    syscall

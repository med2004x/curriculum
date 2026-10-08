# Code map â€” Lesson 064 â€” Systems Lab: xv6 and Kernel Boundaries

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Kernel-boundary lab

Use xv6 or a small OS model to trace one user-to-kernel path. Document the transition points and validate one change with a repeatable test.

### Acceptance criteria

- Start from a clean workspace or the latest lesson project without copying a completed solution.
- Implement the required behavior end-to-end.
- Include automated tests where the technology allows it, plus one deliberate failure test.
- Record design assumptions and the main invariant(s) in Markdown.
- Demonstrate the project from a reproducible command sequence.
- Finish with a short post-mortem: what failed, why, and what you changed.

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.

## Upstream lab: MIT xv6-RISC-V

This lesson uses the real xv6 repository and MIT lab tests. The tiny local C file is only a prompt; it is not an xv6 implementation or replacement lab.

- Upstream source: https://github.com/mit-pdos/xv6-riscv
- Local reference checkout: `../../../../dart-upstreams/xv6-riscv`
- Exact lab handout and tests: https://pdos.csail.mit.edu/6.S081/2026/labs/syscall.html
- Course lab repo: `git://g.csail.mit.edu/xv6-labs-2026`

Clone the course lab repo, check out the `syscall` branch as instructed, and run `make grade`. Record the release/commit, lab, commands, and results in `EVIDENCE.md`. DART acceptance is the upstream test suite plus an explanation of the kernel invariant.

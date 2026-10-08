# Lesson 059 — Queues, Scheduling, and Throughput

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand the tradeoffs between CPU scheduling policies by implementing and measuring several of them directly, connecting the abstract turnaround/response-time metrics to concrete, felt differences in workload behavior.

**The contract.** **The scheduling problem.** With more ready processes than CPU cores, the OS must decide which runs next and for how long.

**Worked example.** Start with the smallest concrete case described by the source material: Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

## 2. Mechanism and vocabulary

**Mechanism.** For a process, `turnaround = completion - arrival`, `waiting = turnaround - burst` for a single-CPU model without I/O blocking, and `response = first_run - arrival`. Keep these metrics distinct; response measures initial latency, not total completion delay.

**Second point.** Non-preemptive SJF selects the shortest burst among jobs ready at each decision point. It minimizes average waiting time for a known set of jobs that are all available together, under the usual single-CPU assumptions; that optimality does not transfer automatically to arbitrary arrival patterns or unknown burst lengths.

**Worked example.** Simulate Round Robin with a stated quantum `q`. The ready queue rotates unfinished processes after each quantum; response time is the first dispatch time minus arrival, while turnaround uses final completion time.

## 3. Worked case from the reference material

Worked example: Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

## 4. Verification, failure, and evidence

**Verification.** Produce a Gantt-style timeline and per-process completion, turnaround, waiting, and response values for FCFS, non-preemptive SJF, and Round Robin with a stated quantum. Check each process's elapsed and service times against the timeline.

**Failure mode.** To illustrate starvation under non-preemptive SJF, start one short job so the CPU is occupied while a long job waits, then let shorter jobs continue arriving often enough that the ready queue always contains a job shorter than the long one. Do not claim SJF can preempt a long job already running; that is a property of SRTF, the preemptive variant.

**Completion criterion.** You can compute scheduling metrics from a timeline, state SJF's assumptions, distinguish SJF from SRTF, and compare responsiveness, waiting, and fairness without claiming one policy dominates under all workloads.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

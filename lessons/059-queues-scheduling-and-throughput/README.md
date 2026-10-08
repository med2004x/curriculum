# Lesson 059 — Queues, Scheduling, and Throughput

## Threads
**Primary:** OS/Math

**Supporting:** schedulers; queue buildup; utilization

**Math companion:** queueing foundations

**Time:** 3–4 hours. **Prerequisites:** Lesson 058.

**Time:** 3–4 hours. **Prerequisites:** Lesson 058.

## What you will be able to do

1. Simulate FCFS, non-preemptive SJF, and Round Robin from arrival/burst-time inputs.
2. Calculate completion, turnaround, waiting, and response times consistently.
3. Explain the assumptions under which SJF minimizes average waiting/turnaround time.
4. Explain how preemption, time quantum, and arrival patterns affect fairness and response.

## Retrieve first (no notes)

1. What is the difference between runnable and blocked process states?
2. What does a syscall trace reveal?
3. What does a parent do to reap an exited child?
4. Why can one library call produce multiple syscalls?

## 1. Problem and model

**What problem this solves.** Understand the tradeoffs between CPU scheduling policies by implementing and measuring several of them directly, connecting the abstract turnaround/response-time metrics to concrete, felt differences in workload behavior.

**The contract.** **The scheduling problem.** With more ready processes than CPU cores, the OS must decide which runs next and for how long.

**Worked example.** Start with the smallest concrete case described by the source material: Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** For a process, `turnaround = completion - arrival`, `waiting = turnaround - burst` for a single-CPU model without I/O blocking, and `response = first_run - arrival`. Keep these metrics distinct; response measures initial latency, not total completion delay.

**Second point.** Non-preemptive SJF selects the shortest burst among jobs ready at each decision point. It minimizes average waiting time for a known set of jobs that are all available together, under the usual single-CPU assumptions; that optimality does not transfer automatically to arbitrary arrival patterns or unknown burst lengths.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Simulate Round Robin with a stated quantum `q`. The ready queue rotates unfinished processes after each quantum; response time is the first dispatch time minus arrival, while turnaround uses final completion time.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Produce a Gantt-style timeline and per-process completion, turnaround, waiting, and response values for FCFS, non-preemptive SJF, and Round Robin with a stated quantum. Check each process's elapsed and service times against the timeline.

**Failure mode.** To illustrate starvation under non-preemptive SJF, start one short job so the CPU is occupied while a long job waits, then let shorter jobs continue arriving often enough that the ready queue always contains a job shorter than the long one. Do not claim SJF can preempt a long job already running; that is a property of SRTF, the preemptive variant.

**Completion criterion.** You can compute scheduling metrics from a timeline, state SJF's assumptions, distinguish SJF from SRTF, and compare responsiveness, waiting, and fairness without claiming one policy dominates under all workloads.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the first completion and response times under FCFS for the supplied workload.
2. Predict which jobs are ready when non-preemptive SJF makes its next choice.
3. Predict how changing Round Robin quantum affects context-switch frequency and response.
4. Predict the difference between turnaround time and waiting time.
5. Predict how an arrival-time shift changes a scheduling decision.
6. Predict when a long job can starve under SJF.
7. Predict how SRTF differs from non-preemptive SJF.
8. State the assumptions used in any optimality claim.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

2. Extend it with one explicit boundary case: Simulate SJF (non-preemptive — once a process starts, it runs to completion, but among *ready* processes at each decision point, pick the shortest) for the same workload, and compute the same metrics.

3. Rebuild the same behavior with one changed constraint: Compare average turnaround and response times, and explain why SJF's all-jobs-ready optimality statement does not apply blindly to every arrival schedule.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Construct a workload specifically designed to make SJF starve a process: have one long job arrive first, and then a continuous stream of short jobs arriving one after another for a long time, each shorter than the long job's remaining time.

2. Simulate SJF on this workload (extend your step 4 simulator, or reason through it by hand for a representative case) and confirm the long job's turnaround time grows unboundedly worse the longer the stream of short jobs continues, even though the long job arrived first and, under FCFS, would have finished at a fixed, predictable time.

3. Explain why this is not a simulator bug but SJF's genuine, provable weakness — optimal *average* turnaround time can come at the cost of individual-job starvation, which is exactly why real systems that use SJF-like priority ideas (e.g.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 059 — Queues, Scheduling, and Throughput: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Define completion, turnaround, waiting, and response time.
2. Which jobs may non-preemptive SJF choose among?
3. Under what assumption does SJF minimize average waiting time?
4. How does Round Robin use its time quantum?
5. How does SRTF differ from SJF?
6. What arrival pattern can starve a long waiting job under SJF?

## Evidence to record

Create `lesson059/EVIDENCE.md` with exactly this structure:

```text
# Lesson 059 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Confusing response time with turnaround time.
- Letting SJF select a job that has not arrived yet.
- Claiming non-preemptive SJF can interrupt a running job.
- Applying SJF optimality to arbitrary arrivals or unknown bursts without qualification.
- Comparing Round Robin results without stating the quantum.

## Precision checks

1. Show the schedule timeline for each policy.
2. Derive all four per-process timing metrics from the timeline.
3. State which processes are ready at every scheduling decision.
4. Qualify the SJF optimality claim and distinguish SRTF.
5. Explain the fairness/response trade-off of the chosen quantum.

Scheduling comparisons are meaningful only when workload assumptions and metric definitions are held constant.
## University revision bridge

For the LIG track, connect queue order and service policy to response/turnaround calculations, stating the workload assumptions behind each comparison.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 059 — Queues, Scheduling, and Throughput without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've hand-computed and then code-verified turnaround/response times for all three policies on the same workload, you've directly demonstrated SJF's starvation failure mode with a constructed adversarial workload, and you can explain — using your own comparison table — why no single one of these three simple policies dominates the others on both metrics sim. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

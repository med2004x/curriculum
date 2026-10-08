# Lesson 059 — Problem Set

**Topic:** Queues, Scheduling, and Throughput.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

2. Extend it with one explicit boundary case: Simulate SJF (non-preemptive — once a process starts, it runs to completion, but among *ready* processes at each decision point, pick the shortest) for the same workload, and compute the same metrics.

3. Rebuild the same behavior with one changed constraint: Compare average turnaround and response times, and explain why SJF's all-jobs-ready optimality statement does not apply blindly to every arrival schedule.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/059-queues-scheduling-and-throughput/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Queue simulator

Build a queue simulation with configurable arrival/service rates. Report throughput, utilization, and waiting behavior as load increases.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

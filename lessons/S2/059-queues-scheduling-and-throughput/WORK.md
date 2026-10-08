# Lesson 059 — Queues, Scheduling, and Throughput Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

2. Simulate SJF (non-preemptive — once a process starts, it runs to completion, but among *ready* processes at each decision point, pick the shortest) for the same workload, and compute the same metrics.

3. Compare average turnaround and response times, and explain why SJF's all-jobs-ready optimality statement does not apply blindly to every arrival schedule.

## Run and check

```powershell
Set-Location lessons/S2/059-queues-scheduling-and-throughput/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Queue simulator

Build a queue simulation with configurable arrival/service rates. Report throughput, utilization, and waiting behavior as load increases.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

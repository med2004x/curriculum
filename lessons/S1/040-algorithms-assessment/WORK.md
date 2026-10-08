# Lesson 040 — Integrated Assessment: Algorithms Under Constraints Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Define the task record and validation policy; build the prerequisite graph and indegrees.

2. Implement topological ordering and cycle detection; verify every returned edge ordering.

3. Compute earliest finish times in topological order and report the critical path duration.

4. Write a concise report leading with validity and completion time, followed by the algorithm, complexity, and limitations.

## Run and check

```powershell
Set-Location lessons/S1/040-algorithms-assessment/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Timed algorithms assessment

Solve a mixed algorithm assessment from a blank workspace under a fixed time budget. Include complexity analysis and one post-mortem on a failed attempt.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

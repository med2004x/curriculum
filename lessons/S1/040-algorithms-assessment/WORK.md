# Lesson 040 — Problem Set

**Topic:** Integrated Assessment: Algorithms Under Constraints.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define the task record and validation policy; build the prerequisite graph and indegrees.

2. Implement topological ordering and cycle detection; verify every returned edge ordering.

3. Compute earliest finish times in topological order and report the critical path duration.

4. Write a concise report leading with validity and completion time, followed by the algorithm, complexity, and limitations.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/040-algorithms-assessment/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Timed algorithms assessment

Solve a mixed algorithm assessment from a blank workspace under a fixed time budget. Include complexity analysis and one post-mortem on a failed attempt.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

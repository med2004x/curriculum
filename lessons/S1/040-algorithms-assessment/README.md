# Lesson 040 — Integrated Assessment: Algorithms Under Constraints

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build and explain a small algorithmic system under explicit input and output constraints. The assessment combines graph modeling, topological ordering, dynamic programming on a DAG, validation, complexity analysis, and technical communication.

**The contract.** Input tasks with unique IDs, nonnegative durations, and prerequisite IDs. Reject duplicate IDs and unknown prerequisites. Return an order in which every prerequisite precedes its dependent task, or report a cycle. For a valid DAG, calculate each task's earliest finish from the maximum earliest finish among its prerequisites plus its own duration. Explain runtime in terms of `V` tasks and `E` prerequisite edges.

**Worked example.** Tasks: `A` duration 2 with no prerequisite; `B` duration 3 depends on A; `C` duration 4 depends on A; `D` duration 1 depends on B and C. A valid order places A before B/C and both before D. Earliest finishes are `A=2`, `B=5`, `C=6`, `D=7`; total completion time is 7.

## 2. Mechanism and vocabulary

**Mechanism.** Kahn's algorithm tracks each vertex's indegree, queues vertices with indegree zero, emits one, and decreases the indegree of each outgoing neighbor. If fewer than `V` vertices are emitted, the graph contains a directed cycle. The traversal is O(V+E) with adjacency lists.

**Second point.** Earliest-finish dynamic programming follows topological order: for task `v`, `EF(v)=duration(v)+max(EF(p))` over prerequisites `p`; the maximum over no prerequisites is zero. This recurrence is valid because all prerequisite values are available before `v` is processed.

The report should lead with whether the input is valid and the completion time, then state the ordering and complexity. Do not call the longest path the topological order; those are different outputs.

**Worked example.** State the result first: “The four-task plan completes in 7 time units; the critical dependency chain is A → C → D.” Then give the validation rules, a valid order, algorithm choice, and O(V+E) time/space assumptions. A reader should find the decision before the derivation.

## 3. Worked case from the reference material

Represent the task input explicitly:

```go
type Task struct {
    ID           string
    Duration     int
    Prerequisites []string
}
```

Build an ID-to-index map, adjacency lists from prerequisite to dependent, and an indegree count. Validate IDs before scheduling. Keep validation errors distinct from cycle detection.

Worked example check: verify every edge in the output order and recompute each earliest-finish value from its predecessors.

## 4. Verification, failure, and evidence

**Verification.** Check empty input, one task, a chain, independent tasks, a diamond dependency, duplicate IDs, unknown prerequisite IDs, negative duration, and a directed cycle. For each valid output, check all prerequisite edges and recompute the reported completion time.

**Failure mode.** Introduce a directed cycle, then separately reference an unknown task ID. Explain why one is a validly parsed but unschedulable graph and the other is invalid input.

**Completion criterion.** You can model, validate, order, detect cycles, compute earliest finish times, state the complexity, and explain the result to a reader who needs the conclusion first.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

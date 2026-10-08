# Lesson 040 — Integrated Assessment: Algorithms Under Constraints

## Threads
**Primary:** Go/Algorithms/Communication

**Supporting:** blank-page implementation; explanation; trade-offs

**Math companion:** discrete math retrieval

**Time:** 3–4 hours. **Prerequisites:** Lesson 039.

**Time:** 3–4 hours. **Prerequisites:** Lesson 039.

## What you will be able to do

1. Implement a dependency planner that validates tasks and orders a directed acyclic graph.
2. Detect cycles and explain why no valid dependency order exists for cyclic input.
3. Calculate the earliest completion time for each task when task durations are supplied.
4. Present the contract, complexity, limitations, and one recommendation in a concise front-loaded report.

## Retrieve first (no notes)

1. Why does sorted input matter for two-pointer elimination?
2. What precondition supports a variable-size sum window?
3. What endpoint convention is required for interval merging?
4. How do you explain which proof changes when an assumption changes?

## 1. Problem and model

**What problem this solves.** Build and explain a small algorithmic system under explicit input and output constraints. The assessment combines graph modeling, topological ordering, dynamic programming on a DAG, validation, complexity analysis, and technical communication.

**The contract.** Input tasks with unique IDs, nonnegative durations, and prerequisite IDs. Reject duplicate IDs and unknown prerequisites. Return an order in which every prerequisite precedes its dependent task, or report a cycle. For a valid DAG, calculate each task's earliest finish from the maximum earliest finish among its prerequisites plus its own duration. Explain runtime in terms of `V` tasks and `E` prerequisite edges.

**Worked example.** Tasks: `A` duration 2 with no prerequisite; `B` duration 3 depends on A; `C` duration 4 depends on A; `D` duration 1 depends on B and C. A valid order places A before B/C and both before D. Earliest finishes are `A=2`, `B=5`, `C=6`, `D=7`; total completion time is 7.

Self-check: distinguish a topological order from a critical path. Several valid task orders may exist, while the completion time is determined by the longest-duration dependency path.

## 2. Mechanism and vocabulary

**Mechanism.** Kahn's algorithm tracks each vertex's indegree, queues vertices with indegree zero, emits one, and decreases the indegree of each outgoing neighbor. If fewer than `V` vertices are emitted, the graph contains a directed cycle. The traversal is O(V+E) with adjacency lists.

**Second point.** Earliest-finish dynamic programming follows topological order: for task `v`, `EF(v)=duration(v)+max(EF(p))` over prerequisites `p`; the maximum over no prerequisites is zero. This recurrence is valid because all prerequisite values are available before `v` is processed.

The report should lead with whether the input is valid and the completion time, then state the ordering and complexity. Do not call the longest path the topological order; those are different outputs.

**Worked example.** State the result first: “The four-task plan completes in 7 time units; the critical dependency chain is A → C → D.” Then give the validation rules, a valid order, algorithm choice, and O(V+E) time/space assumptions. A reader should find the decision before the derivation.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: add a dependency from D back to A; predict why the emitted count falls below the task count.

## 4. Verification, failure, and evidence

**Verification.** Check empty input, one task, a chain, independent tasks, a diamond dependency, duplicate IDs, unknown prerequisite IDs, negative duration, and a directed cycle. For each valid output, check all prerequisite edges and recompute the reported completion time.

**Failure mode.** Introduce a directed cycle, then separately reference an unknown task ID. Explain why one is a validly parsed but unschedulable graph and the other is invalid input.

**Completion criterion.** You can model, validate, order, detect cycles, compute earliest finish times, state the complexity, and explain the result to a reader who needs the conclusion first.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the earliest finish for each task in the four-task example.
2. Predict the number of vertices emitted after adding a back-edge from D to A.
3. Predict how Kahn's algorithm identifies that not all tasks can be ordered.
4. Predict the result of an unknown prerequisite ID.
5. Predict the result of duplicate task IDs.
6. Predict the asymptotic time when the task graph has `V` vertices and `E` edges.
7. Distinguish an arbitrary valid topological order from the critical path duration.
8. Draft the one-sentence conclusion before writing the technical explanation.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Define the task record and validation policy; build the prerequisite graph and indegrees.

2. Implement topological ordering and cycle detection; verify every returned edge ordering.

3. Compute earliest finish times in topological order and report the critical path duration.

4. Write a concise report leading with validity and completion time, followed by the algorithm, complexity, and limitations.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Add a cycle and demonstrate that no topological order can satisfy every prerequisite.

2. Add an unknown prerequisite and ensure it is reported as invalid input, not as a cycle.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one constraint: allow zero-duration tasks, or add multiple independent start tasks. State whether graph ordering and earliest-finish recurrence still apply and update validation accordingly.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What are the input-validation errors distinct from graph cycles?
2. How does Kahn's algorithm construct a topological order?
3. How does the algorithm detect a cycle?
4. How is earliest finish computed from prerequisites?
5. What is the time complexity in terms of vertices and edges?
6. How does the critical path differ from a topological order?

## Evidence to record

Create `lesson040/EVIDENCE.md` with exactly this structure:

```text
# Lesson 040 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Confusing invalid IDs with a directed cycle.
- Treating one topological order as the unique correct schedule.
- Computing earliest finish before all prerequisites have been processed.
- Reporting a critical path duration without stating task-duration assumptions.
- Burying the result below implementation details in the final report.

## Precision checks

1. State the input contract and validation errors.
2. Explain cycle detection using the emitted vertex count.
3. Verify every prerequisite edge in the result.
4. Recompute earliest finishes and the critical path duration.
5. Explain O(V+E) time and the storage used by the adjacency list.

The report is part of the engineering result: state the decision first, then give enough evidence and assumptions to review it.
## University revision bridge

For the LIG track, connect topological ordering and DAG dynamic programming to graph algorithms, then communicate the model and complexity in a concise technical report.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can implement a validated DAG dependency planner, detect cycles, compute the critical-path duration, verify every prerequisite edge, state the complexity, and present the result in a front-loaded engineering summary.

# Lesson 099 — Problem Set

**Topic:** Real Repositories, Bug Reproduction, and Upstreaming.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.

2. Extend it with one explicit boundary case: A project related to something covered elsewhere in this curriculum (a database engine, an HTTP framework, a CLI tool you use) is a reasonable, motivated choice.

3. Rebuild the same behavior with one changed constraint: Identify and trace the project's entry point(s): find `main()` (or equivalent), and trace the first few layers of what it calls, building an initial map of the major subsystems/modules it wires together.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/099-real-repositories-bugs-and-upstreaming/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Real-code bug hunt

Choose a real repository or issue, reproduce a bug from a clean checkout, isolate the cause, write a regression test, and prepare a minimal upstream-quality patch.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

# Lesson 099 — Real Repositories, Bug Reproduction, and Upstreaming Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.

2. A project related to something covered elsewhere in this curriculum (a database engine, an HTTP framework, a CLI tool you use) is a reasonable, motivated choice.

3. Identify and trace the project's entry point(s): find `main()` (or equivalent), and trace the first few layers of what it calls, building an initial map of the major subsystems/modules it wires together.

## Run and check

```powershell
Set-Location lessons/S2/099-real-repositories-bugs-and-upstreaming/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Real-code bug hunt

Choose a real repository or issue, reproduce a bug from a clean checkout, isolate the cause, write a regression test, and prepare a minimal upstream-quality patch.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

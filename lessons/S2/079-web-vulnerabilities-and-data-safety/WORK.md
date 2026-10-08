# Lesson 079 — Problem Set

**Topic:** Web Vulnerabilities and Data Safety.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).

2. Extend it with one explicit boundary case: Fix it using parameterized queries and confirm the same malicious input now has no effect beyond being treated as a literal, harmless string value.

3. Rebuild the same behavior with one changed constraint: Construct an XSS vulnerability (unescaped user content rendered into HTML) and demonstrate it by injecting a harmless but observable script (e.g. one that triggers a visible `alert()` or writes to the page in an obviously detectable way) — confirm it actually executes when the page is viewed.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/079-web-vulnerabilities-and-data-safety/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Web vulnerability lab

Take a deliberately unsafe web endpoint, reproduce a data or authorization flaw, patch it, and write a regression test that would fail before the fix.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.

# Lesson 079 — Web Vulnerabilities and Data Safety Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).

2. Fix it using parameterized queries and confirm the same malicious input now has no effect beyond being treated as a literal, harmless string value.

3. Construct an XSS vulnerability (unescaped user content rendered into HTML) and demonstrate it by injecting a harmless but observable script (e.g. one that triggers a visible `alert()` or writes to the page in an obviously detectable way) — confirm it actually executes when the page is viewed.

## Run and check

```powershell
Set-Location lessons/S2/079-web-vulnerabilities-and-data-safety/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Web vulnerability lab

Take a deliberately unsafe web endpoint, reproduce a data or authorization flaw, patch it, and write a regression test that would fail before the fix.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).

// STUDENT PROBLEM SET — Lesson 079.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).
// 2. Extend it with one explicit boundary case: Fix it using parameterized queries and confirm the same malicious input now has no effect beyond being treated as a literal, harmless string value.
// 3. Rebuild the same behavior with one changed constraint: Construct an XSS vulnerability (unescaped user content rendered into HTML) and demonstrate it by injecting a harmless but observable script (e.g. one that triggers a visible `alert()` or writes to the page in an obviously detectable way) — confirm it actually executes when the page is viewed. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 079: Web Vulnerabilities and Data Safety
// Your task: use parameterized SQL and handle the returned error correctly.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "database/sql"

func loadUser(db *sql.DB, name string) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = sql.DB

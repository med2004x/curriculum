// STUDENT PROBLEM SET — Lesson 067.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Pick a real feature from TARDOC or Mahall that involves an API endpoint you've already built (or plan to build) — for example, TARDOC's clinic subscription activation, or Mahall's product listing creation.
// 2. Extend it with one explicit boundary case: Before looking at (or without regard to) your existing implementation, write a plain-language requirements document answering: what resource is being acted on, what does the client send, what's returned on success, and what are at least 3 distinct ways this specific operation could fail, with a plain-English description of what each failure means.
// 3. Rebuild the same behavior with one changed constraint: For the same operation, explicitly list every field you'd include in the success response, and for each field, justify in one sentence why the client actually needs it (not "it was easy to include" or "it's just there internally") — if you can't justify a field this way, mark it as a candidate for removal from the contract. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 067: Validation, Error Contracts, and API Design
// Your task: define stable validation rules and error categories.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

type CreateUser struct{ Name string }

func (u CreateUser) Validate() error {
    // TODO: implement the assigned behavior.
}

func main() { fmt.Println((CreateUser{Name: "Ada"}).Validate()) }

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Errorf

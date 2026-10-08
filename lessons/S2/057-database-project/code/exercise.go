// STUDENT PROBLEM SET — Lesson 057.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Define task schema, migration sequence, API routes, and repository/service boundaries.
// 2. Implement handlers and a PostgreSQL repository with parameterized SQL and explicit error mapping.
// 3. Add integration tests for CRUD, invalid/not-found cases, transaction rollback, and fixture isolation.
// 4. Document the migration, least-privilege role, and a tested restore path. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 057: Integrated Lab 2: Durable Data Service
// Your task: define the repository boundary and durable transaction semantics.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

func saveOrder(id int) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {
	fmt.Println(saveOrder(1))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Errorf

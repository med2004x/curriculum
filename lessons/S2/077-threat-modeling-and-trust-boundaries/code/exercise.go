// STUDENT PROBLEM SET — Lesson 077.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Draw the data-flow diagram for browser/client, API, database, object store, and operator.
// 2. Identify assets, actors, entry points, trust boundaries, and sensitive flows.
// 3. Apply STRIDE prompts to flows and record prioritized threats, controls, and residual risks. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 077: Threat Modeling and Trust Boundaries
// Your task: make the trust boundaries explicit.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

type Decision struct {
	Authenticated bool
	Authorized    bool
	Allowed       bool
}

func evaluate(authenticated, authorized, trustedInput bool) Decision {
    // TODO: implement the assigned behavior.
    return Decision{}
}

func main() {}

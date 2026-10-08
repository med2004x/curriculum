// STUDENT PROBLEM SET — Lesson 093.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.
// 2. Extend it with one explicit boundary case: Using the math-for-engineering queueing lesson's M/M/1 formulas, and your step 1 peak `λ` estimate, solve for the minimum service rate `μ` required to keep average time-in-system `W` under a target (e.g.
// 3. Rebuild the same behavior with one changed constraint: 2 seconds) — show your work using the actual formulas, not just a rough guess. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 093: Requirements, Capacity, and Queueing
// Your task: calculate utilization and reject invalid rate assumptions.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

func utilization(arrivalRate, serviceRate float64) float64 {
    // TODO: implement the assigned behavior.
    return 0
}

func main() {
	fmt.Println(utilization(80, 100))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println

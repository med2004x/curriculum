# Lesson 067 — Validation, Error Contracts, and API Design

## Threads
**Primary:** Go/API

**Supporting:** input validation; stable errors; contracts

**Math companion:** logic and sets

**Time:** 3–4 hours. **Prerequisites:** Lesson 066.

**Time:** 3–4 hours. **Prerequisites:** Lesson 066, # Lesson 066 — Routing, CRUD, and HTTP Semantics.

## What you will be able to do

1. Explain the central model of # 067 — Validation, Error Contracts, and API Design and name the boundary where it applies.
2. Trace one concrete case of # 067 — Validation, Error Contracts, and API Design before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 066, # Lesson 066 — Routing, CRUD, and HTTP Semantics.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Treat an API as a contract negotiated *before* implementation, not a byproduct of whatever the handler code happens to do — write down the request/response shape and behavioral guarantees first, then build code that matches, rather than the reverse.

**The contract.** **Why "contract first" matters, concretely, not just as good practice.** An API is a promise to every client that calls it — mobile apps, other backend services, third-party integrators — none of whom can see your implementation, only the request/response shape and the documented behavior.

**Worked example.** Start with the smallest concrete case described by the source material: Pick a real feature from TARDOC or Mahall that involves an API endpoint you've already built (or plan to build) — for example, TARDOC's clinic subscription activation, or Mahall's product listing creation.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** If you build the handler first and let the API shape emerge from whatever was convenient to implement, you tend to leak implementation details into the contract (e.g. an internal database column name showing up as a JSON field, or an error message format that changes whenever you refactor internal error handling) — clients then depend on those accidental det.

**Second point.** **What "requirements" means at the API level, specifically.** Before writing a single handler, you should be able to answer, for each operation: what resource does this act on (Lesson 2 covers resource modeling properly), what does the client send, what does the client get back on success, what does the client get back on each distinct failure mode, and what.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Before looking at (or without regard to) your existing implementation, write a plain-language requirements document answering: what resource is being acted on, what does the client send, what's returned on success, and what are at least 3 distinct ways this specific operation could fail, with a plain-English description of what each failure means.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    count, _ := strconv.Atoi("12")
    fmt.Println(count)
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** For step 3, show the specific field(s) you identified as implementation leakage, alongside your reasoning for why a client shouldn't need to depend on it, and what you'd change.

**Failure mode.** Take your step 4 requirements document (for an operation you haven't built) and deliberately have someone else — or, if working alone, revisit it after a meaningful gap (a different day) with fresh eyes — attempt to implement a mock/stub version of the endpoint using *only* your written requirements, without asking you clarifying questions.

**Completion criterion.** You've written a complete requirements document for a real operation before referring to (or instead of referring to) its existing implementation, you've identified at least one real instance of implementation detail leaking into an existing API contract in your own code, and you've stress-tested a requirements document by having it (or a fresh read of it) a.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 067 — Validation, Error Contracts, and API Design case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Pick a real feature from TARDOC or Mahall that involves an API endpoint you've already built (or plan to build) — for example, TARDOC's clinic subscription activation, or Mahall's product listing creation.

2. Extend it with one explicit boundary case: Before looking at (or without regard to) your existing implementation, write a plain-language requirements document answering: what resource is being acted on, what does the client send, what's returned on success, and what are at least 3 distinct ways this specific operation could fail, with a plain-English description of what each failure means.

3. Rebuild the same behavior with one changed constraint: For the same operation, explicitly list every field you'd include in the success response, and for each field, justify in one sentence why the client actually needs it (not "it was easy to include" or "it's just there internally") — if you can't justify a field this way, mark it as a candidate for removal from the contract.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your step 4 requirements document (for an operation you haven't built) and deliberately have someone else — or, if working alone, revisit it after a meaningful gap (a different day) with fresh eyes — attempt to implement a mock/stub version of the endpoint using *only* your written requirements, without asking you clarifying questions.

2. Note every point where the requirements document turned out to be ambiguous or incomplete, requiring an assumption to be made that wasn't actually specified.

3. Explain why this exercise — not just writing requirements, but testing whether they're actually sufficient for someone else to implement against without guessing — is the real test of whether "contract-first" thinking actually happened, versus just writing documentation that still leaves the real design decisions implicit.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 067 — Validation, Error Contracts, and API Design: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 067 — Validation, Error Contracts, and API Design?
2. What two concrete practices are required in the assigned material for # 067 — Validation, Error Contracts, and API Design?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson067/EVIDENCE.md` with exactly this structure:

```text
# Lesson 067 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 067 — Validation, Error Contracts, and API Design.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 067 — Validation, Error Contracts, and API Design without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've written a complete requirements document for a real operation before referring to (or instead of referring to) its existing implementation, you've identified at least one real instance of implementation detail leaking into an existing API contract in your own code, and you've stress-tested a requirements document by having it (or a fresh read of it) a. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

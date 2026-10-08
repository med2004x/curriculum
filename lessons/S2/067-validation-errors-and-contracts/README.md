# Lesson 067 — Validation, Error Contracts, and API Design

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Treat an API as a contract negotiated *before* implementation, not a byproduct of whatever the handler code happens to do — write down the request/response shape and behavioral guarantees first, then build code that matches, rather than the reverse.

**The contract.** **Why "contract first" matters, concretely, not just as good practice.** An API is a promise to every client that calls it — mobile apps, other backend services, third-party integrators — none of whom can see your implementation, only the request/response shape and the documented behavior.

**Worked example.** Start with the smallest concrete case described by the source material: Pick a real feature from TARDOC or Mahall that involves an API endpoint you've already built (or plan to build) — for example, TARDOC's clinic subscription activation, or Mahall's product listing creation.

## 2. Mechanism and vocabulary

**Mechanism.** If you build the handler first and let the API shape emerge from whatever was convenient to implement, you tend to leak implementation details into the contract (e.g. an internal database column name showing up as a JSON field, or an error message format that changes whenever you refactor internal error handling) — clients then depend on those accidental det.

**Second point.** **What "requirements" means at the API level, specifically.** Before writing a single handler, you should be able to answer, for each operation: what resource does this act on (Lesson 2 covers resource modeling properly), what does the client send, what does the client get back on success, what does the client get back on each distinct failure mode, and what.

**Worked example.** Before looking at (or without regard to) your existing implementation, write a plain-language requirements document answering: what resource is being acted on, what does the client send, what's returned on success, and what are at least 3 distinct ways this specific operation could fail, with a plain-English description of what each failure means.

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

## 4. Verification, failure, and evidence

**Verification.** For step 3, show the specific field(s) you identified as implementation leakage, alongside your reasoning for why a client shouldn't need to depend on it, and what you'd change.

**Failure mode.** Take your step 4 requirements document (for an operation you haven't built) and deliberately have someone else — or, if working alone, revisit it after a meaningful gap (a different day) with fresh eyes — attempt to implement a mock/stub version of the endpoint using *only* your written requirements, without asking you clarifying questions.

**Completion criterion.** You've written a complete requirements document for a real operation before referring to (or instead of referring to) its existing implementation, you've identified at least one real instance of implementation detail leaking into an existing API contract in your own code, and you've stress-tested a requirements document by having it (or a fresh read of it) a.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

# Lesson 066 — Routing, CRUD, and HTTP Semantics

## Threads
**Primary:** Go/API/Networking

**Supporting:** routing; verbs; status codes

**Math companion:** state machines

**Time:** 3–4 hours. **Prerequisites:** Lesson 065.

**Time:** 3–4 hours. **Prerequisites:** Lesson 065, # Lesson 065 — Go HTTP, JSON, and SQL.

## What you will be able to do

1. Explain the central model of # 066 — Routing, CRUD, and HTTP Semantics and name the boundary where it applies.
2. Trace one concrete case of # 066 — Routing, CRUD, and HTTP Semantics before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 065, # Lesson 065 — Go HTTP, JSON, and SQL.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Apply HTTP's method semantics (safety, idempotency) and status codes correctly and consistently across a real API, going beyond networking Lesson 4's transport-level mechanics into how an API's design should respect and communicate through these semantics.

**The contract.** **Safety and idempotency, applied at the API-design level (not just protocol trivia).** `GET` must be safe (no side effects) — an API that implements `GET /clinics/{id}/deactivate` (a side-effecting operation exposed via GET) violates this convention and creates real risk:

**Worked example.** Start with the smallest concrete case described by the source material: Audit one of your own existing API endpoints (TARDOC or Mahall) and check whether its HTTP method usage matches this lesson's safety/idempotency conventions — specifically, confirm no GET endpoint has side effects, and check whether your DELETE endpoints behave idempotently (calling delete on an already-deleted resource — what does it currently return, and i.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** `PUT` and `DELETE` should be idempotent — calling `DELETE /clinics/{id}` twice should have the same end effect as calling it once (the clinic is deleted; a second call finding it already gone should return a clear "already deleted" response, typically still treated as success or a specific 404, not an error implying something went wrong).

**Second point.** `POST` is neither safe nor generally idempotent by convention — which is exactly why distributed Lesson 2's request-ID-based idempotency mechanism matters specifically for `POST` operations representing non-idempotent creates, since HTTP itself provides no built-in protection there the way it conventionally implies for `PUT`/`DELETE`.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** For the same endpoint (or a small set of related endpoints), produce a status-code table: every distinct outcome (success and each failure mode) mapped to a specific status code, and confirm you're not reusing the same status code for two semantically different outcomes, or using different codes for what's actually the same kind of outcome across similar end.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
type Bookmark struct {
    ID    int
    URL   string
    Label string
}

type bookmarkStore struct {
    nextID    int
    bookmarks []Bookmark
}

func newBookmarkStore() *bookmarkStore {
    return &bookmarkStore{nextID: 1}
}

func (store *bookmarkStore) create(url, label string) Bookmark {
    bookmark := Bookmark{
        ID:    store.nextID,
        URL:   url,
        Label: label,
    }
    store.nextID++
    store.bookmarks = append(store.bookmarks, bookmark)
    return bookmark
}

func (store *bookmarkStore) find(id int) (Bookmark, bool) {
    for _, bookmark := range store.bookmarks {
        if bookmark.ID == id {
            return bookmark, true
        }
    }
    return Bookmark{}, false
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Present your step 2 status-code table for review, and for step 3, show actual `curl` output (status code and body) for both the "always 200" bad version and your corrected version, demonstrating the practical difference in what a status-code-only check can detect.

**Failure mode.** Take your corrected step 3 endpoint (correct status codes) and write a small generic client function that retries any request receiving a 5xx response (up to 3 times with backoff) but does *not* retry on 4xx.

**Completion criterion.** You've audited a real endpoint against HTTP method safety/idempotency conventions and identified any violations, you've produced a complete, consistent status-code table for a real set of endpoints and fixed at least one status-code misuse you found, and you've demonstrated — with actual working retry-logic code — why correct status code usage is what makes.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 066 — Routing, CRUD, and HTTP Semantics case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Audit one of your own existing API endpoints (TARDOC or Mahall) and check whether its HTTP method usage matches this lesson's safety/idempotency conventions — specifically, confirm no GET endpoint has side effects, and check whether your DELETE endpoints behave idempotently (calling delete on an already-deleted resource — what does it currently return, and i.

2. Extend it with one explicit boundary case: For the same endpoint (or a small set of related endpoints), produce a status-code table: every distinct outcome (success and each failure mode) mapped to a specific status code, and confirm you're not reusing the same status code for two semantically different outcomes, or using different codes for what's actually the same kind of outcome across similar end.

3. Rebuild the same behavior with one changed constraint: Deliberately construct a bad example: an endpoint that returns `200 OK` even when the operation actually failed (with the actual error only discoverable by inspecting the response body), and rewrite it to return an appropriate 4xx or 5xx status code instead, with the error detail in the body.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your corrected step 3 endpoint (correct status codes) and write a small generic client function that retries any request receiving a 5xx response (up to 3 times with backoff) but does *not* retry on 4xx.

2. Test it against both a genuinely transient 5xx scenario (simulate your server temporarily failing, e.g. by adding a deliberate `if firstAttempt { return 500 }` in test code) and a genuine 4xx scenario (send a deliberately malformed request).

3. Confirm the generic retry logic correctly recovers from the transient 5xx (retries succeed) and correctly avoids wasting retries on the un-retriable 4xx (fails fast instead).

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 066 — Routing, CRUD, and HTTP Semantics: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 066 — Routing, CRUD, and HTTP Semantics?
2. What two concrete practices are required in the assigned material for # 066 — Routing, CRUD, and HTTP Semantics?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson066/EVIDENCE.md` with exactly this structure:

```text
# Lesson 066 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 066 — Routing, CRUD, and HTTP Semantics.
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

You pass when you can explain the central model of # 066 — Routing, CRUD, and HTTP Semantics without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've audited a real endpoint against HTTP method safety/idempotency conventions and identified any violations, you've produced a complete, consistent status-code table for a real set of endpoints and fixed at least one status-code misuse you found, and you've demonstrated — with actual working retry-logic code — why correct status code usage is what makes. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

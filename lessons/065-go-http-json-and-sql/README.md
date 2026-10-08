# Lesson 065 — Go HTTP, JSON, and SQL

## Threads
**Primary:** Go/Backend/DB

**Supporting:** handlers; encoding; persistence

**Math companion:** data modeling

**Time:** 3–4 hours. **Prerequisites:** Lesson 064.

**Time:** 3–4 hours. **Prerequisites:** Lesson 064.

## What you will be able to do

1. Decode a JSON request body in a Go `net/http` handler and validate the request shape.
2. Use parameterized SQL to persist and retrieve the corresponding record.
3. Bound request-body size and handle malformed JSON, missing rows, and database failures distinctly.
4. Explain which layer owns HTTP parsing, domain validation, and persistence.

## Retrieve first (no notes)

1. What is the path from user syscall stub through kernel dispatch and back?
2. Why must a kernel syscall validate user-provided addresses?
3. What is the kernel invariant your selected xv6 lab preserves?
4. What does a lab test establish, and what does it leave unproven?

## 1. Problem and model

**What problem this solves.** Carry one request through HTTP, JSON, validation, SQL, and a response while keeping each boundary explicit and failures classifiable.

**The contract.** Read one bounded JSON object from the request body, reject malformed or trailing data according to a stated policy, validate required fields, execute a parameterized query with the request context, and map outcomes to suitable HTTP responses. User input remains a SQL value, never interpolated SQL syntax.

**Worked example.** A `POST /tasks` body contains `{"title":"Review"}`. Decode into a typed request, reject blank title, call a repository method using a bound parameter, and return a JSON response with the created ID and an appropriate success status.

Self-check: what happens if the body exceeds the configured size, contains invalid JSON, or contains a second JSON value after the first object?

## 2. Mechanism and vocabulary

**Mechanism.** `http.MaxBytesReader` bounds the body; `json.Decoder` parses structured input; validation checks domain requirements; `database/sql` parameters bind data values; the response encoder serializes only the public response type.

**Second point.** Keep errors distinct: malformed body is a client error, invalid domain fields are a validation response, no matching record is not-found, and unexpected database failures are server errors. Log internal details safely; do not expose SQL text, credentials, or driver diagnostics to clients.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Query with `db.QueryRowContext(r.Context(), "SELECT ... WHERE id = $1", id).Scan(...)`. Map `sql.ErrNoRows` to a domain not-found error; preserve other errors for internal logging and return a generic server response.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
func handler(w http.ResponseWriter, r *http.Request)
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** For a valid create request, record status, response JSON, and persisted row. Also cover malformed JSON, oversized body, invalid title, SQL constraint failure, canceled context, and read of a missing ID.

**Failure mode.** Replace the parameter placeholder with direct string interpolation in a test-only example and explain the injection boundary. Restore parameterization; never deploy the intentionally vulnerable form.

**Completion criterion.** You can trace a request from HTTP bytes to typed JSON, validated domain input, parameterized SQL, and a stable public response while preserving cancellation and error distinctions.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the response for valid JSON with a valid title.
2. Predict the response for malformed JSON and an oversized body.
3. Predict what happens when a second JSON value follows the first object.
4. Predict the database result for an unknown ID.
5. Predict how a canceled request context affects a context-aware query.
6. Predict why SQL placeholders prevent user text from becoming syntax.
7. Distinguish a validation failure from a driver/query failure.
8. State which response fields are public and which database details remain internal.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Define typed create-request and response structs and a bounded JSON decoder.

2. Validate the request and call a context-aware repository method using parameters.

3. Map success, malformed input, invalid fields, not-found, and internal error to distinct outcomes.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Handle connection failure, query cancellation, no-row results, and a transaction rollback.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one boundary: add a maximum title length, or make the request context expire during persistence. Explain which layer enforces the change and what response contract follows.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. How does `http.MaxBytesReader` constrain a request?
2. Why use a typed request struct for JSON?
3. Why should database calls receive the request context?
4. How should `sql.ErrNoRows` differ from an unexpected SQL error?
5. Why bind values instead of concatenating them into SQL?
6. Which internal details should not appear in an error response?

## Evidence to record

Create `lesson065/EVIDENCE.md` with exactly this structure:

```text
# Lesson 065 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Decoding an unbounded body from an untrusted client.
- Treating successful JSON decoding as domain validation.
- Interpolating request text into SQL.
- Returning every storage error as not-found or as raw driver text.
- Returning success before persistence reports success.

## Precision checks

1. Trace request bytes to decoded fields and validation decisions.
2. Identify the exact parameterized SQL boundary.
3. State the HTTP mapping for malformed input, invalid fields, missing data, and database failure.
4. Explain request-context cancellation.
5. List the body-size and response-data limits.

Each layer should transform data under an explicit contract: HTTP framing, JSON structure, domain rules, and SQL persistence are not interchangeable validations.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 065 — Go HTTP, JSON, and SQL without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can swap storage implementations without rewriting business rules. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

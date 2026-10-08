# Lesson 065 — Go HTTP, JSON, and SQL

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Carry one request through HTTP, JSON, validation, SQL, and a response while keeping each boundary explicit and failures classifiable.

**The contract.** Read one bounded JSON object from the request body, reject malformed or trailing data according to a stated policy, validate required fields, execute a parameterized query with the request context, and map outcomes to suitable HTTP responses. User input remains a SQL value, never interpolated SQL syntax.

**Worked example.** A `POST /tasks` body contains `{"title":"Review"}`. Decode into a typed request, reject blank title, call a repository method using a bound parameter, and return a JSON response with the created ID and an appropriate success status.

## 2. Mechanism and vocabulary

**Mechanism.** `http.MaxBytesReader` bounds the body; `json.Decoder` parses structured input; validation checks domain requirements; `database/sql` parameters bind data values; the response encoder serializes only the public response type.

**Second point.** Keep errors distinct: malformed body is a client error, invalid domain fields are a validation response, no matching record is not-found, and unexpected database failures are server errors. Log internal details safely; do not expose SQL text, credentials, or driver diagnostics to clients.

**Worked example.** Query with `db.QueryRowContext(r.Context(), "SELECT ... WHERE id = $1", id).Scan(...)`. Map `sql.ErrNoRows` to a domain not-found error; preserve other errors for internal logging and return a generic server response.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
func handler(w http.ResponseWriter, r *http.Request)
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** For a valid create request, record status, response JSON, and persisted row. Also cover malformed JSON, oversized body, invalid title, SQL constraint failure, canceled context, and read of a missing ID.

**Failure mode.** Replace the parameter placeholder with direct string interpolation in a test-only example and explain the injection boundary. Restore parameterization; never deploy the intentionally vulnerable form.

**Completion criterion.** You can trace a request from HTTP bytes to typed JSON, validated domain input, parameterized SQL, and a stable public response while preserving cancellation and error distinctions.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

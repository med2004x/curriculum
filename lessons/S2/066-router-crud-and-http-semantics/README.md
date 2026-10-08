# Lesson 066 — Routing, CRUD, and HTTP Semantics

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Apply HTTP's method semantics (safety, idempotency) and status codes correctly and consistently across a real API, going beyond networking Lesson 4's transport-level mechanics into how an API's design should respect and communicate through these semantics.

**The contract.** **Safety and idempotency, applied at the API-design level (not just protocol trivia).** `GET` must be safe (no side effects) — an API that implements `GET /clinics/{id}/deactivate` (a side-effecting operation exposed via GET) violates this convention and creates real risk:

**Worked example.** Start with the smallest concrete case described by the source material: Audit one of your own existing API endpoints (TARDOC or Mahall) and check whether its HTTP method usage matches this lesson's safety/idempotency conventions — specifically, confirm no GET endpoint has side effects, and check whether your DELETE endpoints behave idempotently (calling delete on an already-deleted resource — what does it currently return, and i.

## 2. Mechanism and vocabulary

**Mechanism.** `PUT` and `DELETE` should be idempotent — calling `DELETE /clinics/{id}` twice should have the same end effect as calling it once (the clinic is deleted; a second call finding it already gone should return a clear "already deleted" response, typically still treated as success or a specific 404, not an error implying something went wrong).

**Second point.** `POST` is neither safe nor generally idempotent by convention — which is exactly why distributed Lesson 2's request-ID-based idempotency mechanism matters specifically for `POST` operations representing non-idempotent creates, since HTTP itself provides no built-in protection there the way it conventionally implies for `PUT`/`DELETE`.

**Worked example.** For the same endpoint (or a small set of related endpoints), produce a status-code table: every distinct outcome (success and each failure mode) mapped to a specific status code, and confirm you're not reusing the same status code for two semantically different outcomes, or using different codes for what's actually the same kind of outcome across similar end.

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

## 4. Verification, failure, and evidence

**Verification.** Present your step 2 status-code table for review, and for step 3, show actual `curl` output (status code and body) for both the "always 200" bad version and your corrected version, demonstrating the practical difference in what a status-code-only check can detect.

**Failure mode.** Take your corrected step 3 endpoint (correct status codes) and write a small generic client function that retries any request receiving a 5xx response (up to 3 times with backoff) but does *not* retry on 4xx.

**Completion criterion.** You've audited a real endpoint against HTTP method safety/idempotency conventions and identified any violations, you've produced a complete, consistent status-code table for a real set of endpoints and fixed at least one status-code misuse you found, and you've demonstrated — with actual working retry-logic code — why correct status code usage is what makes.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

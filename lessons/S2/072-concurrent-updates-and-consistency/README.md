# Lesson 072 — Concurrent Updates and Consistency

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Detect and correctly reject conflicting concurrent updates to the same resource using optimistic concurrency control (ETags/version numbers), rather than letting the last write silently win and destroy an earlier, unseen change.

**The contract.** **The lost-update problem, restated at the API level.** Two clients both `GET /clinics/{id}`, both see the same current state, both independently decide on a change, both `PUT` their updated version back — the second `PUT` overwrites the first, silently discarding the first client's change, with neither client aware anything went wrong.

**Worked example.** Start with the smallest concrete case described by the source material: Implement `GET /resource/{id}` returning an `ETag` header (a simple version integer, or a hash of the resource's serialized content, is fine) alongside the resource body.

## 2. Mechanism and vocabulary

**Mechanism.** This is precisely database-internals Lesson 5's lost-update anomaly, except now it's happening across two separate HTTP requests (each internally a properly isolated database transaction) rather than within one database transaction — meaning the database's own isolation guarantees don't automatically protect against it, since each individual `PUT` request is.

**Second point.** **Optimistic concurrency control: detect conflicts instead of preventing them upfront.** Rather than locking the resource for the duration of a client's "think time" between GET and PUT (which would be both impractical over HTTP's stateless request model and a real availability cost, since one slow client would block others), optimistic concurrency lets both.

**Worked example.** Implement `PUT /resource/{id}` requiring an `If-Match` header matching the resource's current ETag.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```sql
UPDATE app.clinics
SET name = $1,
		version = version + 1
WHERE id = $2
	AND version = $3
RETURNING id, name, version;
```

If no row is returned, the ID may not exist or the version is stale; resolve that distinction according to the API's disclosure policy. The version predicate and update occur as one database statement, so two clients with the same old version cannot both update successfully.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** For step 3, show both clients' actual requests (including their `If-Match` headers) and the server's actual responses, confirming the first succeeds (200 with new ETag) and the second is correctly rejected (412), with the resource's final state matching only the first client's intended change, not silently reflecting a merged or overwritten result.

**Failure mode.** Implement a version of `PUT` that *ignores* the `If-Match` header entirely (accepts and applies any update regardless of the client-provided version) and rerun step 3's exact concurrent-update scenario against this broken version.

**Completion criterion.** Your `PUT` endpoint correctly enforces `If-Match` and rejects stale updates with 412, you've directly reproduced the lost-update problem both with and without protection using the same concurrent-client test scenario, and you've implemented a working client-side retry-after-conflict flow that successfully reconciles and resubmits after a detected conflict ra.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

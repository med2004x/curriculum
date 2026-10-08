# Lesson 079 — Web Vulnerabilities and Data Safety

## Threads
**Primary:** Security/API/DB

**Supporting:** injection; XSS/CSRF concepts; safe data handling

**Math companion:** logic and adversarial reasoning

**Time:** 3–4 hours. **Prerequisites:** Lesson 078.

**Time:** 3–4 hours. **Prerequisites:** Lesson 078.

## What you will be able to do

1. Identify trust-boundary crossings for SQL, HTML, browser state-changing requests, and server-side URL fetches.
2. Reproduce each vulnerability only in a local toy application with synthetic data.
3. Apply parameterized SQL, contextual output escaping, CSRF defenses, and restricted outbound destinations.
4. Test vulnerable behavior and patched behavior without targeting real systems.

## Retrieve first (no notes)

1. What does HMAC authenticate, and what does it not encrypt?
2. Why should signature verification use a standard cryptographic comparison?
3. Why is a fast SHA-256 hash unsuitable for storing passwords?
4. What must be true of a public key before trusting a valid signature?

## 1. Problem and model

**What problem this solves.** Reproduce, understand, and patch the core web vulnerability classes — injection, CSRF, SSRF, XSS — in local, controlled examples, understanding each as a specific instance of untrusted input crossing a trust boundary (system-engineering Lesson 10) without adequate handling.

**The contract.** Untrusted input must not be interpreted as control syntax in another context. Parameterized SQL separates values from query syntax. HTML output must be escaped for its context. State-changing browser requests need a CSRF defense appropriate to their authentication model. Server-side fetches must restrict destinations and validate the resolved network target.

**Worked example.** Start with the smallest concrete case described by the source material: In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** SQL placeholders bind values, not identifiers or arbitrary clauses; dynamic identifiers require an allow-list. HTML escaping is context-specific: text, attribute, URL, and script contexts differ. Go's `html/template` performs contextual escaping when templates are used correctly; sanitization is not a universal substitute for output encoding.

**Second point.** CSRF abuses a browser's ambient credentials to trigger an unwanted state change; defenses may include SameSite cookie policy, anti-CSRF tokens, and origin checks. SSRF abuses server-side request capability; validate schemes/hosts, restrict resolved destinations according to policy, and account for redirects and DNS rebinding.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Fix it using parameterized queries and confirm the same malicious input now has no effect beyond being treated as a literal, harmless string value.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** In local fixtures with synthetic data, record the vulnerable observation, patch, and retest: SQL structure remains fixed; an HTML marker renders as text; a cross-site state-changing request is rejected under the selected CSRF policy; an outbound fetch to a prohibited destination is denied.

**Failure mode.** Take your step 2 XSS fix (proper output escaping) and construct a *second*, different injection point in the same application that you didn't originally fix (e.g. if you fixed escaping in a comment-rendering template, try injecting via a different field, like a username, rendered in a different part of the page that uses a different, unescaped rendering path.

**Completion criterion.** You can name each trust boundary, apply its context-specific defense, and show local before/after evidence for SQL injection, XSS, CSRF, and SSRF.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether a SQL parameter can alter the query's syntax tree.
2. Predict how contextual HTML escaping changes a synthetic marker in a local page.
3. Predict which state-changing request is protected by the selected CSRF policy.
4. Predict whether a server-side fetch to a disallowed destination should proceed.
5. Predict how a redirect can affect destination validation.
6. Predict whether escaping one template location protects a second output sink.
7. Identify the trust boundary for each test fixture.
8. State the expected before/after evidence before patching.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).

2. Extend it with one explicit boundary case: Fix it using parameterized queries and confirm the same malicious input now has no effect beyond being treated as a literal, harmless string value.

3. Rebuild the same behavior with one changed constraint: Construct an XSS vulnerability (unescaped user content rendered into HTML) and demonstrate it by injecting a harmless but observable script (e.g. one that triggers a visible `alert()` or writes to the page in an obviously detectable way) — confirm it actually executes when the page is viewed.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your step 2 XSS fix (proper output escaping) and construct a *second*, different injection point in the same application that you didn't originally fix (e.g. if you fixed escaping in a comment-rendering template, try injecting via a different field, like a username, rendered in a different part of the page that uses a different, unescaped rendering path.

2. Confirm this second injection point is still vulnerable despite your first fix, demonstrating that fixing one specific instance of a vulnerability class doesn't automatically fix every instance — each place untrusted data crosses into a rendering (or query, or fetch) context needs its own correct handling, and a security review needs to systematically check.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 079 — Web Vulnerabilities and Data Safety: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Why do SQL parameters keep values separate from query syntax?
2. Why is HTML escaping context-sensitive?
3. What browser behavior does CSRF rely on?
4. What makes a server-side URL fetch a trust boundary?
5. Why must SSRF defenses consider redirects and resolved addresses?
6. Why must every output sink be reviewed independently?

## Evidence to record

Create `lesson079/EVIDENCE.md` with exactly this structure:

```text
# Lesson 079 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating validation as a replacement for parameterized SQL.
- Applying HTML escaping to the wrong output context.
- Assuming CORS alone prevents CSRF.
- Checking an SSRF hostname once while ignoring redirects or resolution changes.
- Fixing one output sink and assuming all templates are safe.

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

You pass when you can explain the central model of # 079 — Web Vulnerabilities and Data Safety without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've successfully demonstrated and then correctly fixed all four vulnerability classes in local, controlled examples, with real before/after evidence for each, and you've demonstrated — via the failure drill — that fixing one instance of a vulnerability class doesn't protect every instance, understanding why a systematic audit across all trust-boundary cro. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

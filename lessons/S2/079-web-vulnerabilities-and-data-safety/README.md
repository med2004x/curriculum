# Lesson 079 — Web Vulnerabilities and Data Safety

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Reproduce, understand, and patch the core web vulnerability classes — injection, CSRF, SSRF, XSS — in local, controlled examples, understanding each as a specific instance of untrusted input crossing a trust boundary (system-engineering Lesson 10) without adequate handling.

**The contract.** Untrusted input must not be interpreted as control syntax in another context. Parameterized SQL separates values from query syntax. HTML output must be escaped for its context. State-changing browser requests need a CSRF defense appropriate to their authentication model. Server-side fetches must restrict destinations and validate the resolved network target.

**Worked example.** Start with the smallest concrete case described by the source material: In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).

## 2. Mechanism and vocabulary

**Mechanism.** SQL placeholders bind values, not identifiers or arbitrary clauses; dynamic identifiers require an allow-list. HTML escaping is context-specific: text, attribute, URL, and script contexts differ. Go's `html/template` performs contextual escaping when templates are used correctly; sanitization is not a universal substitute for output encoding.

**Second point.** CSRF abuses a browser's ambient credentials to trigger an unwanted state change; defenses may include SameSite cookie policy, anti-CSRF tokens, and origin checks. SSRF abuses server-side request capability; validate schemes/hosts, restrict resolved destinations according to policy, and account for redirects and DNS rebinding.

**Worked example.** Fix it using parameterized queries and confirm the same malicious input now has no effect beyond being treated as a literal, harmless string value.

## 3. Worked case from the reference material

Worked example: In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input).

## 4. Verification, failure, and evidence

**Verification.** In local fixtures with synthetic data, record the vulnerable observation, patch, and retest: SQL structure remains fixed; an HTML marker renders as text; a cross-site state-changing request is rejected under the selected CSRF policy; an outbound fetch to a prohibited destination is denied.

**Failure mode.** Take your step 2 XSS fix (proper output escaping) and construct a *second*, different injection point in the same application that you didn't originally fix (e.g. if you fixed escaping in a comment-rendering template, try injecting via a different field, like a username, rendered in a different part of the page that uses a different, unescaped rendering path.

**Completion criterion.** You can name each trust boundary, apply its context-specific defense, and show local before/after evidence for SQL injection, XSS, CSRF, and SSRF.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

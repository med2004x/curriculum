# Lesson 069 — Authentication, Authorization, and Threats

## Threads
**Primary:** Security/API

**Supporting:** identity; permissions; threat boundaries

**Math companion:** sets and predicates

**Time:** 3–4 hours. **Prerequisites:** Lesson 068.

**Time:** 3–4 hours. **Prerequisites:** Lesson 068, # Lesson 068 — Handlers, Services, Repositories, and Boundaries.

## What you will be able to do

1. Explain the central model of # 069 — Authentication, Authorization, and Threats and name the boundary where it applies.
2. Trace one concrete case of # 069 — Authentication, Authorization, and Threats before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 068, # Lesson 068 — Handlers, Services, Repositories, and Boundaries.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Implement token-based authentication and resource-level authorization correctly, and understand precisely why these are two separate, sequential checks (who are you, then what are you allowed to do) rather than one combined concept.

**The contract.** Authentication establishes a principal from a credential; authorization decides whether that principal may perform an action on a particular resource. A valid role is not proof of ownership of every resource with that role's type.

**Worked example.** Issue a short-lived signed access token after credential verification. Middleware validates the signature using an allowed algorithm and validates issuer, audience, and expiration before establishing a principal. A signed JWT is generally readable by its holder; signing provides integrity/authenticity, not confidentiality. Missing or invalid credentials receive `401 Unauthorized` with an appropriate challenge policy.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Authorization evaluates the authenticated principal, action, and resource. A valid token with insufficient general permission is typically `403 Forbidden`; a caller who is allowed to use the endpoint but does not own the requested object must still be denied by a resource-level policy. Some APIs intentionally return 404 to conceal resource existence; document that policy consistently.

**Second point.** Bearer tokens grant access to whoever possesses them. Send them only over TLS, avoid logging them, use short expiry and deliberate rotation/revocation policy, and keep signing keys out of source control. JWT is one token format, not a requirement for all APIs.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Add a role policy and a separate ownership query for `GET /clinics/{id}/records`. Test missing token, invalid token, authenticated user with insufficient role, correct role for another user's clinic, and correct role for the user's clinic. These are distinct authentication/authorization outcomes.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a minimal token-based authentication scheme: a login endpoint that verifies credentials and issues a token (a signed JWT is a reasonable, realistic choice, using a Go library like `golang-jwt/jwt` or equivalent), and middleware that validates the token on subsequent requests, rejecting requests with missing or invalid tokens with `401 Unauthorized`.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Show request/response pairs for missing/invalid credential, insufficient role, wrong resource ownership, and allowed access. Record status and response shape while ensuring no token or sensitive record data is exposed in logs.

**Failure mode.** Deliberately implement the exact vulnerability described in Learn: an endpoint that checks only the user's role (e.g. "is this user clinic-staff") without checking whether the specific requested resource ID actually belongs to them.

**Completion criterion.** You can distinguish authentication, role authorization, and resource authorization; validate token claims securely; and demonstrate that sharing a role does not grant access to another principal's resource.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 069 — Authentication, Authorization, and Threats case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement a minimal token-based authentication scheme: a login endpoint that verifies credentials and issues a token (a signed JWT is a reasonable, realistic choice, using a Go library like `golang-jwt/jwt` or equivalent), and middleware that validates the token on subsequent requests, rejecting requests with missing or invalid tokens with `401 Unauthorized`.

2. Extend it with one explicit boundary case: Implement role-based authorization: attach a role claim to the issued token, and add middleware or handler-level logic that checks the required role for a given endpoint, rejecting requests from an authenticated-but-wrong-role user with `403 Forbidden` (distinct from the 401 in step 1 — test both cases and confirm the status codes differ correctly).

3. Rebuild the same behavior with one changed constraint: Implement resource-level (ownership) authorization on top of role-based checks: for an endpoint like `GET /clinics/{id}/records`, in addition to checking the user has an appropriate role, explicitly check that the requested clinic ID actually belongs to (or is otherwise legitimately accessible by) the authenticated user — and test the specific vulnerability.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Deliberately implement the exact vulnerability described in Learn: an endpoint that checks only the user's role (e.g. "is this user clinic-staff") without checking whether the specific requested resource ID actually belongs to them.

2. Using a real test with two different clinic-staff users (each associated with a different clinic), confirm user A can successfully access user B's clinic's data purely because both share the `clinic-staff` role, despite having no legitimate relationship to each other's clinic.

3. Explain, using this concrete, successfully-reproduced vulnerability, why "the user has the right role" and "the user is allowed to access this specific resource" are genuinely different checks that must both be performed — and why a code review or test suite that only exercises the happy path (a user accessing their *own* resources) would never catch this bu.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one authorization rule: introduce a read-only role, or allow an explicit delegated clinic relationship. State how the principal/resource policy changes and add both allowed and denied tests.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the difference between authentication and authorization?
2. What does a JWT signature establish, and what does it not conceal?
3. Which token claims should a verifier validate besides the signature?
4. Why is a role check insufficient for resource ownership?
5. When might an API use 404 rather than 403 for a denied resource?
6. Why must bearer tokens be protected from logs and plaintext transport?

## Evidence to record

Create `lesson069/EVIDENCE.md` with exactly this structure:

```text
# Lesson 069 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a signed JWT as encrypted or secret.
- Trusting token claims without validating signature, issuer, audience, and expiry.
- Treating a role as authorization for every resource.
- Returning 401 for every authenticated-but-forbidden request without a documented policy.
- Logging bearer tokens or signing secrets.

## Precision checks

1. State the authentication and authorization checks in order.
2. List the token verification checks and key-management assumptions.
3. Identify the role and resource-ownership predicates separately.
4. Show allowed and denied request cases.
5. State the API's 403/404 concealment policy and logging protections.

Security decisions are principal-action-resource checks; a credential is only the start of that decision.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 069 — Authentication, Authorization, and Threats without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've implemented and correctly distinguished 401 (authentication failure) from 403 (authorization failure) in your own working code, you've implemented both role-based and resource-ownership-based authorization as genuinely separate checks, and you've deliberately reproduced the cross-resource-access vulnerability from Learn with two real test users, confi. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

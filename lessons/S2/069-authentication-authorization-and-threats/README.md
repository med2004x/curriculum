# Lesson 069 — Authentication, Authorization, and Threats

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Implement token-based authentication and resource-level authorization correctly, and understand precisely why these are two separate, sequential checks (who are you, then what are you allowed to do) rather than one combined concept.

**The contract.** Authentication establishes a principal from a credential; authorization decides whether that principal may perform an action on a particular resource. A valid role is not proof of ownership of every resource with that role's type.

**Worked example.** Issue a short-lived signed access token after credential verification. Middleware validates the signature using an allowed algorithm and validates issuer, audience, and expiration before establishing a principal. A signed JWT is generally readable by its holder; signing provides integrity/authenticity, not confidentiality. Missing or invalid credentials receive `401 Unauthorized` with an appropriate challenge policy.

## 2. Mechanism and vocabulary

**Mechanism.** Authorization evaluates the authenticated principal, action, and resource. A valid token with insufficient general permission is typically `403 Forbidden`; a caller who is allowed to use the endpoint but does not own the requested object must still be denied by a resource-level policy. Some APIs intentionally return 404 to conceal resource existence; document that policy consistently.

**Second point.** Bearer tokens grant access to whoever possesses them. Send them only over TLS, avoid logging them, use short expiry and deliberate rotation/revocation policy, and keep signing keys out of source control. JWT is one token format, not a requirement for all APIs.

**Worked example.** Add a role policy and a separate ownership query for `GET /clinics/{id}/records`. Test missing token, invalid token, authenticated user with insufficient role, correct role for another user's clinic, and correct role for the user's clinic. These are distinct authentication/authorization outcomes.

## 3. Worked case from the reference material

Worked example: Implement a minimal token-based authentication scheme: a login endpoint that verifies credentials and issues a token (a signed JWT is a reasonable, realistic choice, using a Go library like `golang-jwt/jwt` or equivalent), and middleware that validates the token on subsequent requests, rejecting requests with missing or invalid tokens with `401 Unauthorized`.

## 4. Verification, failure, and evidence

**Verification.** Show request/response pairs for missing/invalid credential, insufficient role, wrong resource ownership, and allowed access. Record status and response shape while ensuring no token or sensitive record data is exposed in logs.

**Failure mode.** Deliberately implement the exact vulnerability described in Learn: an endpoint that checks only the user's role (e.g. "is this user clinic-staff") without checking whether the specific requested resource ID actually belongs to them.

**Completion criterion.** You can distinguish authentication, role authorization, and resource authorization; validate token claims securely; and demonstrate that sharing a role does not grant access to another principal's resource.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).

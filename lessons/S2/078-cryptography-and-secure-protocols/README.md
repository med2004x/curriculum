# Lesson 078 — Cryptography and Secure Protocols

## Threads
**Primary:** Security/Networking

**Supporting:** hashes; MACs; signatures; key exchange

**Math companion:** modular arithmetic

**Time:** 3–4 hours. **Prerequisites:** Lesson 077.

**Time:** 3–4 hours. **Prerequisites:** Lesson 077.

## What you will be able to do

1. Distinguish cryptographic hashes, keyed MACs, and public-key signatures by their guarantees and key requirements.
2. Use standard-library or established-library APIs to compute and verify SHA-256, HMAC, and a digital signature.
3. Explain why password storage needs a salted, deliberately expensive password-hashing function rather than fast SHA-256.
4. Identify key validation, algorithm selection, and constant-time verification requirements at a protocol boundary.

## Retrieve first (no notes)

1. What assets and trust boundaries did the threat model identify?
2. Why does a role check not authorize every resource of that type?
3. How does an HMAC protect a webhook body?
4. Why is a signed token not necessarily confidential?

## 1. Problem and model

**What problem this solves.** Use standard cryptographic primitives (hashing, HMAC, signatures) correctly in real code, and understand precisely why "rolling your own crypto" is dangerous — not as a platitude, but by directly experiencing a specific, common way naive cryptographic code fails.

**The contract.** A cryptographic hash is deterministic and fixed-length; security properties include preimage, second-preimage, and collision resistance. A hash is not encryption and does not in general hide low-entropy input. HMAC combines a hash with a secret key for message authentication; a digital signature uses a private key to sign and a public key to verify.

**Worked example.** Compute SHA-256 for two inputs differing by one byte using the standard library. Observe that the digests differ; do not treat one pair as proof of collision resistance or claim every output bit must change.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Use a password-hashing scheme designed to be salted and computationally expensive, such as Argon2id, scrypt, bcrypt, or PBKDF2 with an appropriate current work factor. Store the encoded parameters/salt/hash format returned by the library; do not invent a custom password hash.

**Second point.** For HMAC verification, compare MAC bytes with the cryptographic library's constant-time verification/comparison function. For digital signatures, ensure the public key is trusted through the protocol's certificate/key distribution rules; mathematical verification alone does not identify a trusted publisher.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Implement HMAC-based message signing and verification (extending API engineering Lesson 11's practical webhook use directly), and confirm that altering either the message or the signature (but not both consistently) causes verification to correctly fail — test both cases explicitly, not just the happy path.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Compute SHA-256 hashes of several inputs (using your language's standard crypto library, never a hand-rolled implementation) and confirm the hash changes completely for even a single-character input change (the "avalanche effect" — a defining property of a good cryptographic hash), demonstrating this isn't a simple checksum where small input changes produce.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Present your step 2 HMAC verification results (both the correct-pass and correct-fail cases), your step 3 signature verification results (correct key pair passes, wrong key pair fails), and your step 4 brute-force time estimates for bcrypt versus plain SHA-256, showing the concrete, large difference in attacker cost.

**Failure mode.** Replace the expected HMAC key or public key with an unrelated key and verify that authentication fails. Separately compare password-verification work for fast SHA-256 and a standard password hash; explain why the fast hash is unsuitable for password storage. Do not implement a real endpoint that accepts a hand-rolled password verifier.

**Completion criterion.** You can state what each primitive guarantees and requires, verify valid and tampered messages using supported libraries, and select a password-hashing scheme rather than a fast general-purpose hash.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict which party can verify an HMAC and what key it needs.
2. Predict which party can create a signature and which key verifies it.
3. Predict what an altered message does to HMAC/signature verification.
4. Explain why a cryptographic hash does not conceal a predictable password.
5. Predict the effect of using a wrong public key.
6. Predict why constant-time MAC comparison is preferred.
7. State which key must remain secret in each construction.
8. Distinguish integrity, authenticity, confidentiality, and password-guessing resistance.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Compute SHA-256 hashes of several inputs (using your language's standard crypto library, never a hand-rolled implementation) and confirm the hash changes completely for even a single-character input change (the "avalanche effect" — a defining property of a good cryptographic hash), demonstrating this isn't a simple checksum where small input changes produce.

2. Extend it with one explicit boundary case: Implement HMAC-based message signing and verification (extending API engineering Lesson 11's practical webhook use directly), and confirm that altering either the message or the signature (but not both consistently) causes verification to correctly fail — test both cases explicitly, not just the happy path.

3. Generate a signing key pair using an established library, sign a message, verify with the correct public key, then verify failure with an unrelated public key.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Implement a deliberately naive "custom" password verification scheme: instead of using bcrypt's built-in, constant-time comparison (which handles the timing-safety concern automatically), implement your own comparison of a computed hash against a stored one using a standard, non-constant-time string comparison (`==` or equivalent).

2. This directly reuses API engineering Lesson 11's constant-time-comparison concern, now applied to password verification specifically rather than webhook signatures.

3. You don't need to demonstrate a successful timing attack (genuinely hard to do reliably in a small exercise), but research and explain, in your own words, why a naive comparison here reintroduces exactly the timing side-channel risk that lesson identified, and why using your standard library's provided, already-hardened comparison function (rather than writi.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one primitive: replace HMAC with a digital signature or change password-hash parameters. Explain which trust/key model changes and which guarantees remain.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does a cryptographic hash provide?
2. How does HMAC differ from an unkeyed hash?
3. How do digital-signature keys differ from HMAC keys?
4. Why is SHA-256 unsuitable for direct password storage?
5. What does constant-time MAC comparison address?
6. Why does a valid signature not by itself establish trust in a public key?

## Evidence to record

Create `lesson078/EVIDENCE.md` with exactly this structure:

```text
# Lesson 078 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a hash as encryption or as proof of publisher identity.
- Storing passwords with a fast general-purpose hash.
- Comparing authentication MACs with a handwritten non-constant-time loop.
- Reusing an HMAC secret as a public verification key.
- Trusting a public key without authenticating its distribution.

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

You pass when you can explain the central model of # 078 — Cryptography and Secure Protocols without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've correctly used hashing, HMAC, and digital signatures in real, working code with both success and failure cases tested explicitly, you've quantitatively demonstrated — with real benchmark numbers — why slow password hashing matters against brute-force attacks, and you can explain, using the constant-time-comparison example, a specific, concrete way tha. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

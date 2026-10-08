# Lesson 046 — TLS, Cryptography, and Trust

## Threads
**Primary:** Networking/Security

**Supporting:** handshake; certificates; confidentiality/integrity

**Math companion:** modular arithmetic foundations

**Time:** 3–4 hours. **Prerequisites:** Lesson 045.

**Time:** 3–4 hours. **Prerequisites:** Lesson 045, # Lesson 045 — HTTP and net/http.

## What you will be able to do

1. Explain the central model of # 046 — TLS, Cryptography, and Trust and name the boundary where it applies.
2. Trace one concrete case of # 046 — TLS, Cryptography, and Trust before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 045, # Lesson 045 — HTTP and net/http.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Understand what TLS actually provides — encryption, authentication, integrity — and specifically what a certificate proves (and doesn't prove), by inspecting a real handshake rather than treating "HTTPS" as an opaque padlock icon.

**The contract.** **Three separate guarantees, commonly conflated into "it's secure."** TLS provides: **confidentiality** (an eavesdropper on the network path can't read the content — the connection is encrypted), **integrity** (the content can't be silently modified in transit without detection — any tampering breaks a cryptographic check), and **authentication** (the client.

**Worked example.** Start with the smallest concrete case described by the source material: Generate a local, self-signed TLS certificate (e.g.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** These are genuinely separable properties; understanding them separately matters because a system can have some without others (e.g. an encrypted-but-unauthenticated connection is still vulnerable to a man-in-the-middle who simply presents their own valid encryption keys, since nothing was checking *whose* keys they were).

**Second point.** **What a certificate actually proves.** A certificate is a signed statement, from a Certificate Authority (CA) that the client already trusts (via a pre-installed list of trusted root CAs), asserting "this public key belongs to this domain name." When your browser connects to `example.com` and receives a certificate, verification means:

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 1 -nodes`) and use it to run a minimal HTTPS server (Go's `net/http` with `ListenAndServeTLS`, pointing at your generated cert/key, is straightforward for this).

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Generate a local, self-signed TLS certificate (e.g.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** For step 3, report the actual certificate chain length (how many certificates were presented) and the negotiated cipher suite for a real production server, and state in your own words what each certificate in the chain is vouching for relative to the next one up the chain.

**Failure mode.** Take your step 1 self-signed server and, using a packet capture tool (Lesson 7 preview) or a network proxy tool, attempt to read the actual HTTP request/response content passing over your `curl -k` connection at the raw TCP byte level (not through the TLS-aware tool itself, but genuinely at the wire).

**Completion criterion.** You've run a real TLS handshake against your own self-signed server and observed both the encrypted-but-unverified case and the properly-rejected-for-lack-of-trust case, you've inspected a real production certificate chain and can explain what each link in it vouches for, and you can state clearly, using your own failure-drill observation of genuinely unread.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 046 — TLS, Cryptography, and Trust case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Generate a local, self-signed TLS certificate (e.g.

2. Extend it with one explicit boundary case: `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 1 -nodes`) and use it to run a minimal HTTPS server (Go's `net/http` with `ListenAndServeTLS`, pointing at your generated cert/key, is straightforward for this).

3. Rebuild the same behavior with one changed constraint: Connect to it with `curl -k` (the `-k` flag skips certificate verification, appropriate here since you know it's self-signed and you're testing locally) and confirm the connection succeeds and is encrypted.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your step 1 self-signed server and, using a packet capture tool (Lesson 7 preview) or a network proxy tool, attempt to read the actual HTTP request/response content passing over your `curl -k` connection at the raw TCP byte level (not through the TLS-aware tool itself, but genuinely at the wire).

2. Confirm the payload is unreadable ciphertext, not plaintext HTTP — directly demonstrating the confidentiality guarantee is real and independent of whether authentication was properly verified in this specific test (since you used `-k` to skip that check).

3. Explain why this is the concrete, hands-on version of the abstract "separable guarantees" point from Learn: your step 1/2 test showed authentication can fail while encryption still works, and this drill shows encryption genuinely works (not just claimed) even in a scenario where you deliberately bypassed authentication.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 046 — TLS, Cryptography, and Trust: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 046 — TLS, Cryptography, and Trust?
2. What two concrete practices are required in the assigned material for # 046 — TLS, Cryptography, and Trust?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson046/EVIDENCE.md` with exactly this structure:

```text
# Lesson 046 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 046 — TLS, Cryptography, and Trust.
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

You pass when you can explain the central model of # 046 — TLS, Cryptography, and Trust without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've run a real TLS handshake against your own self-signed server and observed both the encrypted-but-unverified case and the properly-rejected-for-lack-of-trust case, you've inspected a real production certificate chain and can explain what each link in it vouches for, and you can state clearly, using your own failure-drill observation of genuinely unread. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.

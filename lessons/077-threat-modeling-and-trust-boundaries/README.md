# Lesson 077 — Threat Modeling and Trust Boundaries

## Threads
**Primary:** Security/Systems

**Supporting:** assets; attackers; boundaries; abuse cases

**Math companion:** probability and risk

**Time:** 3–4 hours. **Prerequisites:** Lesson 076.

**Time:** 3–4 hours. **Prerequisites:** Lesson 076.

## What you will be able to do

1. Identify assets, actors, entry points, data flows, and trust boundaries for a small service.
2. Use STRIDE as a prompt to identify concrete threats rather than as a checklist detached from architecture.
3. Map each high-priority threat to a mitigation and a residual-risk statement.
4. Distinguish authentication, authorization, validation, and operational controls in the threat model.

## Retrieve first (no notes)

1. What resource and invariant did the backend capstone define?
2. Which layer maps domain errors to public HTTP responses?
3. Why test ownership authorization separately from role authorization?
4. What cross-feature interaction did the capstone integration test cover?

## 1. Problem and model

**What problem this solves.** Identify what must be protected, who can influence it, and where trust changes before choosing a security mechanism.

**The contract.** Cover assets, actors, trust boundaries, attack surfaces, abuse cases, least privilege, and a practical STRIDE-style review.

**Worked example.** Start with the smallest concrete case described by the source material: Threat-model a small task API with browser clients, an API service, PostgreSQL, object storage, and an administrator.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Draw the data-flow diagram with external actors, processes, data stores, and trust-boundary crossings. For each flow, identify what data crosses, who controls it, and what identity or integrity claim is relied upon.

**Second point.** STRIDE prompts are Spoofing, Tampering, Repudiation, Information disclosure, Denial of service, and Elevation of privilege. Apply them to concrete components/flows, prioritize by impact and feasibility, then identify controls and remaining risk.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** For `browser → API → PostgreSQL/object storage`, consider a forged identity token, a cross-tenant record ID, an oversized upload, a leaked signed URL, and a forged webhook. Map each threat to the relevant trust boundary and control (token validation, ownership check, request limits, short-lived scoped URL, HMAC verification and deduplication).

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Threat-model a small task API with browser clients, an API service, PostgreSQL, object storage, and an administrator.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Produce a diagram and a threat register containing asset, actor, entry point/flow, threat, impact, mitigation, owner, and residual risk. Ensure each mitigation addresses the stated threat, not merely a neighboring concern.

**Failure mode.** Remove the resource-ownership check while leaving authentication and role checks in place. Show which trust assumption permits cross-tenant access and identify the missing authorization predicate.

**Completion criterion.** Every high-priority threat has a concrete boundary, mitigation, and residual-risk owner; controls are justified by the system's assets and flows.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Identify the assets and actors in the task-service diagram.
2. Predict which flow carries an untrusted resource ID.
3. Predict what happens when role checks pass but ownership checks are absent.
4. Identify one STRIDE threat on the webhook flow.
5. Identify one denial-of-service threat on the upload or API boundary.
6. Predict what control mitigates the threat and what residual risk remains.
7. Predict which components store secrets or sensitive records.
8. State the trust boundary before assigning a control.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Draw the data-flow diagram for browser/client, API, database, object store, and operator.

2. Identify assets, actors, entry points, trust boundaries, and sensitive flows.

3. Apply STRIDE prompts to flows and record prioritized threats, controls, and residual risks.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one trust boundary: expose the API to an untrusted integration partner or move object storage behind signed URLs. Update actors, flows, threats, and controls accordingly.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What are the components in a data-flow diagram?
2. What trust-boundary crossing does an external request make?
3. Name the STRIDE categories.
4. Why does a role check not replace resource ownership authorization?
5. What fields belong in a threat register?
6. What is residual risk?

## Evidence to record

Create `lesson077/EVIDENCE.md` with exactly this structure:

```text
# Lesson 077 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Listing generic threats without identifying the flow or asset.
- Adding controls without stating which threat they mitigate.
- Treating authentication as resource authorization.
- Ignoring third-party integrations, operators, or stored secrets as actors/assets.
- Omitting residual risk and ownership after selecting a control.

## Precision checks

1. Show assets, actors, entry points, flows, and trust boundaries.
2. Tie each high-priority threat to a concrete control.
3. State one residual risk and its owner.
4. Demonstrate the cross-tenant failure when ownership checks are removed.
5. Prioritize threats by impact and feasibility rather than list length.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 077 — Threat Modeling and Trust Boundaries without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Every security control in your design has a stated threat or trust-boundary reason. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
